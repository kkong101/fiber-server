package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func ValidateStruct(input any) error {
	return validate.Struct(input)
}

func ParseBody(c *fiber.Ctx, body any) error {
	if err := c.BodyParser(body); err != nil {
		return err
	}
	return nil
}

func ParseAndValidate(c *fiber.Ctx, body any) error {
	v := reflect.ValueOf(body)

	switch v.Kind() {
	case reflect.Ptr:
		ParseBody(c, body)

		return ValidateStruct(v.Elem().Interface())
	case reflect.Struct:
		ParseBody(c, &body)

		return ValidateStruct(v)
	default:
		return nil
	}
}
