/*
	Complete a function named GetCandidates, this function takes a DB object and a slice of
	strings(dynamic array) representing a list of subejcts. The function needs to make some
	checks and then return a slice(dynamic array) of students that match provided subjects
	Go Documentation: https://godoc.org
*/

package q4

// DB object, top level object
type DB struct {
	SubjectInfo []Subject
}

// Subject data structure that contains info abotu subjects
type Subject struct {
	SubjectName string
	Student     []Candidate
}

// Candidate data structure containing name and roll  number of candidate
type Candidate struct {
	Name string
	Roll int
}

// GetCandidates function, please do NOT change it's name :)
func GetCandidates(db DB, subjects []string) []Candidate {
	// Your code goes here
}
