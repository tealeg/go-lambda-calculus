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
