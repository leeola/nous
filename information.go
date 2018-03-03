package nous

// Information is a unit of information within a Nous node.
type Information struct {
	// Content of the information.
	//
	// The text in this field will be indexed and searchable.
	Content string

	// Tags are the main method of organizing information in Nous.
	//
	// The values themselves and how they organize information in
	// Nous is still very in flux. The goal is to have tags
	// represent relationships between pieces of information.
	//
	// Whether tags represent open many to many relationships,
	// or a hierarchy of tags pointing to a single location
	// is undecided.
	Tags []string

	// Value is an optional field representing a singular value for this
	// given information.
	//
	// For example, if an address is to be stored and easily retrieved later,
	// the Content might contain the address with a *description* of the
	// address. The description in the Content helps it be searchable and
	// gives it context. Yet, if you wanted to copy the address to enter
	// in a form, you wouldn't want the Content. Value allows you to store
	// *just* the value.
	Value *string

	//Retain      *Retain
}

// type Retain struct {
// 	Question *string
// 	Answer   *string
// }
//
// type Tree struct {
// 	Node        string
// 	Related     []Tree
// 	Information []Information
// }
