package deref

import (
	"testing"
)

func Test(t *testing.T) {
	var (
		a *float32
		b *float64
		c *int
		d *int16
		e *int32
		f *int64
		g *int8
		h *rune
		i *string
		j *uint
		k *uint16
		l *uint32
		m *uint64
		n *uint8
	)
	if Float32(a) != 0.0 {
		t.Fail()
	}
	if Float64(b) != 0.0 {
		t.Fail()
	}
	if Int(c) != 0 {
		t.Fail()
	}
	if Int16(d) != 0 {
		t.Fail()
	}
	if Int32(e) != 0 {
		t.Fail()
	}
	if Int64(f) != 0 {
		t.Fail()
	}
	if Int8(g) != 0 {
		t.Fail()
	}
	if Rune(h) != 0 {
		t.Fail()
	}
	if String(i) != "" {
		t.Fail()
	}
	if Uint(j) != 0 {
		t.Fail()
	}
	if Uint16(k) != 0 {
		t.Fail()
	}
	if Uint32(l) != 0 {
		t.Fail()
	}
	if Uint64(m) != 0 {
		t.Fail()
	}
	if Uint8(n) != 0 {
		t.Fail()
	}
}
