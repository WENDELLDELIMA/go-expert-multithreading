package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/WENDELLDELIMA/go-expert-multithreading/internal/models"
	"github.com/WENDELLDELIMA/go-expert-multithreading/internal/services"
	"github.com/labstack/echo/v4"
)

func GetAddressHandler(c echo.Context) error {
	time.Sleep(time.Duration(500+rand.Intn(500)) * time.Millisecond)
	cep := c.Param("cep")
	if cep == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "CEP não pode estar vazio"})
	}

	resultChan := make(chan models.Address)
	timeout := 1 * time.Second

	go services.FetchAddress(cep, "https://brasilapi.com.br/api/cep/v1/%s", "BrasilAPI", resultChan)
	go services.FetchAddress(cep, "http://viacep.com.br/ws/%s/json/", "ViaCEP", resultChan)

	select {
	case result := <-resultChan:
		return c.JSON(http.StatusOK, result)
	case <-time.After(timeout):
		return c.JSON(http.StatusRequestTimeout, map[string]string{"error": "Timeout: Nenhuma API respondeu dentro de 1 segundo"})
	}
}
