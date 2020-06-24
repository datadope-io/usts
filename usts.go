package usts

var (
	ilog Logger
)

func init() {

	ilog = createLogger()

}

func SetLogger(l Logger) {
	ilog = l
}
