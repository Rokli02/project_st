package model

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type StandardResponse struct {
	Error    *ResponseError `json:"error"`
	Response any            `json:"response"`
}
