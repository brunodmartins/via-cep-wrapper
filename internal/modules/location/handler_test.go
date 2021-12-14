package location

import (
	"BrunoDM2943/via-cep-wrapper/internal/constants/domain"
	"BrunoDM2943/via-cep-wrapper/internal/gateway/viacep"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getCep_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		responseBody := `
		{
			"cep": "01001-000",
			"logradouro": "Praça da Sé",
			"complemento": "lado ímpar",
			"bairro": "Sé",
			"localidade": "São Paulo",
			"uf": "SP",
			"ibge": "3550308",
			"gia": "1004",
			"ddd": "11",
			"siafi": "7107"
		  }
		`
		rw.Write([]byte(responseBody))
		rw.WriteHeader(http.StatusOK)
	}))
	// Close the server when test finishes
	defer server.Close()
	viacep.Client = server.Client()
	viacep.Host = server.URL

	app := fiber.New()
	SetUpRoutes(app)
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
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}))
	// Close the server when test finishes
	defer server.Close()
	viacep.Client = server.Client()
	viacep.Host = server.URL

	app := fiber.New()
	SetUpRoutes(app)
	req := httptest.NewRequest("GET", fmt.Sprintf("/via_cep_wrapper/%s", zipCode), nil)
	resp, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

