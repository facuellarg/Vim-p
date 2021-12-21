package entities

//User
type User struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	Age       int        `json:"age"`
	Birthdate CustomTime `json:"birthdate"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
}
