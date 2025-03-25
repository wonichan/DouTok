package launcher

type Launcher struct {
}

func New(options ...Option) *Launcher {
	launcher := &Launcher{}
	for _, option := range options {
		option(launcher)
	}
	return launcher
}

func (l *Launcher) Run() {
	// TODO: implement launcher
}
