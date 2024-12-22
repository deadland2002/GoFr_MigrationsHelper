package validator_models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gofr.dev/pkg/gofr"
)

var validate = validator.New()

func BindAndValidate(c *gofr.Context, v interface{}) error {
	if err := c.Request.Bind(v); err != nil {
		return fmt.Errorf("binding failed: %v", err)
	}

	fmt.Println("BindAndValidate", v)
	if err := validate.Struct(v); err != nil {
		return fmt.Errorf("validation failed: %v", err)
	}

	return nil
}

func GetUser(c *gofr.Context) (interface{}, error) {
	var data UserGuard
	if err := BindAndValidate(c, &data); err != nil {
		return nil, err
	}
	return data, nil
}
