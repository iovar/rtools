package tools

import "testing"

func TestJSONMinify(t *testing.T) {
	testJSON := `{
		"propA": "valueA",
		"propB": 2,
		"propC": [ true, true, false ]
	}`
	want := `{"propA":"valueA","propB":2,"propC":[true,true,false]}`

	res := JSONMinify(testJSON)

	if res != want {
		t.Errorf("JSON output %s did not match expectation: %s", res, want)
	}
}

func TestJSONBeautify(t *testing.T) {
	testJSON := `{"propA":"valueA","propB":2,"propC":[true,true,false]}`
	want := "{\n    \"propA\": \"valueA\",\n    \"propB\": 2,\n    \"propC\": [\n        true,\n        true,\n        false\n    ]\n}"

	res := JSONBeautify(testJSON)

	if res != want {
		t.Errorf("JSON output %s did not match expectation: %s", res, want)
	}
}
