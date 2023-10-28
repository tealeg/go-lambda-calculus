package λ

// λ sz.z
func zero(s λ) λ {
	return func(z λ) λ {
		return z
	}
}

// λ sz.sz
func one(s λ) λ {
	return func(z λ) λ {
		return s(z)
	}
}

// λ sz.ssz
func two(s λ) λ {
	return func(z λ) λ {
		return s(s(z))
	}
}

// λ abc.b(abc)
func succ(a λ) λ {
	return func(b λ) λ {
		return func(c λ) λ {
			return b(a(b)(c))
		}
	}
}

// pred = λn.λf.λx.n (λg.λh.h (g f)) (λu.x) (λu.u)
func pred(n λ) λ {
	return func(f λ) λ {
		return func(x λ) λ {
			return n(
				func(g λ) λ {
					return func(h λ) λ {
						return h(g(f))
					}
				},
			)(func(u λ) λ {
				return x
			})(func(u λ) λ {
				return u
			})
		}
	}
}

// λ abc.b(abc)  - identical to succ
func plus(a λ) λ {
	return func(b λ) λ {
		return func(c λ) λ {
			return b(a(b)(c))
		}
	}
}

// λ abc.a(bc)
func mul(a λ) λ {
	return func(b λ) λ {
		return func(c λ) λ {
			return a(b(c))
		}
	}
}

// λa.a FALSE NOT FALSE
func IsZero(a λ) λ {
	return a(False)(Not)(False)
}
