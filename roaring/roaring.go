package roaring

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
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

	key.consumerKey = "HVD7kM1uBFFzSsVKDmpy7YjRfeIa"
	key.secretKey = "buXOYxxSujXJW0lAOaUaxqtpqyUa"

	data := fmt.Sprintf("%s:%s", key.consumerKey, key.secretKey)
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	//fmt.Println(sEnc)
	return sEnc
}

// GetRoaringAccessToken Documentation
func GetRoaringAccessToken() models.Header {
	//requestBody, err := json.Marshal(map[string]string{
	//	"grant_type": "client_credentials",
	//})
	//
	//file, err := ioutil.ReadFile("json/get_access_token.json")
	//if err != nil {
	//	log.Fatalln("Error reading get_access_token.json")
	//fmt.Fprintf(w, "Error reading get_access_token.json - %s", err)
	//w.WriteHeader(http.StatusNotFound)
	//}
	//grant_type := []models.Get_Access_Token{}
	var grant_type models.Get_Access_Token
	//appret[0] := []models.Applicant{}
	//grantret := make([]models.Get_Access_Token, 1, 1)
	grant_type.Credentials = "client_credentials"
	h := json.RawMessage(`{"grant_type": "client_credentials"}`)
	//_ = json.Unmarshal([]byte(file), &grant_type)
	//
	//str := "grant_type: client_credentials"
	//b, err := json.MarshalIndent(&h, "", "\t")
	requestBody := h
	//requestBody, err := json.Marshal(b)
	//
	//fmt.Println(fmt.Sprintf("grant_type: %s", grant_type.Grant_Type))
	//fmt.Println(fmt.Sprintf("grant_type: %v", grant_type))
	//if err != nil {
	//	log.Fatalln(err)
	//}
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	//
	request, err := http.NewRequest("POST", "https://api.roaring.io/token", bytes.NewBuffer(requestBody))
	//request.Header.Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
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
