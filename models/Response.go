package models

// base response for rest api
type Response struct {
	StatusCode string      `json:"statusCode"`
	Message    string      `json:"message"`
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
}

// base response creator  for success
func ResponseSuccess(data interface{}) *Response {
	response := new(Response)
	response.Data = data
	response.Message = "OK"
	response.StatusCode = "200"
	response.Success = true
	return response
}

// base response creator for failure
func ResponseFailure(data interface{}, message string) *Response {
	response := new(Response)
	response.Data = data
	response.Message = message
	response.StatusCode = "400"
	response.Success = false
	return response
}

