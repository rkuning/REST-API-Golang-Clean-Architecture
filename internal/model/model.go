package model

type WebResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}
