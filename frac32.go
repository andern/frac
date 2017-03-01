package frac

type Frac32 struct {
	Num uint32
	Den uint32
	Neg bool
}

func (a Frac32) Mul(b Frac32) Frac32 {
	neg := a.Neg || b.Neg
	if a.Neg && b.Neg {
		neg = false
	}
	return Frac32{
		Num: a.Num * b.Num,
		Den: a.Den * b.Den,
		Neg: neg,
	}
}

func (a Frac32) Div(b Frac32) Frac32 {
	neg := a.Neg || b.Neg
	if a.Neg && b.Neg {
		neg = false
	}
	return Frac32{
		Num: a.Num * b.Den,
		Den: a.Den * b.Num,
		Neg: neg,
	}
}

func (a Frac32) Add(b Frac32) Frac32 {
	if b.Neg {
		return a.Sub(b.Negate())
	}
	if a.Neg {
		return b.Sub(a.Negate())
	}

	lcm := lcm32(a.Den, b.Den)
	return Frac32{
		Num: a.Num*(lcm/a.Den) + b.Num*(lcm/b.Den),
		Den: lcm,
	}
}

func (a Frac32) Sub(b Frac32) Frac32 {
	if b.Neg {
		return a.Add(b.Negate())
	}
	if a.Neg {
		return a.Negate().Add(b).Negate()
	}

	lcm := lcm32(a.Den, b.Den)
	anum := a.Num * (lcm / a.Den)
	bnum := b.Num * (lcm / b.Den)
	return sub32(anum, bnum, lcm)
}

func (a Frac32) Negate() Frac32 {
	return Frac32{
		Num: a.Num,
		Den: a.Den,
		Neg: !a.Neg,
	}
}

func (a Frac32) Equals(b Frac32) bool {
	return a.Num == b.Num && a.Den == b.Den && a.Neg == b.Neg
}

func (a Frac32) Simplify() Frac32 {
	gcd := gcd32(a.Num, a.Den)
	return Frac32{
		Num: a.Num / gcd,
		Den: a.Den / gcd,
		Neg: a.Neg,
	}
}

func sub32(anum, bnum, den uint32) Frac32 {
	if bnum > anum {
		return Frac32{
			Num: bnum - anum,
			Den: den,
			Neg: true,
		}
	}
	return Frac32{
		Num: anum - bnum,
		Den: den,
	}
}

func gcd32(a, b uint32) uint32 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func lcm32(a, b uint32) uint32 {
	return a * b / gcd32(a, b)
}
