package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Salt     string `db:"salt"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `db:"salt"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Salt     string `db:"salt"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (user *User) ToResponse() UserResponse {
	return UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}
}
