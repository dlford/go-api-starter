package models

type ResData[T any] struct {
	Data T `json:"data"`
}

type ResMsg struct {
	Message string `json:"message" example:"Success message here..."`
}

type ResErr struct {
	Error string `json:"error" example:"Error message here..."`
}