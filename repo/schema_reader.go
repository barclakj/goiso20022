package repo

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"

	"realizr.io/iso20022/model"
)

func ReadXMLFile(filePath string) (*model.Iso20022, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var _model model.Iso20022
	var currentTopLevelDictionaryEntry *model.TopLevelDictionaryEntry

	var unknownElements []string

	decoder := xml.NewDecoder(file)
	for {
		tkn, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch token := tkn.(type) {
		case xml.StartElement:
			switch token.Name.Local {
			case "dataDictionary":
				for _, attr := range token.Attr {
					if attr.Name.Local == "xmi:id" {
						_model.DataDictionary.XmiId = attr.Value
					}
				}
			// case "businessProcessCatalogue":

			case "topLevelDictionaryEntry":
				currentTopLevelDictionaryEntry = toTopLevelDictionaryEntry(&token)
				_model.DataDictionary.ListOfTopLevelDictionaryEntry = append(_model.DataDictionary.ListOfTopLevelDictionaryEntry, *currentTopLevelDictionaryEntry)
			case "element":
				element, err := toELement(&token)
				if err != nil {
					return nil, err
				}
				if currentTopLevelDictionaryEntry != nil {
					currentTopLevelDictionaryEntry.ListOfElement = append(currentTopLevelDictionaryEntry.ListOfElement, *element)
				}
				// model.PrintElementAttributes(element)
				// fmt.Println(element.Name + " " + element.Definition)
			default:
				cont := true
				for _, ue := range unknownElements {
					if ue == token.Name.Local {
						cont = false
						break
					}
				}
				if cont {
					unknownElements = append(unknownElements, token.Name.Local)
				}
			}
		case xml.EndElement:
			switch token.Name.Local {
			case "topLevelDictionaryEntry":
				currentTopLevelDictionaryEntry = nil
			}
		case xml.CharData:
			// fmt.Println("data:", token)
		}

	}

	for _, element := range unknownElements {
		fmt.Println("Missing parser for: " + element)
	}

	return &_model, nil
}

func toELement(token *xml.StartElement) (*model.Element, error) {
	var element model.Element
	for _, attr := range token.Attr {
		if attr.Name.Local == "xsi:type" {
			element.XsiType = attr.Value
		}
		if attr.Name.Local == "xmi:id" {
			element.XmiId = attr.Value
		}
		if attr.Name.Local == "name" {
			element.Name = attr.Value
		}
		if attr.Name.Local == "definition" {
			element.Definition = attr.Value
		}
		if attr.Name.Local == "isDerived" {
			boolVal, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return nil, err
			}
			element.IsDerived = boolVal
		}

		if attr.Name.Local == "registrationStatus" {
			element.RegistrationStatus = attr.Value
		}
		if attr.Name.Local == "minOccurs" {
			minOccurs, err := strconv.Atoi(attr.Value)
			if err != nil {
				return nil, err
			}
			element.MinOccurs = minOccurs
		}
		if attr.Name.Local == "maxOccurs" {
			maxOccurs, err := strconv.Atoi(attr.Value)
			if err != nil {
				return nil, err
			}
			element.MaxOccurs = maxOccurs
		}
		if attr.Name.Local == "derivation" {
			element.Derivation = attr.Value
		}
		if attr.Name.Local == "opposite" {
			element.Opposite = attr.Value
		}
		if attr.Name.Local == "type" {
			element.Type = attr.Value
		}
		if attr.Name.Local == "simpleType" {
			element.SimpleType = attr.Value
		}
	}
	return &element, nil
}

func toTopLevelDictionaryEntry(token *xml.StartElement) *model.TopLevelDictionaryEntry {
	var entry model.TopLevelDictionaryEntry
	for _, attr := range token.Attr {
		if attr.Name.Local == "xsi:type" {
			entry.XsiType = attr.Value
		}
		if attr.Name.Local == "xmi:id" {
			entry.XmiId = attr.Value
		}
		if attr.Name.Local == "name" {
			entry.Name = attr.Value
		}
		if attr.Name.Local == "definition" {
			entry.Definition = attr.Value
		}
		if attr.Name.Local == "registrationStatus" {
			entry.RegistrationStatus = attr.Value
		}
		if attr.Name.Local == "subType" {
			entry.SubType = attr.Value
		}
		if attr.Name.Local == "derivationComponent" {
			entry.DerivationComponent = attr.Value
		}
		if attr.Name.Local == "associationDomain" {
			entry.AssociationDomain = attr.Value
		}
		if attr.Name.Local == "derivationElement" {
			entry.DerivationElement = attr.Value
		}
	}

	return &entry
}
