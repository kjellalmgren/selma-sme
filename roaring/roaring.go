package roaring

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Key struct {
	consumer_key string
	secret_key   string
}

//ckey const {"HVD7kM1uBFFzSsVKDmpy7YjRfeIa"}
//skey string {"buXOYxxSujXJW0lAOaUaxqtpqyUa"}
//
func GetAccessToken() string {

	var key Key

	key.consumer_key = "HVD7kM1uBFFzSsVKDmpy7YjRfeIa"
	key.secret_key = "buXOYxxSujXJW0lAOaUaxqtpqyUa"

	data := fmt.Sprintf("%s:%s", key.consumer_key, key.secret_key)
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
}

func X() {
	requestBody, err := json.Marshal(map[string]string{
		"grant_type": "client_credentials",
	})
	if err != nil {
		log.Fatalln(err)
	}
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest("POST", "https://api.roaring.io/token", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", GetAccessToken()))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))

}
