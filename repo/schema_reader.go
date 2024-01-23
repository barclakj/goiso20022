package repo

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"realizr.io/iso20022/model"
)

var complex = "complex"
var simple = "simple"

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

func ExpandElement(identifier string, mdl *model.Iso20022, parent *model.BasicElement) *model.BasicElement {
	var ids []*string
	return ExpandElementWithIds(identifier, mdl, parent, ids)
}

func ExpandElementWithIds(identifier string, mdl *model.Iso20022, parent *model.BasicElement, ids []*string) *model.BasicElement {
	// Check if we have already processed this element
	for _, id := range ids {
		if strings.Compare(*id, identifier) == 0 {
			log.Printf("Element %v already processed\n", identifier)
			return nil
		}
	}
	ids = append(ids, &identifier)
	var element *model.BasicElement
	// log.Default().Printf("Searching for element: %v\n", identifier)

	// Loop through topleveldictionaryentries
	for _, entry := range mdl.DataDictionary.ListOfTopLevelDictionaryEntry {
		if *entry.XmiId == identifier || *entry.Name == identifier {
			// log.Default().Printf("Found element: %v as %v\n", identifier, *entry.Name)
			element = entry.ToElement()
			if parent != nil {
				parent.AddChild(element)
			}

			for _, msgElementChild := range entry.ListOfMessageElement {
				childElement := msgElementChild.ToElement()
				if childElement.Id == nil {
					log.Printf("Child element %v has no id\n", *msgElementChild.Name)
				} else {
					childTopLevelElement := ExpandElementWithIds(*childElement.Id, mdl, element, ids)
					if childTopLevelElement != nil {
						childTopLevelElement.Name = childElement.Name
						childTopLevelElement.MaxOccurs = childElement.MaxOccurs
						childTopLevelElement.MinOccurs = childElement.MinOccurs
						childTopLevelElement.Assess()
					} else {
						element.AddChild(childElement)
					}
				}
			}

			/*	for _, elementChild := range entry.ListOfElement {
				childElement := elementChild.ToElement()
				if childElement.Id == nil {
					log.Printf("Child element %v has no id\n", *elementChild.Name)
				} else {
					childTopLevelElement := ExpandElementWithIds(*childElement.Id, mdl, element, ids)
					if childTopLevelElement != nil {
						childTopLevelElement.Name = childElement.Name
						childTopLevelElement.MaxOccurs = childElement.MaxOccurs
						childTopLevelElement.MinOccurs = childElement.MinOccurs
						childTopLevelElement.Assess()
					} else {
						element.AddChild(childElement)
					}
				}
			} */
		}
		if element != nil {
			break
		}
	}
	if element == nil {
		if parent != nil {
			log.Printf("Element %v not found in %v\n", identifier, *parent.Name)
		} else {
			log.Printf("Element %v not found (nil parent)\n", identifier)
		}
	}
	return element
}

func ListCatalogue(mdl *model.Iso20022) *[]model.CatalogueEntry {
	var elements []model.CatalogueEntry
	for _, catEntry := range mdl.BusinessProcessCatalogue.ListOfTopLevelCatalogueEntries {
		elements = append(elements, catEntry.ToCatalogueEntries()...)
	}
	return &elements
}

func FilterCatalogueByDomain(catalogue *[]model.CatalogueEntry, domain string, latest bool) *[]model.CatalogueEntry {
	var elements []model.CatalogueEntry
	for _, catEntry := range *catalogue {
		if catEntry.Domain != nil && *catEntry.Domain == domain {
			if latest {
				if catEntry.Name != nil && strings.Contains(*catEntry.Name, "Latest version") {
					elements = append(elements, catEntry)
				}
			} else {
				elements = append(elements, catEntry)
			}
		}
	}
	return &elements
}

func FilterMandatoryElements(entry *model.BasicElement) *model.BasicElement {
	var element = entry.DuplicateNoChildren()
	for _, child := range entry.Children {
		if child.Required {
			mandatoryChild := FilterMandatoryElements(child)
			element.AddChild(mandatoryChild)
		}
	}
	return element
}
