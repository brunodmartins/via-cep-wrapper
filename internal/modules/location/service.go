package location

import (
	"BrunoDM2943/via-cep-wrapper/internal/constants/domain"
	"BrunoDM2943/via-cep-wrapper/internal/gateway/viacep"
)

//go:generate mockgen -source=./service.go -destination=./mock/service_mock.go
type Service interface {
	SearchLocation(zipCode string) (*domain.Address, error)
}

type locationService struct {
	api viacep.Gateway
}

func NewLocationService(api viacep.Gateway) Service {
	return &locationService{api: api}
}

//SearchLocation search for a given zipcode
func (service *locationService) SearchLocation(zipCode string) (*domain.Address, error) {
	dto, err := service.api.GetLocation(zipCode)
	if err != nil {
		return nil, err
	}
	return dto.ToAddress(), nil
}
