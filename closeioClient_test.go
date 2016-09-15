package closeio

import "testing"

func Test_ConvertQuery(t *testing.T) {

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
