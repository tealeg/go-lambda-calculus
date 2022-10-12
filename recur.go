package λ

// Y = λf.(λx.f (x x))(λx.f (x x))
func Y(f λ) λ {
	return func(x λ) λ {
		return f(x(x))
	}(func(x λ) λ {
		return f(x(x))
	})
}

// func fib(f λ) λ {
// 	return func(x λ) λ {
// 		if x <= 2 {
// 			return 1
// 		}
// 		return f(x-1)+f(x-2)
// 	}
// }
