package models

type UsersModel struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	IdentificationNumber string `json:"identification_number"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`	
}
