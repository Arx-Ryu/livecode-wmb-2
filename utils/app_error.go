package utils

import (
	"fmt" 
	"net/http"
)

type AppError struct {
	ErrorCode 		string
	ErrorMessage 	string
	ErrorType		int
}

func (e AppError) Error() string {
	return fmt.Sprintf("type: %d, code: %s, err: %s", e.ErrorType, e.ErrorCode, e.ErrorMessage)
}

func RequiredError() error {
	return AppError{
		ErrorCode: "X01",
		ErrorMessage: "Input can't be Empty",
		ErrorType: http.StatusBadRequest,
	}
}

func UnauthorizedError() error {
	return AppError{
		ErrorCode: "X04",
		ErrorMessage: "Unauthorized User",
		ErrorType: http.StatusUnauthorized,
	}
}

func DataNotFoundError() error {
	return AppError{
		ErrorCode: "X02",
		ErrorMessage: "No Data Found",
	}
}

