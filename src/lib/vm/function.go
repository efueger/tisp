package vm

type Function func(...*Thunk) Object

func (f Function) Call(ts ...*Thunk) Object {
	return f(ts...)
}

func NewLazyFunction(f func(...*Thunk) Object) func(...*Thunk) *Thunk {
	return func(ts ...*Thunk) *Thunk {
		return App(Normal(Function(f)), ts...)
	}
}

func NewStrictFunction(f func(...Object) Object) func(...*Thunk) *Thunk {
	return NewLazyFunction(func(ts ...*Thunk) Object {
		for _, t := range ts {
			go t.Eval()
		}

		os := make([]Object, len(ts))

		for i, t := range ts {
			os[i] = t.Eval()
		}

		return f(os...)
	})
}
