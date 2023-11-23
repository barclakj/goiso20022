package iso20022

import (
	"encoding/json"
	"net/http"
	"regexp"
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
