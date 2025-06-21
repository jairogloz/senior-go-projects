package dto

type UserCreateParams struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required,gte=18"`
}
