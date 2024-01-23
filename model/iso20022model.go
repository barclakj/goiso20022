package model

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Iso20022 struct {
	XMLName                  xml.Name                  `xml:"Repository"`
	DataDictionary           *DataDictionary           `xml:"dataDictionary"`
	BusinessProcessCatalogue *BusinessProcessCatalogue `xml:"businessProcessCatalogue"`
	Namespace                *string                   `xml:"namespace,attr"`
}

type DataDictionary struct {
	XmiId                         *string                    `xml:"id,attr"`
	ListOfTopLevelDictionaryEntry []*TopLevelDictionaryEntry `xml:"topLevelDictionaryEntry"`
}

type TopLevelDictionaryEntry struct {
	XsiType              *string           `xml:"type,attr"`
	XmiId                *string           `xml:"id,attr"`
	Name                 *string           `xml:"name,attr"`
	Definition           *string           `xml:"definition,attr"`
	RegistrationStatus   *string           `xml:"registrationStatus,attr"`
	SubType              *string           `xml:"subType,attr"`
	DerivationComponent  *string           `xml:"derivationComponent,attr"`
	AssociationDomain    *string           `xml:"associationDomain,attr"`
	DerivationElement    *string           `xml:"derivationElement,attr"`
	ListOfElement        []*Element        `xml:"element"`
	ListOfMessageElement []*MessageElement `xml:"messageElement"`
	ListOfSemanticMarkup []*SemanticMarkup `xml:"semanticMarkup"`
	ListOfCode           []*Code           `xml:"code"`
	ListOfConstraint     []*Constraint     `xml:"constraint"`
	ListOfXors           []*Xors           `xml:"xors"`
	ListOfExamples       []*Example        `xml:"example"`
	ListOfNamespaceList  []*NamespaceList  `xml:"namespaceList"`
	MinInclusive         *float64          `xml:"minInclusive,attr"`
	TotalDigits          *int              `xml:"totalDigits,attr"`
	FractionDigits       *int              `xml:"fractionDigits,attr"`
	MinLength            *int              `xml:"minLength,attr"`
	MaxLength            *int              `xml:"maxLength,attr"`
	Pattern              *string           `xml:"pattern,attr"`
}

type Example struct {
	Value *string `xml:",chardata"`
}

type NamespaceList struct {
	Value *string `xml:",chardata"`
}

type Doclet struct {
	XmiId    *string `xml:"id,attr"`
	Type     *string `xml:"type,attr"`
	Content  *string `xml:"content,attr"`
	Language *string `xml:"language,attr"`
}

type MessageBuildingBlock struct {
	XmiId                *string           `xml:"id,attr"`
	NextVersions         *string           `xml:"nextVersions,attr"`
	PreviousVersion      *string           `xml:"previousVersion,attr"`
	Name                 *string           `xml:"name,attr"`
	Definition           *string           `xml:"definition,attr"`
	RegistrationStatus   *string           `xml:"registrationStatus,attr"`
	MinOccurs            *int              `xml:"minOccurs,attr"`
	MaxOccurs            *int              `xml:"maxOccurs,attr"`
	XmlTag               *string           `xml:"xmlTag,attr"`
	ComplexType          *string           `xml:"complexType,attr"`
	ListOfSemanticMarkup []*SemanticMarkup `xml:"semanticMarkup"`
}

type MessageDefinition struct {
	XmiId                             *string                        `xml:"id,attr"`
	PreviousVersion                   *string                        `xml:"previousVersion,attr"`
	Name                              *string                        `xml:"name,attr"`
	Definition                        *string                        `xml:"definition,attr"`
	RegistrationStatus                *string                        `xml:"registrationStatus,attr"`
	MessageSet                        *string                        `xml:"messageSet,attr"`
	XmlTag                            *string                        `xml:"xmlTag,attr"`
	RootElement                       *string                        `xml:"rootElement,attr"`
	ListOfMessageBuildingBlock        []*MessageBuildingBlock        `xml:"messageBuildingBlock"`
	ListOfSemanticMarkup              []*SemanticMarkup              `xml:"semanticMarkup"`
	ListOfMessageDefinitionIdentifier []*MessageDefinitionIdentifier `xml:"messageDefinitionIdentifier"`
	ListOfConstraint                  []*Constraint                  `xml:"constraint"`
	ListOfXors                        []*Xors                        `xml:"xors"`
}

type Element struct {
	XsiType              *string           `xml:"xsi:type,attr"`
	XmiId                *string           `xml:"xmi:id,attr"`
	Name                 *string           `xml:"name,attr"`
	Definition           *string           `xml:"definition,attr"`
	RegistrationStatus   *string           `xml:"registrationStatus,attr"`
	MinOccurs            *int              `xml:"minOccurs,attr"`
	MaxOccurs            *int              `xml:"maxOccurs,attr"`
	IsDerived            *bool             `xml:"isDerived,attr"`
	Derivation           *string           `xml:"derivation,attr"`
	Opposite             *string           `xml:"opposite,attr"`
	Type                 *string           `xml:"type,attr"`
	SimpleType           *string           `xml:"simpleType,attr"`
	ComplexType          *string           `xml:"complexType,attr"`
	ListOfSemanticMarkup []*SemanticMarkup `xml:"semanticMarkup"`
}

type MessageElement struct {
	XsiType              *string           `xml:"xsi:type,attr"`
	XmiId                *string           `xml:"xmi:id,attr"`
	Name                 *string           `xml:"name,attr"`
	Definition           *string           `xml:"definition,attr"`
	RegistrationStatus   *string           `xml:"registrationStatus,attr"`
	MinOccurs            *int              `xml:"minOccurs,attr"`
	MaxOccurs            *int              `xml:"maxOccurs,attr"`
	IsDerived            *bool             `xml:"isDerived,attr"`
	ComplexType          *string           `xml:"complexType,attr"`
	SimpleType           *string           `xml:"simpleType,attr"`
	Type                 *string           `xml:"type,attr"`
	BusinessElementTrace *string           `xml:"businessElementTrace,attr"`
	XmlTag               *string           `xml:"xmlTag,attr"`
	ListOfSemanticMarkup []*SemanticMarkup `xml:"semanticMarkup"`
	ListOfConstraint     []*Constraint     `xml:"constraint"`
}

type BusinessProcessCatalogue struct {
	ListOfTopLevelCatalogueEntries []*TopLevelCatalogueEntry `xml:"topLevelCatalogueEntry"`
}

type TopLevelCatalogueEntry struct {
	XsiType                 *string              `xml:"type,attr"`
	XmiId                   *string              `xml:"id,attr"`
	Name                    *string              `xml:"name,attr"`
	Definition              *string              `xml:"definition,attr"`
	RegistrationStatus      *string              `xml:"registrationStatus,attr"`
	ListOfMessageDefinition []*MessageDefinition `xml:"messageDefinition"`
	ListOfBusinessRoles     []*BusinessRole      `xml:"businessRole"`
	ListOfDoclet            []*Doclet            `xml:"doclet"`
}
type BusinessRole struct {
	XmiId                *string           `xml:"id,attr"`
	Name                 *string           `xml:"name,attr"`
	Definition           *string           `xml:"definition,attr"`
	RegistrationStatus   *string           `xml:"registrationStatus,attr"`
	ListOfSemanticMarkup []*SemanticMarkup `xml:"semanticMarkup"`
}
type SemanticMarkup struct {
	XmiId          *string     `xml:"id,attr"`
	Type           *string     `xml:"type,attr"`
	ListOfElements []*Elements `xml:"elements"`
}

type Elements struct {
	XmiId *string `xml:"id,attr"`
	Name  *string `xml:"name,attr"`
	Value *string `xml:"value,attr"`
}

type Code struct {
	XmiId                *string           `xml:"id,attr"`
	Name                 *string           `xml:"name,attr"`
	Definition           *string           `xml:"definition,attr"`
	RegistrationStatus   *string           `xml:"registrationStatus,attr"`
	CodeName             *string           `xml:"codeName,attr"`
	ListOfSemanticMarkup []*SemanticMarkup `xml:"semanticMarkup"`
}

type MessageDefinitionIdentifier struct {
	BusinessArea         *string `xml:"businessArea,attr"`
	MessageFunctionality *string `xml:"messageFunctionality,attr"`
	Flavour              *string `xml:"flavour,attr"`
	Version              *string `xml:"version,attr"`
}

type Constraint struct {
	XmiId              *string `xml:"id,attr"`
	NextVersions       *string `xml:"nextVersions,attr"`
	PreviousVersion    *string `xml:"previousVersion,attr"`
	Name               *string `xml:"name,attr"`
	Definition         *string `xml:"definition,attr"`
	RegistrationStatus *string `xml:"registrationStatus,attr"`
}

type Xors struct {
	XmiId                         *string `xml:"id,attr"`
	Name                          *string `xml:"name,attr"`
	Definition                    *string `xml:"definition,attr"`
	RegistrationStatus            *string `xml:"registrationStatus,attr"`
	ImpactedMessageBuildingBlocks *string `xml:"impactedMessageBuildingBlocks,attr"`
}

func PrintElementAttributes(element *Element) {
	fmt.Println("XsiType:", val(element.XsiType))
	fmt.Println("XmiId:", val(element.XmiId))
	fmt.Println("Name:", val(element.Name))
	fmt.Println("Definition:", val(element.Definition))
	fmt.Println("RegistrationStatus:", val(element.RegistrationStatus))
	fmt.Println("MinOccurs:", element.MinOccurs)
	fmt.Println("MaxOccurs:", element.MaxOccurs)
	fmt.Println("IsDerived:", element.IsDerived)
	fmt.Println("Derivation:", val(element.Derivation))
	fmt.Println("Opposite:", val(element.Opposite))
	fmt.Println("Type:", val(element.Type))
	fmt.Println("SimpleType:", val(element.SimpleType))
}

func val(str *string) string {
	if str == nil {
		return "<nil>"
	} else {
		return *str
	}
}

type BasicElement struct {
	Id          *string         `json:"id,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Children    []*BasicElement `json:"children,omitempty"`
	Type        *string         `json:"type,omitempty"`
	Required    bool            `json:"required"`
	Attribute   bool            `json:"-"`
	MaxLength   *int            `json:"maxLength,omitempty"`
	MinLength   *int            `json:"minLength,omitempty"`
	MinValue    *float64        `json:"minValue,omitempty"`
	Pattern     *string         `json:"pattern,omitempty"`
	MaxOccurs   *int            `json:"maxOccurs,omitempty"`
	MinOccurs   *int            `json:"minOccurs,omitempty"`
	Array       bool            `json:"array"`
}

func (e *BasicElement) DuplicateNoChildren() *BasicElement {
	var element *BasicElement

	element = &BasicElement{
		Id:          e.Id,
		Name:        e.Name,
		Description: e.Description,
		Children:    []*BasicElement{},
		Type:        e.Type,
		Required:    e.Required,
		Attribute:   e.Attribute,
		MaxLength:   e.MaxLength,
		MinLength:   e.MinLength,
		MinValue:    e.MinValue,
		Pattern:     e.Pattern,
		MaxOccurs:   e.MaxOccurs,
		MinOccurs:   e.MinOccurs,
		Array:       e.Array,
	}
	element.Assess()

	return element
}

func (e *BasicElement) AddChild(child *BasicElement) {
	e.Children = append(e.Children, child)
}

func (entry *TopLevelDictionaryEntry) ToElement() *BasicElement {
	var element *BasicElement

	element = &BasicElement{
		Id:          entry.XmiId,
		Name:        entry.Name,
		Description: substNewLines(entry.Definition),
		Children:    []*BasicElement{},
		Type:        nil,
		Required:    false,
		Attribute:   false,
		MaxLength:   entry.MaxLength,
		MinLength:   entry.MinLength,
		MinValue:    entry.MinInclusive,
		Pattern:     entry.Pattern,
		MaxOccurs:   nil,
		MinOccurs:   nil,
		Array:       false,
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
	element.Assess()

	return element
}

func (e *MessageElement) ToElement() *BasicElement {
	var element *BasicElement

	element = &BasicElement{
		Name:        e.Name,
		Description: substNewLines(e.Definition),
		Children:    []*BasicElement{},
		Type:        nil,
		Required:    false,
		Attribute:   false,
		MaxLength:   nil,
		MinLength:   nil,
		MinValue:    nil,
		Pattern:     nil,
		MaxOccurs:   e.MaxOccurs,
		MinOccurs:   e.MinOccurs,
		Array:       false,
	}
	if e.SimpleType != nil {
		element.Id = e.SimpleType
	} else if e.ComplexType != nil {
		element.Id = e.ComplexType
	} else if e.Type != nil {
		element.Id = e.Type
	}
	element.Assess()

	return element
}

func (e *Element) ToElement() *BasicElement {
	var element *BasicElement

	element = &BasicElement{
		Id:          e.XmiId,
		Name:        e.Name,
		Description: substNewLines(e.Definition),
		Children:    []*BasicElement{},
		Type:        nil,
		Required:    false,
		Attribute:   false,
		MaxLength:   nil,
		MinLength:   nil,
		MinValue:    nil,
		Pattern:     nil,
		MaxOccurs:   e.MaxOccurs,
		MinOccurs:   e.MinOccurs,
		Array:       false,
	}
	if e.SimpleType != nil {
		element.Id = e.SimpleType
	} else if e.ComplexType != nil {
		element.Id = e.ComplexType
	} else if e.Type != nil {
		element.Id = e.Type
	}
	element.Assess()

	return element
}

func (e *BasicElement) Assess() {
	if e.MinOccurs != nil && *e.MinOccurs > 0 {
		e.Required = true
	}
	if e.MaxOccurs != nil && *e.MaxOccurs > 1 {
		e.Array = true
	}
	if e.MaxOccurs == nil {
		e.Array = true
	}
}

func substNewLines(str *string) *string {
	if str == nil {
		return nil
	}
	str2 := strings.ReplaceAll(*str, "\r", "")
	str2 = strings.ReplaceAll(str2, "\n", "<p/>")
	return &str2
}

type CatalogueEntry struct {
	Name              *string `json:"name,omitempty"`
	Description       *string `json:"description,omitempty"`
	MessageName       *string `json:"messageName,omitempty"`
	MessageDefinition *string `json:"messageDefinition,omitempty"`
	Domain            *string `json:"domain,omitempty"`
	FunctionalArea    *string `json:"functionalArea,omitempty"`
}

func (e *TopLevelCatalogueEntry) ToCatalogueEntries() []CatalogueEntry {
	var entries []CatalogueEntry

	for _, messageDefinition := range e.ListOfMessageDefinition {
		if messageDefinition.ListOfMessageDefinitionIdentifier != nil {
			for _, messageDefinitionIdentifier := range messageDefinition.ListOfMessageDefinitionIdentifier {
				entry := &CatalogueEntry{
					Name:              e.Name,
					Description:       e.Definition,
					MessageName:       nil,
					MessageDefinition: nil,
					Domain:            nil,
					FunctionalArea:    nil,
				}
				entry.MessageName = messageDefinition.Name
				entry.MessageDefinition = messageDefinition.Definition
				entry.Domain = messageDefinitionIdentifier.BusinessArea
				entry.FunctionalArea = messageDefinitionIdentifier.MessageFunctionality
				entries = append(entries, *entry)
			}
		}
	}

	return entries
}
