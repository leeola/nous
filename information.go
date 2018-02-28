package nous

// Information is a unit of information within a Nous node.
type Information struct {
	Content string

	Tags []string

	// Clipboard is an optional value which will
	Clipboard *string

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
