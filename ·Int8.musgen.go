// Code generated by musgen. DO NOT EDIT.

package musgo_int8

import "github.com/ymz-ncnk/musgo/errs"

// MarshalMUS fills buf with the MUS encoding of v.
func (v Int8) MarshalMUS(buf []byte) int {
	i := 0
	{
		buf[i] = byte(v)
		i++
	}
	return i
}

// UnmarshalMUS parses the MUS-encoded buf, and sets the result to *v.
func (v *Int8) UnmarshalMUS(buf []byte) (int, error) {
	i := 0
	var err error
	{
		if i > len(buf)-1 {
			return i, errs.ErrSmallBuf
		}
		(*v) = Int8(buf[i])
		i++
	}
	return i, err
}

// SizeMUS returns the size of the MUS-encoded v.
func (v Int8) SizeMUS() int {
	size := 0
	{
		_ = v
		size++
	}
	return size
}
