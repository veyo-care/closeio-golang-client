package closeio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConvertQuery(t *testing.T) {

	query := map[string][]string{

		"Custom.Field": []string{"One", "Two", "Three"},
	}

	queryString := convertQueryFields(query)

	expected := `(Custom.Field:"One" OR Custom.Field:"Two" OR Custom.Field:"Three")`

	if queryString != expected {
		t.Errorf("Expected %s, got %s", expected, queryString)
	}
}

func Test_UpdateCustomFields(t *testing.T) {

	lead := &Lead{
		Name:        "Test",
		Description: "description",
		Custom: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
	}

	content, err := LeadToJSON(*lead)
	assert.Nil(t, err)
	retrieved, err := JSONToLead(content)
	assert.Nil(t, err)
	assert.EqualValues(t, lead, retrieved)
}

// Test assumes golang maps are ordered, which they aren't

/*func Test_ConvertQuery(t *testing.T) {

	query := map[string][]string{
		"Status":       []string{"Open", "Lost"},
		"Town":         []string{"Paris"},
		"Custom.Field": []string{"One", "Two", "Three"},
	}

	queryString := convertQueryFields(query)

	expected := `(Status:"Open" OR Status:"Lost") AND (Town:"Paris") AND (Custom.Field:"One" OR Custom.Field:"Two" OR Custom.Field:"Three")`

	if queryString != expected {
		t.Errorf("Expected %s, got %s", expected, queryString)
	}
}
*/
