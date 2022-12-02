package dto

type SuccessResult struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
