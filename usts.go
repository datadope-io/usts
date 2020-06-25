package usts

var (
	ilog Logger
)

func init() {

	ilog = createLogger()

}

// SetLogger add custom logger enables custom format, output and level, if not set
// all logs will be writen ove Stdout with without level filtering (all logs will be shown)
func SetLogger(l Logger) {
	ilog = l
}
