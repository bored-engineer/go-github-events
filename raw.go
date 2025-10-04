package github

// RawEvent is an unparsed event payload.
type RawEvent []byte

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *RawEvent) UnmarshalJSON(data []byte) error {
	*e = append((*e)[:0], data...)
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (e *RawEvent) MarshalJSON() ([]byte, error) {
	return *e, nil
}
