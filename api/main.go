package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Name     string `json:"name"`
	LastName string `json:"hugo"`
	Desc     string `json:"desc"`
}

type Articles []Article

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

func allArticle(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Name: "Angel", LastName: "Martin", Desc: "He is tall"},
	}
	json.NewEncoder(w).Encode(articles)
}

func weather(w http.ResponseWriter, r *http.Request) {

	city := r.FormValue("city")
	var URL = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=8d70211ac55cc3de75d16e818e6336b7&units=metric", city)
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n %v \n\n", response.Body)
	bytes, errRead := ioutil.ReadAll(response.Body)
	defer func() { // al final
		err := response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	if errRead != nil {
		log.Fatal(errRead)
	}
	log.Print(bytes)
	log.Print(string(bytes))

	var wr WeatherResponse
	errUnmarshal := json.Unmarshal(bytes, &wr)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	log.Printf("%+v", wr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wr)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "What sap")
}

func handlerRequests() {
	http.HandleFunc("/", homePage) // rutas
	http.HandleFunc("/martin", allArticle)
	http.HandleFunc("/weather", weather)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handlerRequests()
}
