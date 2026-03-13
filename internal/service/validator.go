package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type ValidationError struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
	Param string      `json:"param"`
}

var (
	ErrNotAStruct = errors.New("not a struct")
)

type ValidatorService interface {
	Validate(request interface{}) ([]ValidationError, error)
}

type validatorService struct {
	validator *validator.Validate
}

func (v *validatorService) Validate(request interface{}) ([]ValidationError, error) {
	iValue := reflect.ValueOf(request)
	if reflect.Indirect(iValue).Kind() != reflect.Struct {
		return nil, ErrNotAStruct
	}

	err := v.validator.Struct(request)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return nil, err
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			result := make([]ValidationError, 0, len(validateErrs))
			for _, e := range validateErrs {
				result = append(result, ValidationError{
					Field: e.Field(),
					Tag:   e.Tag(),
					Value: e.Value(),
					Param: e.Param(),
				})
			}

			return result, nil
		}

		return nil, err
	}

	return nil, nil
}

func NewValidatorService(
	validator *validator.Validate,
) ValidatorService {
	return &validatorService{
		validator: validator,
	}
}
