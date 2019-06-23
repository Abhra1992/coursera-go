package types

// SArray type alias for a string array
type SArray []string

// String representation of the array
func (i *SArray) String() string {
	return "my string representation"
}

// Set adds values to the array
func (i *SArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}
