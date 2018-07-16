package nous

//go:generate stringer -type=Type -output=info_string.go

type Type int

const (
	TypeText Type = iota + 1
	TypeImage
)

// Data is a unit of information within a Nous node.
type Data struct {
	Type Type `json:"nousType"`

	ID string `json:"id"`

	ParentID string `json:"parentId"`

	Meta bool `json:"meta"`

	Name string `json:"name"`

	Text  *DataText  `json:"text,omitempty"`
	Inv   *DataInv   `json:"inv,omitempty"`
	Image *DataImage `json:"image,omitempty"`
}

type DataText struct {
	// Content of the information.
	//
	// The text in this field will be indexed and searchable.
	Content string

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
}

type DataImage struct {
	DataFile
}

type DataInv struct {
}

type DataFile struct {
	// Ext of the file, not including the prefix ".".
	//
	// Prefix '.' removed simply to remove pointless
	// character. Additional '.' chars should be included,
	// as with 'tar.gz' files.
	Ext string `json:"ext"`
}

// Various quiz types, kept around for eventual re-integration.
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
