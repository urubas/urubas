package runtime

type Type interface {
	Name() string
}

var Types map[string]Type

func RegisterType(t Type) {
	Types[t.Name()] = t
}

func FindType(name string) (Type, bool) {
	t, ok := Types[name]

	return t, ok
}
