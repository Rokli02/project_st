package model

type ResponseError struct {
	Code    int
	Message string
}

type StandardResponse[T any] struct {
	Error    *ResponseError
	Response T
}
