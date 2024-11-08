package result

type ErrorCode uint8

const (
	Success ErrorCode = iota
	DBError
	LoggerError
)
