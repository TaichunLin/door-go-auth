package entity

type Group struct {
	Group   string `json:"group"`
	GroupId string `json:"groupId"`
	Door    string `json:"door"`
}

type User struct {
	Username string `json:"username"`
	CardId   string `json:"cardId"`
	GroupId  *Group `json:"groupId"`
}

type UserList struct {
	Username string `json:"username"`
	CardId   string `json:"cardId"`
	Group    string `json:"group"`
	GroupId  string `json:"groupId"`
}
