package λ

// λab.a
func True(a λ) λ {
	return func(b λ) λ {
		return a
	}
}

// λab.b
func False(a λ) λ {
	return func(b λ) λ {
		return b
	}
}

// λab.a(b)
func IfThenElse(a λ) λ {
	return func(b λ) λ {
		return a(b)
	}
}

// λa.a(λ bc.c)(λ de.d)
func Not(a λ) λ {
	return a(func(b λ) λ {
		return func(c λ) λ {
			return c
		}
	})(func(d λ) λ {
		return func(e λ) λ {
			return d
		}
	})
}

// λab.ab(λ xy.y)
func And(a λ) λ {
	return func(b λ) λ {
		return a(b)(False)
	}
}

// λab.a(λ xy.x)b
func Or(a λ) λ {
	return func(b λ) λ {
		return a(True)(b)
	}
}
