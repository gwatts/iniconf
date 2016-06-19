package iniconf

// Reader provides read access to a section
// It is implemented by both Iniconf and ConfChain
type Reader interface {
	HasSection(sectionName string) bool
	EntryString(sectionName, entryName string) (string, error)
	EntryInt(sectionName, entryName string) (int64, error)
	EntryBool(sectionName, entryName string) (bool, error)
}
