package model

import (
	"encoding/xml"
	"fmt"
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
	XmiId                *string           `xml:"id,attr"`
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
	XmiId                *string           `xml:"id,attr"`
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
	fmt.Println("XsiType:", element.XsiType)
	fmt.Println("XmiId:", element.XmiId)
	fmt.Println("Name:", element.Name)
	fmt.Println("Definition:", element.Definition)
	fmt.Println("RegistrationStatus:", element.RegistrationStatus)
	fmt.Println("MinOccurs:", element.MinOccurs)
	fmt.Println("MaxOccurs:", element.MaxOccurs)
	fmt.Println("IsDerived:", element.IsDerived)
	fmt.Println("Derivation:", element.Derivation)
	fmt.Println("Opposite:", element.Opposite)
	fmt.Println("Type:", element.Type)
	fmt.Println("SimpleType:", element.SimpleType)
}
