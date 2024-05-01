package enums

type LogLevel int

const (
	Info LogLevel = iota
	Alert
	Debug
	Warning
	Error
	Fatal
)

func (r LogLevel) String() string {
	return [...]string{"info", "alert", "debug", "warning", "error", "fatal"}[r]
}
