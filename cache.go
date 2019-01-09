package mfc

// Cache - cache common
type Cache interface {
	// Append element
	Append(item interface{}) (err error)

	// Delete element
	Delete(item interface{}) (err error)
}
