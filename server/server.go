package server

import (
	"bytes"
	"log"
	"net/http"

	"realizr.io/iso20022/repo"

	"encoding/json"
)

func GetJSONRepresentation(element *repo.Element) (string, error) {
	if element == nil {
		return "", nil
	}

	jsonBytes, err := json.Marshal(element)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func outputElement(element *repo.Element, buffer *bytes.Buffer) {
	if element != nil {
		buffer.WriteString("<div>" + *element.Name + " - " + *element.Description)
		for _, child := range element.Children {
			outputElement(child, buffer)
		}
		buffer.WriteString("</div>")
	}
}

func AddDefaultHeaders(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}
	w.Header().Add("Access-Control-Allow-Origin", origin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "GET,OPTIONS,PUT,POST,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization, ")
	w.Header().Add("Permissions-Policy", "interest-cohort=()")
	w.Header().Add("Cache-Control", "no-store, max-age=0")
}

func Options(w http.ResponseWriter, r *http.Request) {
	AddDefaultHeaders(w, r)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Default().Println(r.Method, " ", r.URL.Path, " ", 200)
}
