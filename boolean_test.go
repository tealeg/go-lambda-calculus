package λ

import (
	qt "github.com/frankban/quicktest"
	"testing"
)

func TestIfThenElse(t *testing.T) {
	c := qt.New(t)

	result := IfThenElse(True)(one)(two)
	c.Assert(intResult(result), qt.Equals, 1)

	result = IfThenElse(False)(one)(two)
	c.Assert(intResult(result), qt.Equals, 2)
}

func TestNot(t *testing.T) {
	c := qt.New(t)

	c.Assert(intResult(True(one)(two)), qt.Equals, 1)
	c.Assert(intResult(Not(True)(one)(two)), qt.Equals, 2)
	c.Assert(intResult(False(one)(two)), qt.Equals, 2)
	c.Assert(intResult(Not(False)(one)(two)), qt.Equals, 1)
}

func TestAnd(t *testing.T) {
	c := qt.New(t)

	b := And(True)(True)
	c.Assert(intResult(IfThenElse(b)(one)(two)), qt.Equals, 1)
	b = And(False)(True)
	c.Assert(intResult(IfThenElse(b)(one)(two)), qt.Equals, 2)
	b = And(False)(False)
	c.Assert(intResult(IfThenElse(b)(one)(two)), qt.Equals, 2)
	b = And(True)(False)
	c.Assert(intResult(IfThenElse(b)(one)(two)), qt.Equals, 2)

}

func TestOr(t *testing.T) {
	c := qt.New(t)

	b := Or(True)(True)
	c.Assert(intResult(IfThenElse(b)(one)(two)), qt.Equals, 1)
	b = Or(False)(True)
	c.Assert(intResult(IfThenElse(b)(one)(two)), qt.Equals, 1)
	b = Or(True)(False)
	c.Assert(intResult(IfThenElse(b)(one)(two)), qt.Equals, 1)
	b = Or(False)(False)
	c.Assert(intResult(IfThenElse(b)(one)(two)), qt.Equals, 2)

}
