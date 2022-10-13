package λ

import (
	qt "github.com/frankban/quicktest"
	"testing"
)

// makeCounter creates a closure around an integer that is utilised by
// the 3 functions it returns: an incrementor, a getter and a
// resetter.  The incrementor is a λ function, and thus can be used as
// a parameter to, or return value from any other λ type.
// Specifically if you pass it to a church numeral, it will be called
// the number of times that numeral represents, and thus can be used
// for its side-effect of generating an Integer equivalent of any
// given Church numeral.
func makeCounter() (λ, func() int, func()) {
	var i int = 0

	inc := func(f λ) λ {
		i = i + 1
		return f
	}

	get := func() int {
		return i
	}

	reset := func() {
		i = 0
	}

	return inc, get, reset
}

// intResult is a convenience method used in testing.  It wraps up the
// steps required to convert a Church numeral to a Go integer.
func intResult(l λ) int {
	counter, count, _ := makeCounter()
	_ = l(counter)(nil)
	return count()
}

func TestZero(t *testing.T) {
	c := qt.New(t)

	c.Assert(intResult(zero), qt.Equals, 0)
}

func TestOne(t *testing.T) {
	c := qt.New(t)

	c.Assert(intResult(one), qt.Equals, 1)
}

func TestTwo(t *testing.T) {
	c := qt.New(t)

	c.Assert(intResult(two), qt.Equals, 2)
}

func TestSucc(t *testing.T) {
	c := qt.New(t)

	// The succesor of zero is 1
	c.Assert(intResult(succ(zero)), qt.Equals, 1)
	c.Assert(intResult(zero(succ)), qt.Equals, 1)

	// The successor of one is 2
	c.Assert(intResult(succ(one)), qt.Equals, 2)

	// The 0th successor of 0 is 0
	c.Assert(intResult(zero(succ)(zero)), qt.Equals, 0)

	// The 0th succesor of one is 0
	c.Assert(intResult(zero(succ)(one)), qt.Equals, 1)

	// The 1st succesor of 0 is 1
	c.Assert(intResult(one(succ)(zero)), qt.Equals, 1)

	// The 1st succesor of 1 is 2
	c.Assert(intResult(one(succ)(one)), qt.Equals, 2)

	// The 3rd succesor of 3 is 6
	three := succ(two)
	c.Assert(intResult(three), qt.Equals, 3)
	c.Assert(intResult(three(succ)(three)), qt.Equals, 6)
}

func TestPlus(t *testing.T) {
	c := qt.New(t)

	c.Assert(intResult(one(plus)(one)), qt.Equals, 2)
	c.Assert(intResult(two(plus)(two)), qt.Equals, 4)
}

func TestPowers(t *testing.T) {
	c := qt.New(t)

	// 0**1
	c.Assert(intResult(one(zero)), qt.Equals, 0)
	// 1**0
	c.Assert(intResult(zero(one)), qt.Equals, 1)
	// 1**1
	c.Assert(intResult(one(one)), qt.Equals, 1)
	// 1**2
	c.Assert(intResult(two(one)), qt.Equals, 1)
	// 1**2
	c.Assert(intResult(one(two)), qt.Equals, 2)
	// 2**2
	c.Assert(intResult(two(two)), qt.Equals, 4)

}

func TestMul(t *testing.T) {
	c := qt.New(t)

	four := mul(two)(two)
	c.Assert(intResult(four), qt.Equals, 4)

	eight := mul(four)(two)
	c.Assert(intResult(eight), qt.Equals, 8)

}

func TestIsZero(t *testing.T) {
	c := qt.New(t)

	b := IsZero(zero)
	c.Assert(intResult(IfThenElse(b)(one)(two)), qt.Equals, 1)
	b = IsZero(one)
	c.Assert(intResult(IfThenElse(b)(one)(two)), qt.Equals, 2)
}
