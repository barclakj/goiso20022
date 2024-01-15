package repo

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"realizr.io/iso20022/model"
)

type Element struct {
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Children    []*Element `json:"children"`
	Type        *string    `json:"type"`
	Required    bool       `json:"required"`
	Attribute   bool       `json:"attribute"`
}

func loadURL(url string) ([]byte, error) {
	log.Printf("Loading %v\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("Read %v bytes from %v\n", len(data), url)

	return data, nil
}

func ReadXMLFile(filePath string) (*model.Iso20022, error) {
	data, err := loadURL(filePath)
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
			element = &Element{entry.Name, entry.Definition, []*Element{}, nil, false, false}
			if parent != nil {
				parent.Children = append(parent.Children, element)
			}
			// Loop through elements
			for _, child := range entry.ListOfMessageElement {
				var childElement *Element
				// Process each element
				if child.SimpleType != nil {
					c := ExpandElement(*child.SimpleType, model, nil)
					childElement = &Element{child.Name, child.Definition, []*Element{}, c.Name, false, false}
					element.Children = append(element.Children, childElement)
				} else if child.ComplexType != nil {
					complexElement := ExpandElement(*child.ComplexType, model, nil)
					childElement = &Element{child.Name, child.Definition, []*Element{}, complexElement.Name, false, false}
					childElement.Children = append(childElement.Children, childElement)
					childElement.Children = append(childElement.Children, complexElement)
				} else if child.Type != nil {
					childElement = ExpandElement(*child.Type, model, element)
				}

				if childElement != nil {
					if child.MinOccurs != nil && *child.MinOccurs > 0 {
						childElement.Required = true
					}
					if child.Type != nil && *child.Type == "iso20022:MessageAttribute" {
						childElement.Attribute = true
					}
				}
			}
		}
		if element != nil {
			break
		}
	}
	if element == nil {
		log.Printf("Element %v not found in %v\n", identifier, *parent.Name)
	}
	return element
}

func FilterMandatoryElements(entry *Element) *Element {
	var element = &Element{entry.Name, entry.Description, []*Element{}, entry.Type, entry.Required, entry.Attribute}
	for _, child := range entry.Children {
		if child.Required {
			mandatoryChild := FilterMandatoryElements(child)
			element.Children = append(element.Children, mandatoryChild)
		}
	}
	return element
}
