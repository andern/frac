package frac

type Frac64 struct {
	Num uint64
	Den uint64
	Neg bool
}

func (a Frac64) Mul(b Frac64) Frac64 {
	neg := a.Neg != b.Neg
	return Frac64{
		Num: a.Num * b.Num,
		Den: a.Den * b.Den,
		Neg: neg,
	}
}

func (a Frac64) Div(b Frac64) Frac64 {
	neg := a.Neg != b.Neg
	return Frac64{
		Num: a.Num * b.Den,
		Den: a.Den * b.Num,
		Neg: neg,
	}
}

func (a Frac64) Add(b Frac64) Frac64 {
	if b.Neg {
		return a.Sub(b.Negate())
	}
	if a.Neg {
		return b.Sub(a.Negate())
	}

	lcm := lcm64(a.Den, b.Den)
	return Frac64{
		Num: a.Num*(lcm/a.Den) + b.Num*(lcm/b.Den),
		Den: lcm,
	}
}

func (a Frac64) Sub(b Frac64) Frac64 {
	if b.Neg {
		return a.Add(b.Negate())
	}
	if a.Neg {
		return a.Negate().Add(b).Negate()
	}

	lcm := lcm64(a.Den, b.Den)
	anum := a.Num * (lcm / a.Den)
	bnum := b.Num * (lcm / b.Den)
	return sub64(anum, bnum, lcm)
}

func (a Frac64) Negate() Frac64 {
	return Frac64{
		Num: a.Num,
		Den: a.Den,
		Neg: !a.Neg,
	}
}

func (a Frac64) Equals(b Frac64) bool {
	return a.Num == b.Num && a.Den == b.Den && a.Neg == b.Neg
}

func (a Frac64) Simplify() Frac64 {
	gcd := gcd64(a.Num, a.Den)
	return Frac64{
		Num: a.Num / gcd,
		Den: a.Den / gcd,
		Neg: a.Neg,
	}
}

func sub64(anum, bnum, den uint64) Frac64 {
	if bnum > anum {
		return Frac64{
			Num: bnum - anum,
			Den: den,
			Neg: true,
		}
	}
	return Frac64{
		Num: anum - bnum,
		Den: den,
	}
}

func gcd64(a, b uint64) uint64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func lcm64(a, b uint64) uint64 {
	return a * b / gcd64(a, b)
}
