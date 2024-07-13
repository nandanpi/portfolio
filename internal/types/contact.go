package types

type AddContactReq struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
