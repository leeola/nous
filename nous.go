// Nous is a project for storing, retrieving and retaining personal knowledge.
//
// IMPORTANT: This project is unstable and the UX is being dogfooded.
//
// Each piece of "information" should be bite sized pieces of information,
// representing a single, verifiable fact. Knowledge should be browsable,
// and represented in a method similar to a Mind Map.
package nous

type Nous interface {
	// Store information into the Nous implementation.
	Store(Information) (hash string, err error)

	// Retrieve information based on the given tags.
	Retrieve(tag ...string) ([]Information, error)

	// Retain() error
	// Research() error
}
