package validator_models

type UserGuard struct {
	Name string `json:"name" validate:"required,min=1,max=20,alphanum"`
	Age  int    `json:"age" validate:"gte=0,lte=10"`
}
