package app

type Logger interface {
	Printf(format string, values ...interface{})
}
