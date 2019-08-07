package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

type TokenAPI struct {
	Url     string
	Version string
	Token   string
}

func (api TokenAPI) Get(endpoint string) string {
	var bearer = "Bearer " + api.Token
	url := api.Url + api.Version + endpoint
	log.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Erro ao realizar request.\n[ERRO] -", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string([]byte(body))

}
