package frac

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func (a Frac32) String() string {
	if a.Neg {
		return fmt.Sprintf("-%v/%v", a.Num, a.Den)
	}
	return fmt.Sprintf("%v/%v", a.Num, a.Den)
}

func TestAdd32(t *testing.T) {
	for _, s := range add_t {
		a, b, c := readTest32(s)
		res := a.Add(b).Simplify()
		if !res.Equals(c) {
			fmt.Printf("%v + %v = %v\n", a, b, res)
			t.Fail()
		}
	}
}

func TestSub32(t *testing.T) {
	for _, s := range sub_t {
		a, b, c := readTest32(s)
		res := a.Sub(b).Simplify()
		if !res.Equals(c) {
			fmt.Printf("%v - %v = %v\n", a, b, res)
			t.Fail()
		}
	}
}

func TestMul32(t *testing.T) {
	for _, s := range mul_t {
		a, b, c := readTest32(s)
		res := a.Mul(b).Simplify()
		if !res.Equals(c) {
			fmt.Printf("%v * %v = %v\n", a, b, res)
			t.Fail()
		}
	}
}

func TestDiv32(t *testing.T) {
	for _, s := range div_t {
		a, b, c := readTest32(s)
		res := a.Div(b).Simplify()
		if !res.Equals(c) {
			fmt.Printf("%v / %v = %v\n", a, b, res)
			t.Fail()
		}
	}
}

func readTest32(t []string) (a, b, c Frac32) {
	return readFrac32(t[0]), readFrac32(t[1]), readFrac32(t[2])
}

func readFrac32(f string) Frac32 {
	var neg bool
	var i int
	if f[0] == '-' {
		neg = true
		i = 1
	}
	slash := strings.IndexRune(f, '/')
	num, _ := strconv.Atoi(f[i:slash])
	den, _ := strconv.Atoi(f[slash+1:])
	return Frac32{
		Num: uint32(num),
		Den: uint32(den),
		Neg: neg,
	}
}
