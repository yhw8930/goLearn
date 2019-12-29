package parser

import (
	"goLearn/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "xsd")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
	expexted := model.Profile{
		Name:      "",
		Age:       0,
		Height:    0,
		Weight:    0,
		Income:    "",
		Residence: "",
		House:     "",
	}
	if profile != expexted {
		t.Errorf("%v==%v", profile, expexted)
	}
}
