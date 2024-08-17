package trigger

type Handler interface {
	Passed() bool
}
