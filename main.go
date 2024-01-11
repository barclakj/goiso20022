package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"realizr.io/iso20022/server" // This is a local package

	"realizr.io/iso20022/model"
	"realizr.io/iso20022/repo"
)

type Payment struct {
	amount float32
}

type ReferredDocumentInfo struct {
	Number            string
	RelatedDate       string
	CodeOrProprietary string
}

func validateRelatedDate(relatedDate string) bool {
	// Regular expression pattern for "yyyy-mm-dd" format
	pattern := `^\d{4}-\d{2}-\d{2}$`

	// Match the related date against the pattern
	match, _ := regexp.MatchString(pattern, relatedDate)

	return match
}

func validateReferredDocumentInfo(referredDocInfo ReferredDocumentInfo) bool {
	if referredDocInfo.Number == "" ||
		referredDocInfo.RelatedDate == "" ||
		referredDocInfo.CodeOrProprietary == "" ||
		validateRelatedDate(referredDocInfo.RelatedDate) == false {
		return false
	}
	return true
}

func HandleReferredDocumentInfo(w http.ResponseWriter, r *http.Request) {
	var referredDocInfo ReferredDocumentInfo

	// Decode the JSON body into the ReferredDocumentInfo struct
	err := json.NewDecoder(r.Body).Decode(&referredDocInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the ReferredDocumentInfo struct
	// ...

	// Send a response back
	w.WriteHeader(http.StatusOK)
}

func main() {
	iso := loadModel()
	log.Printf("Loaded model with %d top level dictionary entries", len(iso.DataDictionary.ListOfTopLevelDictionaryEntry))
	// Start the web server
	server.RunWebServer(iso)
}

func loadModel() *model.Iso20022 {
	filePath := "schema/20230719_ISO20022_2013_eRepository.iso20022"

	// Call the function being tested
	iso20022model, err := repo.ReadXMLFile(filePath)
	if err != nil {
		log.Fatalf("Error reading XML file: %v", err)
		panic(err)
	}
	return iso20022model
}
