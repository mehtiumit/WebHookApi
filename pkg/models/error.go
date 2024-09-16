package models

type CustomError struct {
	Code        int    `json:"code"`
	ErrorDetail string `json:"errorDetail"`
}

func (e CustomError) Error() string {
	return e.ErrorDetail
}
