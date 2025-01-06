package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/WENDELLDELIMA/go-expert-multithreading/internal/models"
)

func FetchAddress(cep, url, apiSource string, resultChan chan models.Address) {
	resp, err := http.Get(fmt.Sprintf(url, cep))
	if err != nil {
		fmt.Printf("Erro na %s: %v\n", apiSource, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Erro: %s retornou status: %d\n", apiSource, resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao ler o corpo da resposta da %s: %v\n", apiSource, err)
		return
	}

	var address models.Address
	if err := json.Unmarshal(body, &address); err != nil {
		fmt.Printf("Erro ao fazer Unmarshal da %s: %v\n", apiSource, err)
		return
	}

	address.ApiSource = apiSource
	resultChan <- address
}
