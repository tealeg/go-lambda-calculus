package λ


func Y (f λ) λ {
	return func(le λ) λ {
		return func(g λ) λ{
			return g(g)
		}(
			func(h λ) λ {
				return le(func(x λ) λ {
					return h(h)(x)
				})
			},
		)
	}
}


// F  = λ f. λ n. cond (isZero n) 1 (Mult n (f (Pred n))
func F(f λ) λ {
	return func(n λ) λ {
		return IfThenElse(IsZero(n))(one)(mul(n)(f(pred(n))))
	}
}
