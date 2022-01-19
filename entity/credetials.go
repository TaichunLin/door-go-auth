package entity

type TokenMetadata struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type Accounts struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password []byte `json:"password"`
}

type AccessDetails struct {
	AccessUuid string
	Email      string
}

type Todo struct {
	Id      string `json:"id"`
	Account string `json:"account"`
	Title   string `json:"title"`
}
