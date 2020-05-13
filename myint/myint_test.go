package myint

import (
	"testing"
)

// Add is adding to Myint.
func TestAdd(t *testing.T) {
	data := []struct {
		title  string
		in     MyInt
		param  int
		should MyInt
	}{
		{"a", 1, 1, 2},
		{"b", 2, 2, 4},
	}
	for _, d := range data {
		d.in.Add(d.param)
		if d.in != d.should {
			t.Errorf("fo test %v got %v but should got %v", d.title, d.in, d.should)
		}
	}
}

func TestSub(t *testing.T) {
	data := []struct {
		title  string
		in     MyInt
		param  int
		should MyInt
	}{
		{"a", 1, 1, 0},
	}
	for _, d := range data {
		got := d.in.Sub(d.param)
		if got != d.should {
			t.Errorf("fo test %v got %v but should got %v", d.title, got, d.should)
		}
	}
}

func TestDivide(t *testing.T) {
	data := []struct {
		title  string
		in     MyInt
		param  int
		should MyInt
		err    error
	}{
		{"a", 1, 1, 1, nil},
		{"b", 1, 0, 1, ErrorDivideByZero},
	}
	for _, d := range data {
		got, err := d.in.Divide(d.param)
		if d.err != err {
			t.Errorf("fo test %v for error got %v but should got %v", d.title, err, d.err)
		}
		if got != d.should {
			t.Errorf("fo test %v got %v but should got %v", d.title, got, d.should)
		}
	}
}

func TestMultiply(t *testing.T) {
	data := []struct {
		title  string
		in     MyInt
		param  int
		should MyInt
	}{
		{"a", 2, 2, 4},
	}
	for _, d := range data {
		got := d.in.Multiply(d.param)
		if got != d.should {
			t.Errorf("fo test %v got %v but should got %v", d.title, got, d.should)
		}
	}
}
