package main

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestParsing(t *testing.T) {
	payload := generatePayload("https://youtu.be/testid1?t=465")

	var result map[string]interface{}
	err := json.Unmarshal([]byte(payload), &result)
	if err != nil {
		t.Error("Invalid json")
	}

	resumeValues := result["params"].(map[string]interface{})["options"].(map[string]interface{})["resume"].(map[string]interface{})

	if resumeValues["hours"].(float64) != 0 || resumeValues["minutes"].(float64) != 7 || resumeValues["seconds"].(float64) != 30 {
		t.Error("Invalid parsed time")
	}

	urlString := result["params"].(map[string]interface{})["item"].(map[string]interface{})["file"].(string)
	url, err := url.Parse(urlString)
	if err != nil {
		t.Error("Invalid url")
	}

	if url.Query().Get("videoid") != "testid1" {
		t.Error("Invalid video id")
	}
}
