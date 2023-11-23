package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadXMLFile(t *testing.T) {
	filePath := "../schema/20230719_ISO20022_2013_eRepository.iso20022"

	// Call the function being tested
	iso20022model, err := ReadXMLFile(filePath)

	if err != nil {
		t.Errorf("Error reading XML file: %v", err)
	}

	assert.NotNil(t, iso20022model)

}
