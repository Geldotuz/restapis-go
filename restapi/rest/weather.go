package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main Main   `json:"main"`
}

type Main struct {
	Temp          float64 `json:"temp"`
	TempFeelsLike float64 `json:"feels_like"`
	TempMin       float64 `json:"temp_Min"`
	TempMax       float64 `json:"temp_Max"`
	Pressure      float64 `json:"pressure"`
}

func (a *APP) Weather(w http.ResponseWriter, r *http.Request) {

	city := r.FormValue("city")
	var URL = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=8d70211ac55cc3de75d16e818e6336b7&units=metric", city)
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n %v \n\n", response.Body)
	bytes, errRead := ioutil.ReadAll(response.Body)
	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	if errRead != nil {
		log.Fatal(errRead)
	}

	var wr WeatherResponse
	errUnmarshal := json.Unmarshal(bytes, &wr)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wr)

}
