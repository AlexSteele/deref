package deref

func Float32(f *float32) float32 {
	if f == nil {
		return 0.0
	}
	return *f
}

func Float64(f *float64) float64 {
	if f == nil {
		return 0.0
	}
	return *f
}

func Int(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func Int16(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}

func Int32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func Int64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func Int8(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}

func Rune(r *rune) rune {
	if r == nil {
		return 0
	}
	return *r
}

func String(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func Uint(i *uint) uint {
	if i == nil {
		return 0
	}
	return *i
}

func Uint16(i *uint16) uint16 {
	if i == nil {
		return 0
	}
	return *i
}

func Uint32(i *uint32) uint32 {
	if i == nil {
		return 0
	}
	return *i
}

func Uint64(i *uint64) uint64 {
	if i == nil {
		return 0
	}
	return *i
}

func Uint8(i *uint8) uint8 {
	if i == nil {
		return 0
	}
	return *i
}
