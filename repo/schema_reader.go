package repo

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"realizr.io/iso20022/model"
)

type Element struct {
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Children    []*Element `json:"children,omitempty"`
	Type        *string    `json:"type,omitempty"`
	Required    bool       `json:"required"`
	Attribute   bool       `json:"attribute"`
	MaxLength   *int       `json:"maxLength,omitempty"`
	MinLength   *int       `json:"minLength,omitempty"`
	MinValue    *float64   `json:"minValue,omitempty"`
	Pattern     *string    `json:"pattern,omitempty"`
	MaxOccurs   *int       `json:"maxOccurs,omitempty"`
}

type CatalogueEntry struct {
	Name              *string `json:"name,omitempty"`
	Description       *string `json:"description,omitempty"`
	MessageName       *string `json:"messageName,omitempty"`
	MessageDefinition *string `json:"messageDefinition,omitempty"`
	Domain            *string `json:"domain,omitempty"`
	FunctionalArea    *string `json:"functionalArea,omitempty"`
}

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

func ExpandElement(identifier string, model *model.Iso20022, parent *Element) *Element {
	var element *Element

	// Loop through toplevelcatalogueentries
	for _, entry := range model.DataDictionary.ListOfTopLevelDictionaryEntry {
		// Process each entry
		// ...
		if *entry.XmiId == identifier || *entry.Name == identifier {
			element = &Element{entry.Name, substNewLines(entry.Definition), []*Element{}, nil, false, false, nil, nil, nil, nil, nil}
			if parent != nil {
				parent.Children = append(parent.Children, element)
			}

			// Loop through elements
			for _, child := range entry.ListOfMessageElement {
				var childElement *Element
				// Process each element
				if child.SimpleType != nil {
					simpleElement := ExpandElement(*child.SimpleType, model, nil)
					childElement = &Element{child.Name, substNewLines(child.Definition), []*Element{}, simpleElement.Type, false, false, simpleElement.MaxLength, simpleElement.MinLength, simpleElement.MinValue, simpleElement.Pattern, nil}
					element.Children = append(element.Children, childElement)
				} else if child.ComplexType != nil {
					complexElement := ExpandElement(*child.ComplexType, model, nil)
					childElement = &Element{child.Name, substNewLines(child.Definition), []*Element{}, complexElement.Name, false, false, nil, nil, nil, nil, nil}
					childElement.Children = append(childElement.Children, complexElement.Children...)
					element.Children = append(element.Children, childElement)
				} else if child.Type != nil {
					otherElement := ExpandElement(*child.Type, model, nil)
					childElement = &Element{child.Name, substNewLines(child.Definition), []*Element{}, otherElement.Name, false, false, nil, nil, nil, nil, nil}
					childElement.Children = append(childElement.Children, otherElement.Children...)
					element.Children = append(element.Children, childElement)
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
			if entry.MinInclusive != nil {
				element.MinValue = entry.MinInclusive
			}
			if entry.TotalDigits != nil && (entry.FractionDigits == nil || *entry.FractionDigits == 0) {
				var tp = "integer"
				element.Type = &tp
				element.MaxLength = entry.TotalDigits
			} else if entry.TotalDigits != nil {
				var tp = "double"
				element.Type = &tp
				element.MaxLength = entry.TotalDigits
			} else if entry.MinLength != nil {
				var tp = "string"
				element.Type = &tp
				element.MaxLength = entry.MaxLength
				element.MinLength = entry.MinLength
			} else if entry.Pattern != nil {
				var tp = "string"
				element.Type = &tp
				element.Pattern = entry.Pattern
			} else {
				element.Type = entry.Name
			}
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

func ExpandCatalogue(identifier string, model *model.Iso20022) *Element {
	var element *Element

	// Loop through toplevelcatalogueentries
	for _, entry := range model.BusinessProcessCatalogue.ListOfTopLevelCatalogueEntries {
		for _, messageDefinitionChild := range entry.ListOfMessageDefinition {
			if *messageDefinitionChild.Name == identifier {
				element = &Element{entry.Name, substNewLines(entry.Definition), []*Element{}, nil, false, false, nil, nil, nil, nil, nil}

				for _, buildingBlock := range messageDefinitionChild.ListOfMessageBuildingBlock {
					if buildingBlock.ComplexType != nil {
						complexElement := ExpandElement(*buildingBlock.ComplexType, model, nil)
						subElement := &Element{buildingBlock.Name, substNewLines(buildingBlock.Definition), []*Element{}, complexElement.Name, false, false, nil, nil, nil, nil, nil}
						subElement.Children = append(subElement.Children, complexElement.Children...)
						if buildingBlock.MinOccurs != nil && *buildingBlock.MinOccurs > 0 {
							element.Required = true
						}
						if buildingBlock.MaxOccurs != nil {
							element.MaxOccurs = buildingBlock.MaxOccurs
						}
						element.Children = append(element.Children, subElement)
					}
				}
			}
			if element != nil {
				break
			}
		}
	}
	return element
}

func ListCatalogue(model *model.Iso20022) *[]CatalogueEntry {
	var elements []CatalogueEntry
	for _, catEntry := range model.BusinessProcessCatalogue.ListOfTopLevelCatalogueEntries {
		for _, messageDefinitionChild := range catEntry.ListOfMessageDefinition {
			if messageDefinitionChild.ListOfMessageDefinitionIdentifier != nil {
				entry := CatalogueEntry{catEntry.Name, catEntry.Definition, messageDefinitionChild.Name, messageDefinitionChild.Definition, nil, nil}
				for _, messageDefinitionIdentifier := range messageDefinitionChild.ListOfMessageDefinitionIdentifier {
					entry.Domain = messageDefinitionIdentifier.BusinessArea
					entry.FunctionalArea = messageDefinitionIdentifier.MessageFunctionality
				}
				elements = append(elements, entry)
			}
		}
	}
	return &elements
}

func FilterCatalogueByDomain(catalogue *[]CatalogueEntry, domain string, latest bool) *[]CatalogueEntry {
	var elements []CatalogueEntry
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

func FilterMandatoryElements(entry *Element) *Element {
	var element = &Element{entry.Name, substNewLines(entry.Description), []*Element{}, entry.Type, entry.Required, entry.Attribute, entry.MaxLength, entry.MinLength, entry.MinValue, entry.Pattern, nil}
	for _, child := range entry.Children {
		if child.Required {
			mandatoryChild := FilterMandatoryElements(child)
			element.Children = append(element.Children, mandatoryChild)
		}
	}
	return element
}

func substNewLines(str *string) *string {
	if str == nil {
		return nil
	}
	str2 := strings.ReplaceAll(*str, "\r", "")
	str2 = strings.ReplaceAll(str2, "\n", "<p/>")
	return &str2
}
