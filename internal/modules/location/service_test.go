package location

import (
	"BrunoDM2943/via-cep-wrapper/internal/gateway/viacep"
	mock_viacep "BrunoDM2943/via-cep-wrapper/internal/gateway/viacep/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const zipCode = "01001-000"

func TestSearchLocation_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_viacep.NewMockGateway(ctrl)
	service := NewLocationService(apiMock)

	apiMock.EXPECT().GetLocation(gomock.Eq(zipCode)).Return(&viacep.ViaCepResponse{Cep: "01001-000"}, nil)

	result, err := service.SearchLocation(zipCode)
	assert.Equal(t, result.ZipCode, zipCode)
	assert.Nil(t, err)
}

func TestSearchLocation_NOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_viacep.NewMockGateway(ctrl)
	service := NewLocationService(apiMock)

	apiMock.EXPECT().GetLocation(gomock.Eq(zipCode)).Return(nil, errors.New("generic error"))


	_, err := service.SearchLocation(zipCode)
	assert.NotNil(t, err)
}


