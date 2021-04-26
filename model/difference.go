package model

// DifferenceEntryType define the type of an entry
type DifferenceEntryType int

const (
	DET_ROOT DifferenceEntryType = iota
	DET_OBJECT
	DET_ARRAY
	DET_KEY_VALUE
)

// DifferenceEntryStatus enum to mark entry as added, deleted or neutral
type DifferenceEntryStatus int

const (
	DES_NEUTRAL DifferenceEntryStatus = iota
	DES_ADD
	DES_REMOVE
)

// DifferenceEntry define a difference item that can be rendered
type DifferenceEntry struct {
	Parent                *DifferenceEntry
	ChildEntries          []*DifferenceEntry
	Key                   string
	Value                 interface{}
	ArrayHelpMarker       bool
	DifferenceEntryStatus DifferenceEntryStatus
	DifferenceEntryType   DifferenceEntryType
}

func (de *DifferenceEntry) AddChild(key string, value interface{}, des DifferenceEntryStatus, det DifferenceEntryType) *DifferenceEntry {
	newChild := &DifferenceEntry{
		Parent:                de,
		ChildEntries:          []*DifferenceEntry{},
		Key:                   key,
		Value:                 value,
		DifferenceEntryStatus: des,
		DifferenceEntryType:   det,
	}

	de.ChildEntries = append(de.ChildEntries, newChild)

	return newChild
}
