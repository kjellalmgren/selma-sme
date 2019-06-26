package roaring

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
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

	// QA
	key.consumerKey = "hApFNht4DgDjY4QVLtiZNnpgJWEa"
	key.secretKey = "dfab3Krxb80p71QaqITP9P8bmKga"
	// PROD
	//key.consumerKey = ""
	//key.secretKey = ""

	data := fmt.Sprintf("%s:%s", key.consumerKey, key.secretKey)
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	//fmt.Println(sEnc)
	return sEnc
}

// GetRoaringAccessToken Documentation
func GetRoaringAccessToken() models.Header {
	//
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	//
	bodystring := []byte("grant_type=client_credentials")
	request, err := http.NewRequest("POST", "https://api.roaring.io/token", bytes.NewBuffer(bodystring))
	request.Header.Set("Access-Control-Allow-Credentials", "true")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//
	//request.Header.Set("Accept-Encoding", "gzip")
	request.Header.Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", getEncodedString()))
	request.Header.Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID, access_token, access_type, expires_in, scope")
	//
	//dump := func(request *http.Request) {
	output, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
		//return
	}
	fmt.Println(fmt.Sprintf("%s", string(output)))
	//
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	//
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	fmt.Println(fmt.Sprintf("Status: %v", resp.Status))
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fmt.Sprintf("%s", string(body)))
	//
	var header models.Header
	header.AccessToken = resp.Header.Get("access_token")
	header.TokenType = resp.Header.Get("access_type")
	header.Scope = resp.Header.Get("scope")
	header.ExpiresIn = fmt.Sprintf("%v", resp.Header.Get("expires_in"))
	return header
}
