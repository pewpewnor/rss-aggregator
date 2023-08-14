package response

type successResponse struct {
	Status  string                 `json:"status"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

func GenerateSimpleSuccessResponse(message string) successResponse {
	return successResponse{
		Status:  "success",
		Message: message,
	}
}

func GenerateSuccessResponse(data map[string]interface{}, message string) successResponse {
	return successResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}
