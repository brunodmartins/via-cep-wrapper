package location

import (
	"BrunoDM2943/via-cep-wrapper/internal/constants/domain"
	mock_location "BrunoDM2943/via-cep-wrapper/internal/modules/location/mock"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getCep_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock_location.NewMockService(ctrl)
	app := fiber.New()
	NewHandler(service).SetUpRoutes(app)

	service.EXPECT().SearchLocation(gomock.Eq(zipCode)).Return(&domain.Address{ZipCode: zipCode}, nil)

	req := httptest.NewRequest("GET", fmt.Sprintf("/via_cep_wrapper/%s", zipCode), nil)
	resp, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	result := &domain.Address{}
	_ = json.Unmarshal(body, result)
	assert.Equal(t, zipCode, result.ZipCode)

}

func Test_getCep_NOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock_location.NewMockService(ctrl)
	app := fiber.New()
	NewHandler(service).SetUpRoutes(app)

	service.EXPECT().SearchLocation(gomock.Eq(zipCode)).Return(nil, errors.New("generic error"))
	req := httptest.NewRequest("GET", fmt.Sprintf("/via_cep_wrapper/%s", zipCode), nil)
	resp, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

