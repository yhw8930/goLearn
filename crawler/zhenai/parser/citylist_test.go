package parser

import (
	"io/ioutil"
	"testing"
)

const resultSize = 470

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	if len(result.Items) != resultSize || len(result.Requests) != resultSize {
		t.Errorf("Result should have %d "+"but had %d  %d", resultSize, len(result.Items), len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("Expected url #%d: %s; but "+"was %s", i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("Expected city #%d: %s; but "+"was %s", i, city, result.Items[i].(string))
		}
	}
}