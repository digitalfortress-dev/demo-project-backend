package user

type (
	User struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		AccessToken string `json:"access_token"`
	}
)
