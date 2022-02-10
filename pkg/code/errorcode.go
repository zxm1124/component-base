package code

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var codes map[int]string

type ErrorCode struct {
	code int

	httpStatus int

	message string

	details []string
}

func NewCode(code, httpStatus int, message string) *ErrorCode {
	if _, ok := codes[code]; ok {
		logrus.Panicf("code：%d already exists, please replace it", code)
	}
	// 将该错误码存放到map中
	codes[code] = message

	httpCode := []int{200, 403, 400, 401, 404, 500}
	for i := 0; i < len(httpCode); i++ {
		if httpCode[i] != httpStatus {
			logrus.Panicf("http: %d code not in `200, 403, 400, 401, 404, 500`", httpCode)
		}
	}

	return &ErrorCode{
		code:       code,
		httpStatus: httpStatus,
		message:    message,
	}
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.code, e.message)
}
func (e *ErrorCode) Code() int {
	return e.code
}
func (e *ErrorCode) Msg() string {
	return e.message
}
func (e *ErrorCode) Details() []string {
	return e.details
}
func (e *ErrorCode) HttpStatus() int {
	return e.httpStatus
}
func (e *ErrorCode) WithDetails(details ...string) *ErrorCode {
	for _, v := range details {
		e.details = append(e.details, v)
	}

	return e
}
