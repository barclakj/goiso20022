package repo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
		fmt.Printf("Child: %v - %v - %v\n", *child.Name, *child.Description, *&child.Required)
	}
	assert.Equal(t, "OrganisationModification2", *el.Name)

	el = ExpandElement("CardPaymentTransaction53", iso20022model, nil)
	mandatoryEl := FilterMandatoryElements(el)
	assert.NotNil(t, mandatoryEl)
	assert.True(t, len(mandatoryEl.Children) > 0)
	for _, child := range mandatoryEl.Children {
		assert.True(t, child.Required)
		fmt.Printf("Mandatory Child: %v - %v - %v\n", *child.Name, *child.Description, *&child.Required)
	}
}
