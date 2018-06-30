// Nous is a project for storing, retrieving and retaining personal knowledge.
//
// IMPORTANT: This project is unstable and the UX is being dogfooded.
//
// Each piece of "information" should be bite sized pieces of information,
// representing a single, verifiable fact. Knowledge should be browsable,
// and represented in a method similar to a Mind Map.
//
//
// IMPORTANT: The v0 flatdisk implementation is a brute force, ugly
// implementation of the Nous interface. Strictly designed for a PoC to suss
// out the user experience and methods of storing (and retrieving) information
// relationships.
package nous

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config for nous.
type Config struct {
	// Path to store information entries.
	Path string
}

// Nous information storage and .
type Nous struct {
	root string
}

// New constructs a new nous instance from the given config.
func New(c Config) (*Nous, error) {
	if c.Path == "" {
		return nil, errors.New("missing required config: Path")
	}

	return &Nous{
		root: c.Path,
	}, nil
}

func (n *Nous) Store(i Information) error {
	if i.Content == "" {
		return errors.New("information content cannot be empty")
	}

	b, err := json.Marshal(i)
	if err != nil {
		return fmt.Errorf("failed to marshal information: %s", err)
	}

	err = ioutil.WriteFile(filepath.Join(n.root), b, 0644)
	if err != nil {
		return fmt.Errorf("failed to write information to disk: %s", err)
	}

	return nil
}

// NOTE: this is an especially inefficient implementation of retrieving
// documents with matching information.
func (n *Nous) Retrieve(tags ...string) ([]Information, error) {
	fis, err := ioutil.ReadDir(n.root)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %s", err)
	}

	var matched []Information
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

func openInfo(path string) (Information, error) {
	f, err := os.Open(path)
	if err != nil {
		return Information{}, fmt.Errorf("failed to open file %s: %s", path, err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return Information{}, fmt.Errorf("failed to open file %s: %s", path, err)
	}

	var info Information
	if err := json.Unmarshal(b, &info); err != nil {
		return Information{}, fmt.Errorf("failed to marshal info: %s", err)
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
