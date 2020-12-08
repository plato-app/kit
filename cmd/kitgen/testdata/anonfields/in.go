package foo

// from https://github.com/plato-app/kit/pull/589#issuecomment-319937530
type Service interface {
	Foo(context.Context, int, string) (int, error)
}
