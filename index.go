package mfc

// Index common index
type Index interface {
	// Append element
	Append(item interface{}) (err error)

	// Delete element
	Delete(item interface{}) (err error)
}
