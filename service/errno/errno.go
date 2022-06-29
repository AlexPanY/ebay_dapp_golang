package errno

import "fmt"

type Err struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
	RawErr  error  `json:"-"`
}

//NewErr
func NewErr(code string, msg string, err error) Err {
	return Err{
		Code:    code,
		Message: msg,
		RawErr:  err,
	}
}

//WithRawErr
func (e *Err) WithRawErr(err error) Err {
	return Err{
		Code:    e.Code,
		Message: e.Message,
		RawErr:  err,
	}
}

//WithFmtAndRawErr
func (e *Err) WithFmtAndRawErr(f string, err error) Err {
	errs := Err{
		Code:    e.Code,
		Message: e.Message,
		RawErr:  err,
	}
	errs.Message = fmt.Sprintf(errs.Message, f)

	return errs
}
