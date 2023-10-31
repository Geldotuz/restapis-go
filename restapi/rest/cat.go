package rest

import (
	"fmt"
	"net/http"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

const url = "https://catFact.ninja/fact"

func (a *APP) GetCatFact(w http.ResponseWriter, r *http.Request) {
	var catFact CatFact

	err := GetJson(url, &catFact)
	if err != nil {
		fmt.Printf("error getting cat fact: %s\n", err.Error())
		return
	}
	fmt.Printf("a super interesting Cat Fact: %s\n", catFact.Fact)

	/*response, err := http.Get(url)
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

	errUnmarshal := json.Unmarshal(bytes, &catFact)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	log.Printf("%+v", catFact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(catFact)*/

}
