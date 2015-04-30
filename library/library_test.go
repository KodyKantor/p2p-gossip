package library

import (
	"bytes"
	"testing"
)

//Create a file with some json in it, have the function read it, then
//destroy the file to clean up.
func TestCreateResourceFromJSON(t *testing.T) {
	var buf bytes.Buffer

	buf.WriteString("{ \"name\": \"cats.jpg\", \"description\": \"cat picture of cats\", ")
	buf.WriteString("\"location\": \"/tmp/somepath\", \"mimetype\": \"image/jpeg\"}")

	str := buf.String()
	t.Logf("Testing the json string: %v", str)

	result, err := CreateResourceFromJSON(str)
	if err != nil {
		t.Errorf("Error creating resource from JSON string: %v", err)
	}

	t.Logf("Resource was decoded as:\n %v", result)

	str = ""
	_, err = CreateResourceFromJSON(str)
	if err == nil {
		t.Errorf("CreateResourceFromJSON should return an error for zero-length string")
	}
}
