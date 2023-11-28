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
	var currentTopLevelCatalogueEntry *model.TopLevelCatalogueEntry
	var currentBusinessRole *model.BusinessRole
	var currentSemanticMarkup *model.SemanticMarkup
	var currentMessageElement *model.MessageElement
	var currentElement *model.Element
	var currentMessageBuildingBlock *model.MessageBuildingBlock
	var currentMessageDefinition *model.MessageDefinition

	var stack Stack
	var lastName *string

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
			// fmt.Println("Next:", token.Name.Local)
			lastName = stack.Push(token.Name.Local)
			switch token.Name.Local {
			case "Repository":
			case "businessProcessCatalogue":
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
				currentElement, err = toELement(&token)
				if err != nil {
					return nil, err
				}
				if currentTopLevelDictionaryEntry != nil {
					currentTopLevelDictionaryEntry.ListOfElement = append(currentTopLevelDictionaryEntry.ListOfElement, *currentElement)
				}
			case "topLevelCatalogueEntry":
				currentTopLevelCatalogueEntry = toTopLevelCatalogueEntry(&token)
				_model.BusinessProcessCatalogue.ListOfTopLevelCatalogueEntries = append(_model.BusinessProcessCatalogue.ListOfTopLevelCatalogueEntries, *currentTopLevelCatalogueEntry)
			case "businessRole":
				currentBusinessRole = toBusinessRole(&token)
				if currentTopLevelCatalogueEntry != nil {
					currentTopLevelCatalogueEntry.ListOfBusinessRoles = append(currentTopLevelCatalogueEntry.ListOfBusinessRoles, *currentBusinessRole)
				} else {
					fmt.Println("Nothing to attach businessRole to " + *lastName)
				}
			case "semanticMarkup":
				currentSemanticMarkup = toSemanticMarkup(&token)
				switch *lastName {
				case "topLevelDictionaryEntry":
					currentTopLevelDictionaryEntry.ListOfSemanticMarkup = append(currentTopLevelDictionaryEntry.ListOfSemanticMarkup, *currentSemanticMarkup)
				case "businessRole":
					currentBusinessRole.ListOfSemanticMarkup = append(currentBusinessRole.ListOfSemanticMarkup, *currentSemanticMarkup)
				case "messageElement":
					currentMessageElement.ListOfSemanticMarkup = append(currentMessageElement.ListOfSemanticMarkup, *currentSemanticMarkup)
				case "element":
					currentElement.ListOfSemanticMarkup = append(currentElement.ListOfSemanticMarkup, *currentSemanticMarkup)
				case "messageBuildingBlock":
					currentMessageBuildingBlock.ListOfSemanticMarkup = append(currentMessageBuildingBlock.ListOfSemanticMarkup, *currentSemanticMarkup)
				case "messageDefinition":
					currentMessageDefinition.ListOfSemanticMarkup = append(currentMessageDefinition.ListOfSemanticMarkup, *currentSemanticMarkup)
				default:
					fmt.Println("Nothing to attach semanticMarkup to " + *lastName)
				}

			case "elements":
				elements := toELements(&token)
				if currentSemanticMarkup != nil {
					currentSemanticMarkup.ListOfElements = append(currentSemanticMarkup.ListOfElements, *elements)
				} else {
					fmt.Println("Nothing to attach elements to " + *lastName)
				}
			case "messageElement":
				currentMessageElement = toMessageElement(&token)
				if currentTopLevelDictionaryEntry != nil {
					currentTopLevelDictionaryEntry.ListOfMessageElement = append(currentTopLevelDictionaryEntry.ListOfMessageElement, *currentMessageElement)
				} else {
					fmt.Println("Nothing to attach messageElement to " + *lastName)
				}
			case "messageBuildingBlock":
				currentMessageBuildingBlock = toMessageBuildingBlock(&token)
			case "messageDefinition":
				currentMessageDefinition = toMessageDefinition(&token)
				currentTopLevelCatalogueEntry.ListOfMessageDefinition = append(currentTopLevelCatalogueEntry.ListOfMessageDefinition, *currentMessageDefinition)
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
			// fmt.Println("Pushed:", lastName)
		case xml.EndElement:
			lastName, err = stack.Pop()
			// fmt.Println("Popped:", lastName)
			if err != nil {
				return nil, err
			}
			switch token.Name.Local {
			case "topLevelDictionaryEntry":
				currentTopLevelDictionaryEntry = nil
			case "topLevelCatalogueEntry":
				currentTopLevelCatalogueEntry = nil
			case "businessRole":
				currentBusinessRole = nil
			case "semanticMarkup":
				currentSemanticMarkup = nil
			case "messageElement":
				currentMessageElement = nil
			case "element":
				currentElement = nil
			case "messageBuildingBlock":
				currentMessageBuildingBlock = nil
			case "messageDefinition":
				currentMessageDefinition = nil
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

func toMessageDefinition(token *xml.StartElement) *model.MessageDefinition {
	var messageDefinition model.MessageDefinition
	for _, attr := range token.Attr {
		if attr.Name.Local == "xmi:id" {
			messageDefinition.XmiId = attr.Value
		}
		if attr.Name.Local == "name" {
			messageDefinition.Name = attr.Value
		}
		if attr.Name.Local == "definition" {
			messageDefinition.Definition = attr.Value
		}
		if attr.Name.Local == "registrationStatus" {
			messageDefinition.RegistrationStatus = attr.Value
		}
		if attr.Name.Local == "messageSet" {
			messageDefinition.MessageSet = attr.Value
		}
		if attr.Name.Local == "xmlTag" {
			messageDefinition.XmlTag = attr.Value
		}
		if attr.Name.Local == "rootElement" {
			messageDefinition.RootElement = attr.Value
		}
	}
	return &messageDefinition
}

func toMessageBuildingBlock(token *xml.StartElement) *model.MessageBuildingBlock {
	var messageBuildingBlock model.MessageBuildingBlock
	for _, attr := range token.Attr {
		if attr.Name.Local == "xmi:id" {
			messageBuildingBlock.XmiId = attr.Value
		}
		if attr.Name.Local == "name" {
			messageBuildingBlock.Name = attr.Value
		}
		if attr.Name.Local == "definition" {
			messageBuildingBlock.Definition = attr.Value
		}
		if attr.Name.Local == "registrationStatus" {
			messageBuildingBlock.RegistrationStatus = attr.Value
		}
		if attr.Name.Local == "minOccurs" {
			minOccurs, err := strconv.Atoi(attr.Value)
			if err != nil {
				return nil
			}
			messageBuildingBlock.MinOccurs = minOccurs
		}
		if attr.Name.Local == "maxOccurs" {
			maxOccurs, err := strconv.Atoi(attr.Value)
			if err != nil {
				return nil
			}
			messageBuildingBlock.MaxOccurs = maxOccurs
		}
		if attr.Name.Local == "complexType" {
			messageBuildingBlock.ComplexType = attr.Value
		}
		if attr.Name.Local == "xmlTag" {
			messageBuildingBlock.XmlTag = attr.Value
		}
	}
	return &messageBuildingBlock
}

func toMessageElement(token *xml.StartElement) *model.MessageElement {
	var messageElement model.MessageElement
	for _, attr := range token.Attr {
		if attr.Name.Local == "xsi:type" {
			messageElement.XsiType = attr.Value
		}
		if attr.Name.Local == "xmi:id" {
			messageElement.XmiId = attr.Value
		}
		if attr.Name.Local == "name" {
			messageElement.Name = attr.Value
		}
		if attr.Name.Local == "definition" {
			messageElement.Definition = attr.Value
		}
		if attr.Name.Local == "registrationStatus" {
			messageElement.RegistrationStatus = attr.Value
		}
		if attr.Name.Local == "minOccurs" {
			minOccurs, err := strconv.Atoi(attr.Value)
			if err != nil {
				return nil
			}
			messageElement.MinOccurs = minOccurs
		}
		if attr.Name.Local == "maxOccurs" {
			maxOccurs, err := strconv.Atoi(attr.Value)
			if err != nil {
				return nil
			}
			messageElement.MaxOccurs = maxOccurs
		}
		if attr.Name.Local == "isDerived" {
			boolVal, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return nil
			}
			messageElement.IsDerived = boolVal
		}
		if attr.Name.Local == "complexType" {
			messageElement.ComplexType = attr.Value
		}
		if attr.Name.Local == "businessElementTrace" {
			messageElement.BusinessElementTrace = attr.Value
		}
		if attr.Name.Local == "xmlTag" {
			messageElement.XmlTag = attr.Value
		}
	}
	return &messageElement
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

func toTopLevelCatalogueEntry(token *xml.StartElement) *model.TopLevelCatalogueEntry {
	var catalogueEntry model.TopLevelCatalogueEntry
	for _, attr := range token.Attr {
		if attr.Name.Local == "xsi:type" {
			catalogueEntry.XsiType = attr.Value
		}
		if attr.Name.Local == "xmi:id" {
			catalogueEntry.XmiId = attr.Value
		}
		if attr.Name.Local == "name" {
			catalogueEntry.Name = attr.Value
		}
		if attr.Name.Local == "definition" {
			catalogueEntry.Definition = attr.Value
		}
		if attr.Name.Local == "registrationStatus" {
			catalogueEntry.RegistrationStatus = attr.Value
		}
	}
	return &catalogueEntry
}

func toBusinessRole(token *xml.StartElement) *model.BusinessRole {
	var businessRole model.BusinessRole
	for _, attr := range token.Attr {
		if attr.Name.Local == "xmi:id" {
			businessRole.XmiId = attr.Value
		}
		if attr.Name.Local == "name" {
			businessRole.Name = attr.Value
		}
		if attr.Name.Local == "definition" {
			businessRole.Definition = attr.Value
		}
		if attr.Name.Local == "registrationStatus" {
			businessRole.RegistrationStatus = attr.Value
		}
	}
	return &businessRole
}

func toSemanticMarkup(token *xml.StartElement) *model.SemanticMarkup {
	var semanticMarkup model.SemanticMarkup
	for _, attr := range token.Attr {
		if attr.Name.Local == "xmi:id" {
			semanticMarkup.XmiId = attr.Value
		}
		if attr.Name.Local == "type" {
			semanticMarkup.Type = attr.Value
		}
	}
	return &semanticMarkup
}

func toELements(token *xml.StartElement) *model.Elements {
	var elements model.Elements
	for _, attr := range token.Attr {
		if attr.Name.Local == "xmi:id" {
			elements.XmiId = attr.Value
		}
		if attr.Name.Local == "name" {
			elements.Name = attr.Value
		}
		if attr.Name.Local == "value" {
			elements.Value = attr.Value
		}
	}
	return &elements
}
