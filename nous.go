// Nous is a project for storing, retrieving and retaining personal knowledge.
//
// IMPORTANT: This project is unstable and the UX is being dogfooded.
//
// Each piece of "information" should be bite sized pieces of information,
// representing a single, verifiable fact. Knowledge should be browsable,
// and represented in a method similar to a Mind Map.
package nous

import (
	"errors"
	"fmt"

	"github.com/leeola/fixity"
)

// Nous information storage and .
type Nous struct {
	s fixity.Store
}

// New constructs a new nous instance from the given store.
func New(s fixity.Store) (*Nous, error) {
	return &Nous{
		s: s,
	}, nil
}

func (n *Nous) Store(d Data) error {
	switch d.Type {
	case TypeText:
		if d.Text == nil {
			return fmt.Errorf("text cannot be nil with %s", d.Type)
		}
	default:
		return fmt.Errorf("unexpected data type: %s", d.Type)
	}

	return errors.New("not implemented")
}
