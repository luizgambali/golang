package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type DadosIp struct {
	City              string  `json:"city"`
	Region            string  `json:"region"`
	Country           string  `json:"country"`
	CountryCode       string  `json:"countryCode"`
	Postal            string  `json:"zip"`
	Latitude          float64 `json:"lat"`
	Longitude         float64 `json:"long"`
	Timezone          string  `json:"timeZone"`
	CountryArea       float64 `json:"countryArea"`
	CountryPopulation int     `json:"countryPopulation"`
	Org               string  `json:"org"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	localIp := "8.8.8.8" // IP externo do Google

	// a API é paga, então funciona apenas com os endereços de IP abaixo
	// porém, deixei a função ObterIpExterno, para demonstrar como ficaria para
	// obter o seu IP externo e passar o valor

	url := "http://ip-api.com/json/" + localIp

	ObterDadosIp(url)
}

func ObterDadosIp(url string) {
	var dados DadosIp

	response, err := myClient.Get(url)
	response.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Não foi possível obter os dados do IP informado")
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&dados)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("City: ", dados.City)
		fmt.Println("Region: ", dados.Region)
		fmt.Println("Country: ", dados.Country)
		fmt.Println("CountryCode: ", dados.CountryCode)
		fmt.Println("Postal: ", dados.Postal)
		fmt.Println("Latitude: ", dados.Latitude)
		fmt.Println("Longitude: ", dados.Longitude)
		fmt.Println("Timezone: ", dados.Timezone)
		fmt.Println("CountryArea: ", dados.CountryArea)
		fmt.Println("CountryPopulation: ", dados.CountryPopulation)
		fmt.Println("Org: ", dados.Org)
	}

}

func ObterIpExterno() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
