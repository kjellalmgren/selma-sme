package roaring

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"selmasme/models"
	"time"
)

// Key struct documentation
type Key struct {
	consumerKey string
	secretKey   string
}

//ckey const {"HVD7kM1uBFFzSsVKDmpy7YjRfeIa"}
//skey string {"buXOYxxSujXJW0lAOaUaxqtpqyUa"}
//
func getEncodedString() string {

	var key Key

	key.consumerKey = "HVD7kM1uBFFzSsVKDmpy7YjRfeIa"
	key.secretKey = "buXOYxxSujXJW0lAOaUaxqtpqyUa"

	data := fmt.Sprintf("%s:%s", key.consumerKey, key.secretKey)
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	//fmt.Println(sEnc)
	return sEnc
}

// GetRoaringAccessToken Documentation
func GetRoaringAccessToken() models.Header {
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

	request.Header.Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	request.Header.Set("Access-Control-Allow-Credentials", "true")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", getEncodedString()))
	request.Header.Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, access_token, access_type, expires_in, scope")
	if err != nil {
		log.Fatalln(err)
	}
	//
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	var header models.Header
	header.AccessToken = resp.Header.Get("access_token")
	header.TokenType = resp.Header.Get("access_type")
	header.Scope = resp.Header.Get("scope")
	header.ExpiresIn = fmt.Sprintf("%v", resp.Header.Get("expires_in"))
	return header
}
