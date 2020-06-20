package custombool

// Bool allows 0/1 to also become boolean.
type Bool bool

// UnmarshalJSON provide custom bool json decoding
func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*bit = Bool(txt == "1" || txt == "true")
	return nil
}
