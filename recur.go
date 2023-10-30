package λ

// This implementation is only a "candidate" - I haven't tested it yet

// Y = λf.(λx.f (x x))(λx.f (x x))
// func Y (f λ) λ {
// 	return func (x λ) λ {
// 		return f(x(x))
// 	}(func(x λ) λ {
// 		return f(x(x))
// 	})
// }

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
