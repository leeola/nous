// Nous is a project for storing, retrieving and retaining personal knowledge.
//
// IMPORTANT: This project is unstable and the UX is being dogfooded.
//
// Each piece of "information" should be bite sized pieces of information,
// representing a single, verifiable fact. Knowledge should be browsable,
// and represented in a method similar to a Mind Map.
package nous

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/leeola/fixity"
	"github.com/leeola/fixity/value"
)

const FixityNamespace = "nous"

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

func (n *Nous) Store(ctx context.Context, d Data) error {
	switch d.Type {
	case TypeText:
		if d.Text == nil {
			return fmt.Errorf("text cannot be nil with %s", d.Type)
		}
	default:
		return fmt.Errorf("unexpected data type: %s", d.Type)
	}

	v := textToValues(d)

	id := strconv.FormatInt(time.Now().Unix(), 10)

	_, err := n.s.WriteNamespace(ctx, id, FixityNamespace, v, nil)
	if err != nil {
		return fmt.Errorf("store: %v", err)
	}

	return nil
}

func textToValues(info Data) fixity.Values {
	v := fixity.Values{
		"parentId": value.String(info.ParentID),
		"name":     value.String(info.Name),
		"content":  value.String(info.Text.Content),
		"value":    value.String(info.Text.Value),
	}
	return v
}
