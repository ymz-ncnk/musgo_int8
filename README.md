# MusGo_int8
Provides serialization of the Golang's `int8` type.

# How to use
Simply cast `int8` to `musgo_int8.Int8`:
```go
	var n int8 = 5
	// Marshal
	buf := make([]byte, musgo_int8.Int8(n).SizeMUS())
	musgo_int8.Int8(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_int8.Int8)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

