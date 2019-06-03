package bankid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// EndUserInfoType documentation
type EndUserInfoType struct {
	EndUserInfo             []EndUserInfo `json:"endUserInfo"`
	RequirementAlternatives interface{}   `json:"requirementAlternatives"`
	DisplayName             string        `json:"displayName"`
	TransactionID           string        `json:"transactionId"`
	PersonalNumber          string        `json:"personalNumber"`
	Device                  string        `json:"device"`
}

// EndUserInfo struct
type EndUserInfo struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// AuthenticateResponse documentation
type AuthenticateResponse struct {
	TransactionID  string      `json:"transactionId"`
	OrderRef       string      `json:"orderRef"`
	AutoStartToken string      `json:"autoStartToken"`
	ErrorMessage   interface{} `json:"errorMessage"`
}

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// Configuration documentation
type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}

// NewConfiguration documentation
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://securetest2.landshypotek.se",
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
	}
	return cfg
}

// KeyValue documentation
type KeyValue struct {
	Key   string
	Value string
}

// AuthenticatePerson documentation
func AuthenticatePerson(_ssn string) int {

	var auth BasicAuth
	auth.UserName = "testQ"
	auth.Password = "testQtest"
	cfg := NewConfiguration()
	cfg.Host = "https://securetest2.landshypotek.se/api/identity/authenticate"
	h := make(map[string]string)
	fmt.Println("1")
	h["content-type"] = string("application/x-www-form-urlencoded")
	fmt.Println("2")
	//h["content-type"] = string("application/json")
	cfg.DefaultHeader = h
	var enduser EndUserInfoType
	enduserinfo := make([]EndUserInfo, 1, 1)
	enduserinfo[0].Type = "IP_ADDR"
	enduserinfo[0].Value = "Null"
	enduser.EndUserInfo = enduserinfo
	enduser.Device = "SAME"
	enduser.DisplayName = "Test mobilt BankID"
	enduser.PersonalNumber = _ssn
	enduser.RequirementAlternatives = nil
	enduser.TransactionID = "d6732e0c-a1ec-4db2-b58a-53ad53beb2b1"
	out, err := json.Marshal(enduser)
	if err != nil {
		panic(err)
	}
	bodystring := []byte(out)
	request, err := http.NewRequest("POST", cfg.Host, bytes.NewBuffer(bodystring))
	request.Header.Set("Access-Control-Allow-Credentials", "true")
	request.SetBasicAuth(auth.UserName, auth.Password)
	request.Header.Set("content-type", "application/x-www-form-urlencoded")
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
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

	return 0
}
