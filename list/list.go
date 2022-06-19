package list

// Structure: Linked List
// Description: Contains Nodes. Each node holds data and a reference to another list node.
// Functions: 
// 		- Insert onto the end of list
//		- Insert at a certain index
// 		- Print the entire list
// 		- Search for index of specific value
//    - Return a node from a given index
// 		- Delete node from list given index
// 		- Delete node from list given value

type Node struct {
	value		int
	next 		*Node
}

type LinkedList struct {
	head 		*Node
	len 		int
}

// Insert onto the end of list
func (l *LinkedList) Insert(value int) {}
// Insert at a certain index
func (l *LinkedList) InsertAt(value int, index int) {}
// Print the entire list
func (l *LinkedList) Print() {}
// Search for index of specific value
func (l *LinkedList) Search(value int) int {}
// Return a node from a given index
func (l *LinkedList) GetAt(index int) *Node {}
// Delete node from list given index
func (l *LinkedList) DeleteAt(index int) error {}
// Delete node from list given value
func (l *LinkedList) DeleteValue(value int) error {}
