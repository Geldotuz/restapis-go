package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RandomUser struct {
	Results []UserResult
}

type UserResult struct {
	Name    UserName
	Email   string
	Picture UserPicture
}

type UserName struct {
	Title string
	First string
	Last  string
}

type UserPicture struct {
	Large     string
	Medium    string
	Thumbnail string
}

func (a *APP) GetRandomUser(w http.ResponseWriter, r *http.Request) {
	url := "https://randomUser.me/api/?inc=name,email,picture"
	var user RandomUser

	err := GetJson(url, &user)
	if err != nil {
		fmt.Printf("error getting json: %s\n", err.Error())
		return
	}
	/*
		fmt.Fprintf(w, "User: %s %s %s\nEmail: %s\nThumbnail: %s",
			user.Results[0].Name.Title,
			user.Results[0].Name.First,
			user.Results[0].Name.Last,
			user.Results[0].Email,
			user.Results[0].Picture.Thumbnail)
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	//data, err := json.Marshal(user)
	//w.Write(data)
	//fmt.Fprintf(w, "%+v", user)

}
