package handlers

// Stores the name, real name and description of an Avenger
type Avenger struct {
	Name        string `json:"name"`
	RealName    string `json:"realname"`
	Description string `json:"desc"`
}

// Dummy data to test API
var db []Avenger = []Avenger{
	{
		Name:        "Iron Man",
		RealName:    "Tony Stark",
		Description: "A genius millionare who creates battle suits with high tech weaponry",
	},
	{
		Name:        "Captain America",
		RealName:    "Steve Rogers",
		Description: "A soldier with superhuman capabilities and strength",
	},
}
