package model

import "fmt"

type Iso20022 struct {
	DataDictionary           DataDictionary
	BusinessProcessCatalogue BusinessProcessCatalogue
	Namespace                string
}

type DataDictionary struct {
	XmiId                         string
	ListOfTopLevelDictionaryEntry []TopLevelDictionaryEntry
}

type TopLevelDictionaryEntry struct {
	XsiType              string
	XmiId                string
	Name                 string
	Definition           string
	RegistrationStatus   string
	SubType              string
	DerivationComponent  string
	AssociationDomain    string
	DerivationElement    string
	ListOfElement        []Element
	ListOfMessageElement []MessageElement
	ListOfSemanticMarkup []SemanticMarkup
}

type MessageBuildingBlock struct {
	XmiId                string
	NextVersions         string
	PreviousVersion      string
	Name                 string
	Definition           string
	RegistrationStatus   string
	MinOccurs            int
	MaxOccurs            int
	XmlTag               string
	ComplexType          string
	ListOfSemanticMarkup []SemanticMarkup
}

type MessageDefinition struct {
	XmiId                      string
	PreviousVersion            string
	Name                       string
	Definition                 string
	RegistrationStatus         string
	MessageSet                 string
	XmlTag                     string
	RootElement                string
	ListOfMessageBuildingBlock []MessageBuildingBlock
	ListOfSemanticMarkup       []SemanticMarkup
}

type Element struct {
	XsiType              string
	XmiId                string
	Name                 string
	Definition           string
	RegistrationStatus   string
	MinOccurs            int
	MaxOccurs            int
	IsDerived            bool
	Derivation           string
	Opposite             string
	Type                 string
	SimpleType           string
	ListOfSemanticMarkup []SemanticMarkup
}

type MessageElement struct {
	XsiType              string
	XmiId                string
	Name                 string
	Definition           string
	RegistrationStatus   string
	MinOccurs            int
	MaxOccurs            int
	IsDerived            bool
	ComplexType          string
	BusinessElementTrace string
	XmlTag               string
	ListOfSemanticMarkup []SemanticMarkup
}

type BusinessProcessCatalogue struct {
	ListOfTopLevelCatalogueEntries []TopLevelCatalogueEntry
}

type TopLevelCatalogueEntry struct {
	XsiType                 string
	XmiId                   string
	Name                    string
	Definition              string
	RegistrationStatus      string
	ListOfMessageDefinition []MessageDefinition
	ListOfBusinessRoles     []BusinessRole
}
type BusinessRole struct {
	XmiId                string
	Name                 string
	Definition           string
	RegistrationStatus   string
	ListOfSemanticMarkup []SemanticMarkup
}
type SemanticMarkup struct {
	XmiId          string
	Type           string
	ListOfElements []Elements
}

type Elements struct {
	XmiId string
	Name  string
	Value string
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
