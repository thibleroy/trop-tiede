package main

import (
	"back/lib"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getRequest(url string) []byte{
	var response []byte
	var resp, err = http.Get(url)
	if resp != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode == 200 {
			response = body
			fmt.Println("body value", string(body))
		} else {
			fmt.Println("http error", resp.Status)
		}
	} else {
		fmt.Println("Connexion error", err)
	}
	return response
}

func postRequest(url string) string{
	track := lib.IRoom{

	}
	str, _ := json.Marshal(track)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(str))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	fmt.Println("header loc", resp.Header["Location"])
	return resp.Header["Location"][0]
}

func putRequest(url string, track lib.IRoom){
	fmt.Println("url:", url)
	str, _ := json.Marshal(track)
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(str))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status put:", resp.Status)
	fmt.Println("response Headers put:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body put:", string(body))
	createdId := resp.Header
	fmt.Println("header loc", createdId)
}

func deleteRequest(url string) {
	fmt.Println("url:", url)
	str, _ := json.Marshal(lib.IRoom{})
	req, _ := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(str))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status del:", resp.Status)
	fmt.Println("response Headers del:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body del:", string(body))
	createdId := resp.Header
	fmt.Println("header loc", createdId)
}

func main(){
	url := "http://localhost:8080/rooms/123"
	fmt.Println("post ok")
	var room lib.IRoom
	val := getRequest(url)
	err := json.Unmarshal(val, &room)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created track", room)

}
