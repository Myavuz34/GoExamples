package dataloaders

import (
	"encoding/json"

	. "../models"
	util "../utils"
)

func LoadUsers() []User {
	bytes, _ := util.LoadFile("../json/users.json")
	var users []User
	json.Unmarshal([]byte(bytes), &users)
	return users
}

func LoadInterests() []Interest {
	bytes, _ := util.LoadFile("../json/interest.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterestMappings() []InterestMapping {
	bytes, _ := util.LoadFile("../json/userInterestMapping.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
