package location

import (
	"BrunoDM2943/via-cep-wrapper/internal/constants/domain"
	"BrunoDM2943/via-cep-wrapper/internal/gateway/viacep"
)

func SearchLocation(zipCode string) (*domain.Address, error) {
	dto, err := viacep.GetLocation(zipCode)
	if err != nil {
		return nil, err
	}
	return dto.ToAddress(), nil
}
