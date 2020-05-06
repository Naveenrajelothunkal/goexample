package userModel

type UserDetails struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Place      string `json:"place"`
	Created_On string `json:"created_on"`
	Updated_On string `json:"updated_on"`
}
