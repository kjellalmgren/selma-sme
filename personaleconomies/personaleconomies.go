package personaleconomies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"selmasme/models"
)

// GetPersonalEconomies
func GetPersonalEconomies(w http.ResponseWriter, r *http.Request) {

	//
	processid := r.Header.Get("X-process-Id")
	fmt.Printf("getPersonalEconomies executed: processId: %s...\r\n", processid)
	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	personaleconomies := []models.PersonalEconomyType{}
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		file, err := ioutil.ReadFile("json/personaleconomies.json")
		if err != nil {
			fmt.Fprintf(w, "Error reading personaleconomies.json - %s", err)
			w.WriteHeader(http.StatusNotFound)
		}
		_ = json.Unmarshal([]byte(file), &personaleconomies)
	}
	//
	if err := json.NewEncoder(w).Encode(personaleconomies); err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//
}

//
// GetPersonalEconomy
func GetPersonalEconomy(w http.ResponseWriter, r *http.Request) {

	//
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://app.swaggerhub.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me, X-process-ID")
	//
	processid := r.Header.Get("X-process-Id")
	//
	var data models.PersonalEconomyID
	var r1 []byte
	r1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	//
	json.NewDecoder(bytes.NewReader([]byte(r1))).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	//
	fmt.Printf("getPersonalEconomies executed: processId: %s /personalEconomyId: %s...\r\n",
		processid, data.PersonalEconomyID)
	//
	personaleconomies := []models.PersonalEconomyType{}
	file, err := ioutil.ReadFile("json/personaleconomies.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading personaleconomies.json - %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	peconret := make([]models.PersonalEconomyType, 1, 1)
	_ = json.Unmarshal([]byte(file), &personaleconomies)
	//
	switch processid {
	case "9a65d28a-46bb-4442-b96d-6a09fda6b18b":
		for i := 0; i < len(personaleconomies); i++ {
			if personaleconomies[i].PersonalEconomyID == data.PersonalEconomyID {
				peconret[0] = personaleconomies[i]
			}
		}
	}
	if peconret[0].PersonalEconomyID == data.PersonalEconomyID {
		if err := json.NewEncoder(w).Encode(peconret); err != nil {
			w.WriteHeader(http.StatusNotFound)
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "PersonalEconomy not Found", 404)
		//w.WriteHeader(http.StatusNotFound)
	}
}
