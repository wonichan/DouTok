package launcher

type Option func(*Launcher)

func WithConfig(c interface{}) Option {
	return func(l *Launcher) {
		l.configValue = c
	}
}
