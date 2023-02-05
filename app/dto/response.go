package dto

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}
