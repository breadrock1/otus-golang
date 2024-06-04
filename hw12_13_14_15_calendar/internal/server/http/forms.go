package internalhttp

import "time"

// DatetimeForm example.
type DatetimeForm struct {
	Datetime *time.Time `json:"datetime" example:"2024-05-10T10:07:35Z"`
}

// ResponseForm example.
type ResponseForm struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"Done"`
}

// BadRequestForm example.
type BadRequestForm struct {
	Status  int    `json:"status" example:"400"`
	Message string `json:"message" example:"Bad Request message"`
}

// ServerErrorForm example.
type ServerErrorForm struct {
	Status  int    `json:"status" example:"503"`
	Message string `json:"message" example:"Server Error message"`
}

func createStatusResponse(status int, msg string) *ResponseForm {
	return &ResponseForm{Status: status, Message: msg}
}
