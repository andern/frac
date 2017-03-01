package frac

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var add_t = [][]string{
	/* a + b = c */
	[]string{"1/2", "3/4", "5/4"},
	[]string{"1/2", "-3/4", "-1/4"},
	[]string{"-1/2", "3/4", "1/4"},
	[]string{"-1/2", "-3/4", "-5/4"},
	[]string{"3/4", "1/2", "5/4"},
	[]string{"3/4", "-1/2", "1/4"},
	[]string{"-3/4", "1/2", "-1/4"},
	[]string{"-3/4", "-1/2", "-5/4"},

	[]string{"123/321", "456/654", "12601/11663"},
	[]string{"123/321", "-456/654", "-3663/11663"},
	[]string{"-123/321", "456/654", "3663/11663"},
	[]string{"-123/321", "-456/654", "-12601/11663"},
	[]string{"456/654", "123/321", "12601/11663"},
	[]string{"456/654", "-123/321", "3663/11663"},
	[]string{"-456/654", "123/321", "-3663/11663"},
	[]string{"-456/654", "-123/321", "-12601/11663"},
}
var sub_t = [][]string{
	/* a - b = c */
	[]string{"1/2", "3/4", "-1/4"},
	[]string{"1/2", "-3/4", "5/4"},
	[]string{"-1/2", "3/4", "-5/4"},
	[]string{"-1/2", "-3/4", "1/4"},
	[]string{"3/4", "1/2", "1/4"},
	[]string{"3/4", "-1/2", "5/4"},
	[]string{"-3/4", "1/2", "-5/4"},
	[]string{"-3/4", "-1/2", "-1/4"},

	[]string{"123/321", "456/654", "-3663/11663"},
	[]string{"123/321", "-456/654", "12601/11663"},
	[]string{"-123/321", "456/654", "-12601/11663"},
	[]string{"-123/321", "-456/654", "3663/11663"},
	[]string{"456/654", "123/321", "3663/11663"},
	[]string{"456/654", "-123/321", "12601/11663"},
	[]string{"-456/654", "123/321", "-12601/11663"},
	[]string{"-456/654", "-123/321", "-3663/11663"},
}
var mul_t = [][]string{
	/* a * b = c */
	[]string{"1/2", "3/4", "3/8"},
	[]string{"1/2", "-3/4", "-3/8"},
	[]string{"-1/2", "3/4", "-3/8"},
	[]string{"-1/2", "-3/4", "3/8"},
	[]string{"3/4", "1/2", "3/8"},
	[]string{"3/4", "-1/2", "-3/8"},
	[]string{"-3/4", "1/2", "-3/8"},
	[]string{"-3/4", "-1/2", "3/8"},

	[]string{"123/321", "456/654", "3116/11663"},
	[]string{"123/321", "-456/654", "-3116/11663"},
	[]string{"-123/321", "456/654", "-3116/11663"},
	[]string{"-123/321", "-456/654", "3116/11663"},
	[]string{"456/654", "123/321", "3116/11663"},
	[]string{"456/654", "-123/321", "-3116/11663"},
	[]string{"-456/654", "123/321", "-3116/11663"},
	[]string{"-456/654", "-123/321", "3116/11663"},
}
var div_t = [][]string{
	/* a / b = c */
	[]string{"1/2", "3/4", "2/3"},
	[]string{"1/2", "-3/4", "-2/3"},
	[]string{"-1/2", "3/4", "-2/3"},
	[]string{"-1/2", "-3/4", "2/3"},
	[]string{"3/4", "1/2", "3/2"},
	[]string{"3/4", "-1/2", "-3/2"},
	[]string{"-3/4", "1/2", "-3/2"},
	[]string{"-3/4", "-1/2", "3/2"},

	[]string{"123/321", "456/654", "4469/8132"},
	[]string{"123/321", "-456/654", "-4469/8132"},
	[]string{"-123/321", "456/654", "-4469/8132"},
	[]string{"-123/321", "-456/654", "4469/8132"},
	[]string{"456/654", "123/321", "8132/4469"},
	[]string{"456/654", "-123/321", "-8132/4469"},
	[]string{"-456/654", "123/321", "-8132/4469"},
	[]string{"-456/654", "-123/321", "8132/4469"},
}

func (a Frac64) String() string {
	if a.Neg {
		return fmt.Sprintf("-%v/%v", a.Num, a.Den)
	}
	return fmt.Sprintf("%v/%v", a.Num, a.Den)
}

func TestAdd(t *testing.T) {
	for _, s := range add_t {
		a, b, c := readTest64(s)
		res := a.Add(b).Simplify()
		if !res.Equals(c) {
			fmt.Printf("%v + %v = %v\n", a, b, res)
			t.Fail()
		}
	}
}

func TestSub(t *testing.T) {
	for _, s := range sub_t {
		a, b, c := readTest64(s)
		res := a.Sub(b).Simplify()
		if !res.Equals(c) {
			fmt.Printf("%v - %v = %v\n", a, b, res)
			t.Fail()
		}
	}
}

func TestMul(t *testing.T) {
	for _, s := range mul_t {
		a, b, c := readTest64(s)
		res := a.Mul(b).Simplify()
		if !res.Equals(c) {
			fmt.Printf("%v * %v = %v\n", a, b, res)
			t.Fail()
		}
	}
}

func TestDiv(t *testing.T) {
	for _, s := range div_t {
		a, b, c := readTest64(s)
		res := a.Div(b).Simplify()
		if !res.Equals(c) {
			fmt.Printf("%v / %v = %v\n", a, b, res)
			t.Fail()
		}
	}
}

func readTest64(t []string) (a, b, c Frac64) {
	return readFrac64(t[0]), readFrac64(t[1]), readFrac64(t[2])
}

func readFrac64(f string) Frac64 {
	var neg bool
	var i int
	if f[0] == '-' {
		neg = true
		i = 1
	}
	slash := strings.IndexRune(f, '/')
	num, _ := strconv.Atoi(f[i:slash])
	den, _ := strconv.Atoi(f[slash+1:])
	return Frac64{
		Num: uint64(num),
		Den: uint64(den),
		Neg: neg,
	}
}
