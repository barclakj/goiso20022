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
	XsiType             string
	XmiId               string
	Name                string
	Definition          string
	RegistrationStatus  string
	SubType             string
	DerivationComponent string
	AssociationDomain   string
	DerivationElement   string
	ListOfElement       []Element
}

type Element struct {
	XsiType            string
	XmiId              string
	Name               string
	Definition         string
	RegistrationStatus string
	MinOccurs          int
	MaxOccurs          int
	IsDerived          bool
	Derivation         string
	Opposite           string
	Type               string
	SimpleType         string
}

type BusinessProcessCatalogue struct {
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
