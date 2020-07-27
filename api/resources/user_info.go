package resources

type UserInfoRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
