package fibonacci

import (
	"math/big"
	"reflect"
	"testing"
)

func TestSequenceNormal(t *testing.T) {
	var tests = []struct {
		in   int
		want []*big.Int
	}{
		{0, []*big.Int{}},
		{1, []*big.Int{big.NewInt(0)}},
		{2, []*big.Int{big.NewInt(0), big.NewInt(1)}},
		{3, []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(1)}},
		{6, []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(5)}},
	}

	for _, c := range tests {
		var got = Sequence(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Error("Sequence(%q) -> %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSequenceLarge(t *testing.T) {
	var s = Sequence(100)
	var strings = []string{
		s[97].String(),
		s[98].String(),
		s[99].String(),
	}
	var want = []string{
		"83621143489848422977",
		"135301852344706746049",
		"218922995834555169026",
	}

	if !reflect.DeepEqual(strings, want) {
		t.Error("Last 3 elements of Sequence(100) -> %q, want %q",
			strings, want)
	}
}

func TestSequenceXlarge(t *testing.T) {
	var s = Sequence(100000)
	var s_str = s[99999].String()
	const length = 20899

	if len(s_str) != length {
		t.Error("length mismatch: got %q, want %q", len(s_str), length)
	}
	if s_str[:32] != "16052857682729819697035016991663" {
		t.Error("First 32 bytes mismatch")
	}
	if s_str[length-32:] != "35545120747688390605016278790626" {
		t.Error("Last 32 bytes mismatch")
	}
}
