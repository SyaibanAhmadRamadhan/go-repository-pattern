package helpers

type Response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponse(code int, message string, data any) any {
	status := "ok"
	if code >= 400 {
		status = "failed"
	}
	if data != nil {
		return &Response{
			Status:  status,
			Code:    code,
			Message: message,
			Data:    data,
		}
	} else {
		return &ResponseWithoutData{
			Status:  status,
			Code:    code,
			Message: message,
		}
	}
}
