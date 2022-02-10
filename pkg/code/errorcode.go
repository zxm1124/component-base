package code

import (
	"fmt"
	"github.com/novalagung/gubrak"
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
	if len(codes) == 0 {
		codes = make(map[int]string)
	}

	if _, ok := codes[code]; ok {
		logrus.Panicf("code：%d already exists, please replace it", code)
	}
	// 将该错误码存放到map中
	codes[code] = message

	found, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 500}, httpStatus)
	if !found {
		logrus.Panicf("type code is: %d, the code not in `200, 403, 400, 401, 404, 500`", httpStatus)
	}

	return &ErrorCode{
		code:       code,
		httpStatus: httpStatus,
		message:    message,
	}
}

func (e *ErrorCode) Error() string {
	if len(e.details) > 0 {
		str := e.details[0]
		for i := 1; i < len(e.details); i++ {
			str = fmt.Sprintf("%s.%s", str, e.details[i])
		}

		return fmt.Sprintf("code：%d，msg：%s, details: %s",
			e.code, e.message, e.details)
	}
	return fmt.Sprintf("code：%d，msg：%s", e.code, e.message)
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
