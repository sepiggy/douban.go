package main

import (
	"net/http"
	"log"
	"fmt"
	"bytes"
	"encoding/json"
)

type UserInfo struct {
	Id       string `json:"id"`
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Alt      string `json:"alt"`
	Relation string `json:"relation"`
	Created  string `json:"created"`
	LocID    string `json:"loc_id"`
	Desc     string `json:"desc"`
}

func main() {
	client := &http.Client{}
	url := "https://api.douban.com/v2/user/90513287"

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	response, _ := client.Do(request)

	status := response.StatusCode
	if status != 200 {
		log.Fatalln("status is not 200")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	jsonStr := buf.String()

	userInfo := UserInfo{}
	err = json.Unmarshal([]byte(jsonStr), &userInfo)
	fmt.Printf("%+v\n", userInfo)

}
