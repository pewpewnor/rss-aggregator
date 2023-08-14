package res

import (
	"strconv"
)

type errorResponse struct {
	ErrorData errorResponseContent `json:"error"`
}

func (e *errorResponse) AddValidation(validation ErrorResponseValidation) {
	e.ErrorData.ValidationErrors = append(e.ErrorData.ValidationErrors, validation)
}

func (e errorResponse) Error() string {
	return e.ErrorData.Message
}

type errorResponseContent struct {
	Code             string                    `json:"code"`
	Message          string                    `json:"message"`
	Details          string                    `json:"details"`
	ValidationErrors []ErrorResponseValidation `json:"validationErrors"`
}

type ErrorResponseValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func SimpleErrorResponse(message string) errorResponse {
	return errorResponse{
		ErrorData: errorResponseContent{
			Message: message,
		},
	}
}

func SimpleErrorResponseFromError(err error) errorResponse {
	return errorResponse{
		ErrorData: errorResponseContent{
			Message: err.Error(),
		},
	}
}

func ErrorResponse(code int, message string, details string, validationErrors []ErrorResponseValidation) errorResponse {
	return errorResponse{
		ErrorData: errorResponseContent{
			strconv.Itoa(code),
			message,
			details,
			validationErrors,
		},
	}
}
