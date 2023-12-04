package repo

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	"realizr.io/iso20022/model"
)

func ReadXMLFile(filePath string) (*model.Iso20022, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var _model model.Iso20022
	err = xml.Unmarshal(data, &_model)
	if err != nil {
		return nil, err
	}

	return &_model, nil
}
