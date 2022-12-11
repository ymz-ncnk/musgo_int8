package musgo_int8

import "testing"

func TestMusgoInt8(t *testing.T) {
	var n int8 = 5
	buf := make([]byte, Int8(n).SizeMUS())
	Int8(n).MarshalMUS(buf)

	var an int8
	(*Int8)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
