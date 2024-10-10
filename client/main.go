package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	myGetMethod("Dog")
	myPostMethod("Dog", "NotAris")
	myPostMethod("Cat", "Aris")
	myGetMethod("Dog")
	myDeleteMethod("Dog")
	myGetMethod("Dog")
}

func myPostMethod(key string, value string) {
	requestData := map[string]string{
		"key":   key,
		"value": value,
	}

	//to json
	myJson, err := json.Marshal(requestData)
	if err != nil {
		log.Fatal("json marshal error", err)
	}

	//POST request
	request1, err := http.NewRequest("POST", "http://localhost:13281/mutate", bytes.NewBuffer(myJson))
	if err != nil {
		log.Fatalf("Ooops, something went wrong %s!", err)
	}

	//setting a header
	request1.Header.Set("Content-Type", "application/json; charset=UTF-8")

	//sending request
	client := &http.Client{}
	response, err := client.Do(request1)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	//reading a response from the server
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s", err)
	}

	//printing response status and body
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)

	fmt.Println("response Body:", string(body))
}

func myGetMethod(key string) {

	url := fmt.Sprintf("http://localhost:13281/mutate?key=%s", key)

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request: %s", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s", err)
	}
	fmt.Println("Response Status:", response.Status)
	fmt.Println("response Body:", string(body))
}

func myDeleteMethod(key string) {

	url := fmt.Sprintf("http://localhost:13281/mutate?key=%s", key)

	//new DELETE request
	request2, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Ooops, something went wrong %s!", err)
	}

	//send the request
	client := &http.Client{}
	response, err := client.Do(request2)
	if err != nil {
		panic(err)
	}

	//reading a response from the server
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s", err)
	}

	//printing response status and body
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)

	fmt.Println("response Body:", string(body))

}

// resp, err := client.Get("http://localhost:13281/")
// if err != nil {
// 	log.Fatalf("Ooops, something went wrong %s!", err)
// }

// client.Post("http://localhost:13281/", "contentType", http.NoBody)
// client.Get("http://localhost:13281/")
