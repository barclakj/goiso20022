package repo

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"realizr.io/iso20022/model"
)

func TestReadXMLFile(t *testing.T) {
	filePath := "https://storage.googleapis.com/media.nonfunctionalarchitect.com/20230719_ISO20022_2013_eRepository.iso20022"

	// Call the function being tested
	iso20022model, err := ReadXMLFile(filePath)

	if err != nil {
		t.Errorf("Error reading XML file: %v", err)
	}

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

	el := ExpandElement("OrganisationModification2", iso20022model, nil)
	assert.NotNil(t, el)
	fmt.Printf("Element: %v = %v\n", *el.Name, *el.Description)

	// Iterate over el children and print each child
	for _, child := range el.Children {
		fmt.Printf("Child: %v - %v - %v\n", *child.Name, *child.Description, child.Required)
	}
	assert.Equal(t, "OrganisationModification2", *el.Name)

	el = ExpandElement("CardPaymentTransaction53", iso20022model, nil)
	mandatoryEl := FilterMandatoryElements(el)
	assert.NotNil(t, mandatoryEl)
	assert.True(t, len(mandatoryEl.Children) > 0)
	for _, child := range mandatoryEl.Children {
		assert.True(t, child.Required)
		fmt.Printf("Mandatory Child: %v - %v - %v\n", *child.Name, *child.Description, child.Required)
	}

	fmt.Printf("Fetching element CardPaymentTransactionDetails42 for underlying type testing\n")
	el = ExpandElement("CardPaymentTransactionDetails42", iso20022model, nil)
	el = FilterMandatoryElements(el)
	assert.NotNil(t, el)
	assert.True(t, len(el.Children) > 0)
	for _, child := range el.Children {
		assert.True(t, child.Required)
		fmt.Printf("Mandatory Child: %v - %v - %v\n", *child.Name, *child.Description, *child.Type)
		if *child.Name == "TotalAmount" {
			assert.Equal(t, "double", *child.Type)
		}
	}

	testCatalogue(t, iso20022model)

	listCatalogue(t, iso20022model)

	testMaxOccurs(t, iso20022model)

	testNoLoop(t, iso20022model)
}

func testMaxOccurs(t *testing.T, mdl *model.Iso20022) {
	postlAddress := ExpandElement("PostalAddress4", mdl, nil)
	assert.NotNil(t, postlAddress)

	// Find the AddressLine element
	var addressLine *Element
	for _, element := range postlAddress.Children {
		if *element.Name == "AddressLine" {
			addressLine = element
			log.Default().Printf("Found Address Line: %v\n", *addressLine.Name)
			log.Default().Printf("Found Max Occurs: %v\n", *addressLine.MaxOccurs)
			log.Default().Printf("Found Required: %v\n", addressLine.Required)

			break
		}
	}

	// Verify the maxOccurs is 2
	assert.NotNil(t, addressLine)
	assert.Equal(t, 2, *addressLine.MaxOccurs)
	assert.Equal(t, false, addressLine.Required)
}

func testNoLoop(t *testing.T, mdl *model.Iso20022) {
	postlAddress := ExpandElement("PostalAddress", mdl, nil)
	assert.NotNil(t, postlAddress)
}

func testCatalogue(t *testing.T, model *model.Iso20022) {
	catalogue := ExpandCatalogue("CustomerCreditTransferInitiationV02", model)
	assert.NotNil(t, catalogue)
}

func listCatalogue(t *testing.T, model *model.Iso20022) {
	catalogueEntries := ListCatalogue(model)
	assert.NotNil(t, catalogueEntries)
	assert.True(t, len(*catalogueEntries) > 0)
	for _, entry := range *catalogueEntries {
		assert.NotNil(t, entry.Name)
		fmt.Printf("Name: %v\n", *entry.Name)
		if entry.Description != nil {
			fmt.Printf("Description: %v\n", *entry.Description)
		}
		assert.NotNil(t, entry.MessageName)
		fmt.Printf("MessageName: %v\n", *entry.MessageName)
		if entry.MessageDefinition != nil {
			fmt.Printf("MessageDefinition: %v\n", *entry.MessageDefinition)
		}
		assert.NotNil(t, entry.FunctionalArea)
		fmt.Printf("FunctionalArea: %v\n", *entry.FunctionalArea)
		assert.NotNil(t, entry.Domain)
		fmt.Printf("Domain: %v\n", *entry.Domain)
	}
}
