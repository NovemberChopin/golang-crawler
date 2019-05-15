package model

import "encoding/json"

type Profile struct {
	Name     string
	Gender   string
	Age      int
	Height   int
	Weight   int
	Income   string
	Marriage string
	Address  string
}

func FromJsonObj(obj interface{}) (Profile, error) {
	var profile Profile
	string, err := json.Marshal(obj)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(string, &profile)
	return profile, err
}
