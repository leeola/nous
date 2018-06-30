// The flatdisk implementation is a brute force, ugly implementation of the
// Nous interface. Strictly designed for a PoC to suss out the user experience
// and methods of storing (and retrieving) information relationships.
package flatdisk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/leeola/nous"
	multihash "github.com/multiformats/go-multihash"
)

var multihashCode = multihash.Names["blake2b-32"]

// Config for the flatdisk nous implementation.
type Config struct {
	// Path to store information entries.
	Path string
}

// Nous flatdisk implementation.
type Nous struct {
	root string
}

// New constructs a new flatdisk nous implementation.
func New(c Config) (*Nous, error) {
	if c.Path == "" {
		return nil, errors.New("missing required config: Path")
	}

	return &Nous{
		root: c.Path,
	}, nil
}

func (n *Nous) Store(i nous.Information) (string, error) {
	if i.Content == "" {
		return "", errors.New("information content cannot be empty")
	}

	// TODO(leeola): hash the whole information.
	//
	// hash the info content to get the address.
	// -1 uses the default size
	infoAddrM, err := multihash.Sum([]byte(i.Content), multihashCode, -1)
	if err != nil {
		return "", fmt.Errorf("failed to hash information: %s", err)
	}

	b, err := json.Marshal(i)
	if err != nil {
		return "", fmt.Errorf("failed to marshal information: %s", err)
	}

	err = ioutil.WriteFile(filepath.Join(n.root, infoAddrM.HexString()), b, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write information to disk: %s", err)
	}

	return infoAddrM.B58String(), nil
}

// NOTE: this is an especially inefficient implementation of retrieving
// documents with matching information.
func (n *Nous) Retrieve(tags ...string) ([]nous.Information, error) {
	fis, err := ioutil.ReadDir(n.root)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %s", err)
	}

	var matched []nous.Information
	for _, fi := range fis {
		info, err := openInfo(filepath.Join(n.root, fi.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed to open %s: %s", fi.Name(), err)
		}

		if !matchTags(tags, info.Tags) {
			continue
		}

		matched = append(matched, info)
	}

	return matched, nil
}

func openInfo(path string) (nous.Information, error) {
	f, err := os.Open(path)
	if err != nil {
		return nous.Information{}, fmt.Errorf("failed to open file %s: %s", path, err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nous.Information{}, fmt.Errorf("failed to open file %s: %s", path, err)
	}

	var info nous.Information
	if err := json.Unmarshal(b, &info); err != nil {
		return nous.Information{}, fmt.Errorf("failed to marshal info: %s", err)
	}

	return info, nil
}

func matchTags(required, tags []string) bool {
	// looking up via a map might be faster, but the flatdisk
	// implementation is super temporary, so i'm not worried.
	for _, req := range required {
		var hasReq bool
		for _, tag := range tags {
			if req == tag {
				hasReq = true
				break
			}
		}

		if !hasReq {
			return false
		}
	}

	return true
}
