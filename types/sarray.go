package types

type SArray []string

func (i *SArray) String() string {
	return "my string representation"
}

// Set ...
func (i *SArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}
