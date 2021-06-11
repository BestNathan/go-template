package example

type ExampleApp struct {
}

func New() *ExampleApp {
	return &ExampleApp{}
}

func (e ExampleApp) Name() string {
	return "ExampleApp"
}

func (e *ExampleApp) Serve() error {
	return nil
}
