package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type climaPrevisao struct{
	Location struct{
		Name string `json:"name"`
		Region string `json:"region"`
		Country string `json:"country"`
	} `json:""location""`
	Current struct{
		Temp_c float64 `json:"temp_c"`
		Condition struct{
			Text string `json:"text"`
		}`json:"condition"`
	} `json:"current"`
	Forecast struct{
		Forecastday []struct{
			Hour []struct{
				TimeEpoch int64 `json:"time_epoch"`
				Temp_c float64 `json:"temp_c"`
				Condition struct{
					Text string `json:"text"`
				}`json:"condition"`
				Humidity int8 `json:"humidity"`
				ChanceOfRain float64 `json:"chance_of_rain"`
				}`json:"hour"`
		}`json:"forecastday"`
	}`json:"forecast"`
}

func main() {
	//carega o .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env: ",err)
	}

	//importa os dados .env
	var API_KEY = os.Getenv("API_KEY")
	var LAT = os.Getenv("LAT")
	var LONG = os.Getenv("LONG")
	var URL string

	if len(os.Args) >= 2{
		 URL = fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no",API_KEY,os.Args[1])
	}else{
		 URL  = fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s,%s&days=1&aqi=no&alerts=no",API_KEY,LAT,LONG)
	}

	res, err := http.Get(URL)
	if err != nil{
		log.Panic("Erro ao comunicar com a api: ",err)
	}
	defer res.Body.Close();

	if res.StatusCode != 200{
		log.Panic("API com problemas: ")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil{
		log.Panic(err)
	}

	var previsao climaPrevisao 
	err = json.Unmarshal(body,&previsao)
	if err != nil{
		log.Panic(err)
	}

	fmt.Printf("Local: %s, %s, %s\n",
	previsao.Location.Name,
	previsao.Location.Region,
	previsao.Location.Country)

	fmt.Printf("Hora | Temp | Condição | Humidade | Chance Chuva\n")

	fmt.Printf("Agora|  %1.fc  | %s\n",
	previsao.Current.Temp_c,
	previsao.Current.Condition.Text)
	for _,hora := range previsao.Forecast.Forecastday[0].Hour{
		hor := time.Unix(hora.TimeEpoch,0)
		fmt.Printf("%s| %.1fc | %s | %d | %1.f\n",
		hor.Format("15:04"),
		hora.Temp_c,
		hora.Condition.Text,
		hora.Humidity,
		hora.ChanceOfRain)
	}

}
