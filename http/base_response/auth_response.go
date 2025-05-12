package baseresponse

type AuthResponse struct {
	AccessToken string `json:"token"`
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
}
