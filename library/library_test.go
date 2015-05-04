package library

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
)

//Create a file with some json in it, have the function read it, then
//destroy the file to clean up.
func TestCreateResourceFromJSON(t *testing.T) {
	var buf bytes.Buffer

	buf.WriteString("{ \"name\": \"cats.jpg\", \"description\": \"cat picture of cats\", ")
	buf.WriteString("\"location\": \"/tmp/somepath\", \"mimetype\": \"image/jpeg\"}")

	buffer := buf.Bytes()
	//t.Logf("Testing the json buffer: %v", buffer)

	result, err := CreateResourceFromJSON(buffer)
	if err != nil {
		t.Errorf("Error creating resource from JSON string: %v", err)
	}

	t.Logf("Resource was decoded as:\n %v", result)

	buffer = make([]byte, 0)
	_, err = CreateResourceFromJSON(buffer)
	if err == nil {
		t.Errorf("CreateResourceFromJSON should return an error for zero-length slice")
	}

	buffer = nil
	_, err = CreateResourceFromJSON(buffer)
	if err == nil {
		t.Errorf("CreateResourceFromJSON should return an error for nil slice")
	}
}

func TestLoadResourcesFrom(t *testing.T) {
	var buf bytes.Buffer

	path := "/tmp/testfile"

	buf.WriteString("[{ \"name\": \"cats.jpg\", \"description\": \"cat picture of cats\", ")
	buf.WriteString("\"location\": \"/tmp/somepath\", \"mimetype\": \"image/jpeg\"}]")

	ioutil.WriteFile(path, buf.Bytes(), 0644)

	_, err := LoadResourcesFrom(path)
	if err != nil {
		t.Errorf("Error loading resources from file: %v", err)
	}

	os.Remove(path) //clean up the test file

}

func TestAssignRandomIDs(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetLevel(logrus.DebugLevel)

	path := "/tmp/testfile"

	buf.WriteString("[{ \"name\": \"cats.jpg\", \"description\": \"cat picture of cats\", ")
	buf.WriteString("\"location\": \"/tmp/somepath\", \"mimetype\": \"image/jpeg\"}]")

	ioutil.WriteFile(path, buf.Bytes(), 0644)

	slice, err := LoadResourcesFrom(path)
	if err != nil {
		t.Errorf("Error loading resources from file: %v", err)
	}

	os.Remove(path) //clean up the test file

	err = AssignRandomIDs(slice)
	if err != nil {
		t.Errorf("Error assigning random IDs: %v", err)
	}

	if slice[0].resourceID == nil {
		t.Error("Resource ID was not set")
	}

}
