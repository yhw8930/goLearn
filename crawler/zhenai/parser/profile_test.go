package parser

import (
	"goLearn/crawler/engine"
	"goLearn/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "http://album.zhenai.com/u/1794937843", "雪花")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	actual := result.Items[0]
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1794937843",
		Type: "zhenai",
		Id:   "1794937843",
		Payload: model.Profile{
			Name:      "雪花",
			Age:       26,
			Height:    161,
			Weight:    53,
			Income:    "8千-1.2万",
			Residence: "湖北武汉",
			House:     "武汉武昌区",
		},
	}

	if actual != expected {
		t.Errorf("%v==%v", actual, expected)
	}
}
