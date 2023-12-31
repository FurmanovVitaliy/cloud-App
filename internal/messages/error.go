package messages

import (
	"encoding/json"
	"fmt"
)

var (
	ErrNotFound = NewAppError(nil, "not found", "", "US-000001")

	ErrStartGame           = NewAppError(nil, "error while starting game", "", "WO-000001")
	ErrStartCapture        = NewAppError(nil, "error while starting capture", "", "WO-000002")
	ErrCreateVirtualDevice = NewAppError(nil, "error while creating virtual device", "", "WO-000003")
)

type AppError struct {
	Err             error  `json:"-"`
	Massage         string `json:"massage,omitempty"`
	DevelopeMassage string `json:"develope_massafe,omitempty"`
	Code            string `json:"code,omitempty"`
}

func (e *AppError) Error() string { return e.Massage }

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}
func NewAppError(err error, massage, developeMassage, code string) *AppError {
	return &AppError{
		Err:             fmt.Errorf(massage),
		Massage:         massage,
		DevelopeMassage: developeMassage,
		Code:            code,
	}
}

func wrapSystemError(err error) *AppError {
	return NewAppError(err, "internal system error", err.Error(), "US-000000")
}
