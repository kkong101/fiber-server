package request

type UserRequest struct {
	Phone string `json:"phone" validate:"required"`
}
