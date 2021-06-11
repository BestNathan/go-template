package example2

type Example2App struct {
}

func New() *Example2App {
	return &Example2App{}
}

func (e Example2App) Name() string {
	return "Example2App"
}

func (e *Example2App) Serve() error {
	return nil
}
