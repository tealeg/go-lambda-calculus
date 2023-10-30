package λ

import (
	qt "github.com/frankban/quicktest"
	"testing"
)

func TestFact(t *testing.T) {
	c := qt.New(t)

	// Y returns the fixed point of F, which is the recursive factorial function
	// fact := Y(F)
	fact := Y(F)
	c.Assert(fact, qt.IsNotNil)

	// factorial of 0 is 1 (See here if you don't get why:
	// https://www.thoughtco.com/why-does-zero-factorial-equal-one-3126598
	// )
	c.Assert(intResult(fact(zero)), qt.Equals, 1)

	// factorial of 1 is 1
	// This code will cause the program to panic - here we reach the limi t for Go
	// c.Assert(intResult(fact(one)), qt.Equals, 1)

	
}
