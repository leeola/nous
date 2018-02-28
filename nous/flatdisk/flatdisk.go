package flatdisk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/leeola/nous"
	multihash "github.com/multiformats/go-multihash"
)

var (
	multihashCode uint64
)

func init() {
	c, ok := multihash.Names["blake2b-256"]
	if !ok {
		panic("multihash name not found")
	}
	multihashCode = c
}

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
