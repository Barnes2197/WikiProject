package main

import (
	"bytes"
	"os"
	"testing"
)

func TestPage(t *testing.T) {
	testPage := &Page{Title: "testPage", Body: []byte("This is a test")}
	err := testPage.save()
	if err != nil {
		t.Error("Page saving failed")
		return
	}
	result, err := loadPage("testPage")
	if err != nil {
		t.Error("Page loading failed")
		return
	}
	match := bytes.Compare(result.Body, testPage.Body)
	if match != 0 && result.Title != testPage.Title {
		t.Error("Result and test page don't match")
	}

	_ = os.Remove("data/" + testPage.Title + ".txt")

}
