package repo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"realizr.io/iso20022/model"
)

type Element struct {
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Children    []*Element `json:"children"`
	Type        *string    `json:"type"`
}

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

func ExpandElement(identifier string, model *model.Iso20022, parent *Element) *Element {
	var element *Element

	// Loop through toplevelcatalogueentries
	for _, entry := range model.DataDictionary.ListOfTopLevelDictionaryEntry {
		// Process each entry
		// ...
		if *entry.XmiId == identifier || *entry.Name == identifier {
			element = &Element{entry.Name, entry.Definition, []*Element{}, nil}
			if parent != nil {
				parent.Children = append(parent.Children, element)
			}
			// Loop through elements
			for _, child := range entry.ListOfMessageElement {
				// Process each element
				if child.SimpleType != nil {
					c := ExpandElement(*child.SimpleType, model, nil)

					childElement := &Element{child.Name, child.Definition, []*Element{}, c.Name}
					element.Children = append(element.Children, childElement)
				}
				if child.ComplexType != nil {
					complexElement := ExpandElement(*child.ComplexType, model, nil)
					childElement := &Element{child.Name, child.Definition, []*Element{}, complexElement.Name}
					childElement.Children = append(childElement.Children, childElement)
					childElement.Children = append(childElement.Children, complexElement)
				}
				if child.Type != nil {
					ExpandElement(*child.Type, model, element)
				}
			}
		}
		if element != nil {
			break
		}
	}
	if element == nil {
		fmt.Printf("Element %v not found\n", identifier)
	}
	return element
}
