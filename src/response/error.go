package response

import "strconv"

type errorResponse struct {
	Error errorResponseContent `json:"error"`
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

func SimpleErrorResponseFromError(err error) errorResponse {
	return errorResponse{
		Error: errorResponseContent{
			Message: err.Error(),
		},
	}
}

func ErrorResponse(code int, message string, details string, validationErrors []ErrorResponseValidation) errorResponse {
	return errorResponse{
		Error: errorResponseContent{
			strconv.Itoa(code),
			message,
			details,
			validationErrors,
		},
	}
}
