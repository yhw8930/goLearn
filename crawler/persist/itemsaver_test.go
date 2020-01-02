package persist

import (
	"context"
	"encoding/json"
	"goLearn/crawler/engine"
	"goLearn/crawler/model"
	"testing"

	"github.com/olivere/elastic"
)

func Test_save(t *testing.T) {
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

	// TODO: Try to start up elastic search
	// Here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	const index = "dating_test"
	//Save expected item
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	//Fetch saved item
	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual = engine.Item{}
	err = json.Unmarshal(resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	//Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
