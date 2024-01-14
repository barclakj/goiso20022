package function

import (
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"realizr.io/iso20022/model"
	"realizr.io/iso20022/repo"
	"realizr.io/iso20022/server"
)

var isoModel *model.Iso20022

func loadModel() *model.Iso20022 {
	url := "https://storage.googleapis.com/media.nonfunctionalarchitect.com/20230719_ISO20022_2013_eRepository.iso20022"

	// Call the function being tested
	iso20022model, err := repo.ReadXMLFile(url)
	if err != nil {
		log.Fatalf("Error reading XML file: %v", err)
		panic(err)
	}
	return iso20022model
}

func init() {
	functions.HTTP("ISO20022", ISO20022)
	isoModel = loadModel()
}

func ISO20022(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		element := repo.ExpandElement(id, isoModel, nil)
		if element != nil {
			json, err := server.GetJSONRepresentation(element)
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
	}
}
