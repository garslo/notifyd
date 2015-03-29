package runners

type Runner interface {
	RunNext() []error
	RunForever()
}
