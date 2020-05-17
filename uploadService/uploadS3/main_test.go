package main

import (
	"encoding/json"
	"testing"
)

func TestUpload(t *testing.T) {
	resp, err := Handler(nil, Reqeust{})

	if resp.StatusCode != 200 {
		t.Errorf("Status code did not equal 200")
	}

	if err != nil {
		t.Errorf("Error occured")
	}

	var body map[string]interface{}

	json.Unmarshal([]byte(resp.Body), &body)

	if body["message"] != "Okay" {
		t.Errorf("Wrong message, was %v", body["message"])
	}
}
