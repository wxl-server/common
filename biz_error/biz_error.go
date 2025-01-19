package biz_error

type BizError struct {
	code    int64
	message string
	extra   []string
}

func (e *BizError) Error() string {
	return e.message
}

func newBizError(code int64, message string) *BizError {
	return &BizError{
		code:    code,
		message: message,
	}
}

func NewBizError(baseBizErr BizError, extra ...string) *BizError {
	baseBizErr.extra = extra
	return &baseBizErr
}

var (
	SystemError = newBizError(1000, "system error")
	DBError     = newBizError(1001, "db error")
	CacheError  = newBizError(1002, "cache error")
	MQError     = newBizError(1003, "mq error")
	ESError     = newBizError(1004, "es error")
)
