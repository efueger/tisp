package core

// Signature represents function signature.
type Signature struct {
	positionals argumentSet
	keywords    argumentSet
}

// NewSignature defines a new Signature.
func NewSignature(
	pr []string, po []OptionalArgument, pp string,
	kr []string, ko []OptionalArgument, kk string) Signature {
	return Signature{
		positionals: argumentSet{
			requireds: pr,
			optionals: po,
			rest:      pp,
		},
		keywords: argumentSet{
			requireds: kr,
			optionals: ko,
			rest:      kk,
		},
	}
}

func NewSimpleSignature(pr ...string) Signature {
	return NewSignature(
		pr, []OptionalArgument{}, "",
		[]string{}, []OptionalArgument{}, "",
	)
}

// Bind binds Arguments to names defined in Signature and returns full
// arguments to be passed to a function.
func (s Signature) Bind(args Arguments) ([]*Thunk, *Thunk) {
	ts := make([]*Thunk, 0, s.arity())

	for _, name := range s.positionals.requireds {
		t := args.searchKeyword(name)

		if t == nil {
			t = args.nextPositional()
		}

		if t == nil {
			return nil, argumentError("Could not bind an required positional argument.")
		}

		ts = append(ts, t)
	}

	for _, o := range s.positionals.optionals {
		t := args.searchKeyword(o.name)

		if t == nil {
			t = args.nextPositional()
		}

		if t == nil {
			t = o.defaultValue
		}

		ts = append(ts, t)
	}

	if s.positionals.rest != "" {
		t := args.searchKeyword(s.positionals.rest)

		if t == nil {
			t = args.restPositionals()
		}

		ts = append(ts, t)
	}

	for _, name := range s.keywords.requireds {
		t := args.searchKeyword(name)

		if t == nil {
			return nil, argumentError("Could not bind an required positional argument.")
		}

		ts = append(ts, t)
	}

	for _, o := range s.keywords.optionals {
		t := args.searchKeyword(o.name)

		if t == nil {
			t = o.defaultValue
		}

		ts = append(ts, t)
	}

	if s.keywords.rest != "" {
		t := args.searchKeyword(s.keywords.rest)

		if t == nil {
			t = args.restKeywords()
		}

		ts = append(ts, t)
	}

	if len(ts) != s.arity() {
		return nil, argumentError("Number of arguments bound to names is different from signature's arity.")
	}

	return ts, nil
}

func (s Signature) arity() int {
	return s.positionals.size() + s.keywords.size()
}

func argumentError(m string) *Thunk {
	return NewError("ArgumentError", m)
}
