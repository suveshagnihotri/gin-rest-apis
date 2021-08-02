package Models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Password string `json:"password"`
	Address string `json:"address"`
}

func (b *User) TableName() string {
	return "user"
}
