package interfaces

type Logger interface {
	Error(message string, err error)
	Info(message string, fields ...interface{})
}
