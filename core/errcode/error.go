package errcode

import "log"

// 内存保存code msg的引用
// 同时Error和code、msg保持对应的关系
// 实际返回的时候返回Error
//
var codes = map[int] string {}

// Error 错误码
type Error struct {
	Code int
	Msg string
	details [] string	
}


// NewError 错误码
func NewError(code int, message string) (* Error) {
	if _, ok :=codes[code]; ok {
		log.Fatalf("错误码 %d 已经存在", code)
	}
	codes[code] = message
	return &Error{
		Code: code,
		Msg: message,
	}
}

// StatusCode 获取Http状态码
func (e *Error) StatusCode(code int) int {
	return 1
}