package server

import (
	"log"
	"net/http"

	"bytes"

	"realizr.io/iso20022/model"
	"realizr.io/iso20022/repo"

	"encoding/json"
)

func RunWebServer(model *model.Iso20022) {
	port := ":8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		element := repo.ExpandElement(id, model, nil)
		if element != nil {
			json, err := GetJSONRepresentation(element)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(json))
			}
		} else {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
	})

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

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
