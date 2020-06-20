package custombool

import "testing"

func TestBool_UnmarshalJSON(t *testing.T) {
	var bit Bool
	args := map[string]bool{
		"1":        true,
		"true":     true,
		"0":        false,
		"false":    false,
		"whatever": false,
	}
	wantErr := false

	for arg, want := range args {
		if err := (&bit).UnmarshalJSON([]byte(arg)); (err != nil) != wantErr {
			t.Errorf("Bool.UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if bool(bit) != want {
			t.Errorf("Bool.UnmarshalJSON() error, want %t, got %t", want, bit)
		}
	}
}
