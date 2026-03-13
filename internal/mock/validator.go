package mock

import "golang-entrypoint/internal/service"

type ValidatorService struct{}

func (*ValidatorService) Validate(request interface{}) ([]service.ValidationError, error) {
	return nil, nil
}
