package repo

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"realizr.io/iso20022/model"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

var iso20022model *model.Iso20022

func setup() {
	filePath := "https://storage.googleapis.com/media.nonfunctionalarchitect.com/20230719_ISO20022_2013_eRepository.iso20022"
	var err error
	// Call the function being tested
	iso20022model, err = ReadXMLFile(filePath)
	if err != nil {
		log.Printf("Error reading XML file: %v", err)
		iso20022model = nil
	}
}

func TestReadXMLFile(t *testing.T) {
	assert.NotNil(t, iso20022model)

	// Loop through toplevelcatalogueentries
	for _, entry := range iso20022model.BusinessProcessCatalogue.ListOfTopLevelCatalogueEntries {
		// Process each entry
		// ...
		assert.NotNil(t, entry.Name)
		if *entry.Name == "Organisation31" {
			fmt.Printf("ProcessCatalogue: %v = %v\n", entry.Name, entry.Definition)
		}
	}
}

func TestExpandElement(t *testing.T) {
	el := ExpandElement("OrganisationModification2", iso20022model, nil)
	assert.NotNil(t, el)
	fmt.Printf("Element: %v = %v\n", *el.Name, *el.Description)

	// Iterate over el children and print each child
	for _, child := range el.Children {
		assert.NotNil(t, child.Name)
		// fmt.Printf("Child: %v - %v - %v\n", *child.Name, *child.Description, child.Required)
	}
	assert.Equal(t, "OrganisationModification2", *el.Name)

	el = ExpandElement("CardPaymentTransaction53", iso20022model, nil)
	mandatoryEl := FilterMandatoryElements(el)
	assert.NotNil(t, mandatoryEl)
	assert.True(t, len(mandatoryEl.Children) > 0)
	for _, child := range mandatoryEl.Children {
		assert.True(t, child.Required)
		// fmt.Printf("Mandatory Child: %v - %v - %v\n", *child.Name, *child.Description, child.Required)
	}

	fmt.Printf("Fetching element CardPaymentTransactionDetails42 for underlying type testing\n")
	el = ExpandElement("CardPaymentTransactionDetails42", iso20022model, nil)
	el = FilterMandatoryElements(el)
	assert.NotNil(t, el)
	assert.True(t, len(el.Children) > 0)
	for _, child := range el.Children {
		assert.True(t, child.Required)
		// fmt.Printf("Mandatory Child: %v - %v - %v\n", *child.Name, *child.Description, *child.Type)
		if *child.Name == "TotalAmount" {
			assert.Equal(t, "double", *child.Type)
		}
	}
}

func TestMaxOccurs(t *testing.T) {
	postlAddress := ExpandElement("PostalAddress4", iso20022model, nil)
	assert.NotNil(t, postlAddress)
	log.Printf("Postal address has %v children\n", len(postlAddress.Children))

	// Find the AddressLine element
	var addressLine *model.BasicElement
	for _, element := range postlAddress.Children {
		log.Printf("Child: %v\n", *element.Name)
		if *element.Name == "AddressLine" {
			addressLine = element
			log.Printf("Found Address Line: %v\n", *addressLine.Name)
			log.Printf("Found Max Occurs: %v\n", *addressLine.MaxOccurs)
			log.Printf("Found Required: %v\n", addressLine.Required)

			break
		}
	}

	// Verify the maxOccurs is 2
	assert.NotNil(t, addressLine)
	assert.Equal(t, 2, *addressLine.MaxOccurs)
	assert.Equal(t, false, addressLine.Required)
}

func TestNoLoop(t *testing.T) {
	postlAddress := ExpandElement("PostalAddress", iso20022model, nil)
	assert.NotNil(t, postlAddress)
}

// func TestCatalogue(t *testing.T) {
// catalogue := ExpandCatalogue("CustomerCreditTransferInitiationV02", model)
// assert.NotNil(t, catalogue)
// }

func TestListCatalogue(t *testing.T) {
	catalogueEntries := ListCatalogue(iso20022model)
	assert.NotNil(t, catalogueEntries)
	assert.True(t, len(*catalogueEntries) > 0)
	for _, entry := range *catalogueEntries {
		assert.NotNil(t, entry.Name)

		// fmt.Printf("Name: %v\n", *entry.Name)

		assert.NotNil(t, entry.MessageName)
		// fmt.Printf("MessageName: %v\n", *entry.MessageName)
		// if entry.MessageDefinition != nil {
		// 	fmt.Printf("MessageDefinition: %v\n", *entry.MessageDefinition)
		// }
		assert.NotNil(t, entry.FunctionalArea)
		// fmt.Printf("FunctionalArea: %v\n", *entry.FunctionalArea)
		assert.NotNil(t, entry.Domain)
		// fmt.Printf("Domain: %v\n", *entry.Domain)
	}
}
