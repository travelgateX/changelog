package model

// Category : Commit category.
type Category string

// Added : that comment its mandatory to export type?
const (
	Added      = Category("ADDED")
	Changed    = Category("CHANGED")
	Deprecated = Category("DEPRECATED")
	Removed    = Category("REMOVED")
	Fixed      = Category("FIXED")
	Security   = Category("SECURITY")
)
