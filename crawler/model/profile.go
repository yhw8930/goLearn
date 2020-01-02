package model

import "encoding/json"

type Profile struct {
	Name      string
	Age       int
	Height    int
	Weight    int
	Income    string
	Residence string
	House     string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
