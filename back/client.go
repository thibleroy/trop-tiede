package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"back/src/lib"
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
	track := lib.ITrack{
		Title:     "myCreatedTitle",
		Artist:    "myArtist",
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

func putRequest(url string, track lib.ITrack){
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
	str, _ := json.Marshal(lib.ITrack{})
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
	url := "http://localhost:8080"
	uri := postRequest(url + "/tracks")
	fmt.Println("post ok")
	var track lib.ITrack
	val := getRequest(uri)
	err := json.Unmarshal(val, &track)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created track", track)
	track.Artist = "modified"
	fmt.Println("put request : " + uri, track)
	putRequest(uri, track)
	var upTrack lib.ITrack
	updatedTrack := getRequest(uri)
	err2 := json.Unmarshal(updatedTrack, &upTrack)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("updated", upTrack)
	deleteRequest(uri)
	fmt.Println("deleted")
	getRequest(uri)
}
