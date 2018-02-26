package nous

type Nous interface {
	// Store information into the Nous implementation.
	Store(Information) (hash string, err error)

	// Retrieve(node ...string) (Node, error)

	// Retain() error
	// Research() error
}
