package λ

// This implementation is only a "candidate" - I haven't tested it yet

// Y = λf.(λx.f (x x))(λx.f (x x))
func Y(f λ) λ {
	return func(x λ) λ {
		return f(x(x))
	}(func(x λ) λ {
		return f(x(x))
	})
}
