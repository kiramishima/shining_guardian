package domain

import (
	"time"
)

type User struct {
	ID        string    `json:"id" db:"id" bson:"_id,omitempty"`
	Email     string    `json:"email" db:"email" bson:"email"`
	UserName  string    `json:"username,omitempty,string" db:"username,omitempty,string" bson:"username,omitempty,string"`
	RoleID    string    `json:"-" db:"role_id,omitempty,string" bson:"role_id,omitempty,string"`
	Password  string    `json:"-" db:"password" bson:"password"`
	CreatedAt time.Time `json:"-" db:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at" bson:"updated_at"`
}

// NewUser crea un nuevo usuario
func NewUser(username, password, email string) (*User, error) {
	user := &User{
		Email:     email,
		UserName:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	/*if err := user.Validate(); err != nil {
		return nil, err
	}*/

	return user, nil
}

// Validate valida al usuario
/*func (user *User) Validate() error {
	if user.UserName == "" || user.Password == "" || user.Email == "" {
		return ErrEmptyUserField
	}

	if strings.ContainsAny(user.UserName, " \t\r\n") || strings.ContainsAny(user.Password, " \t\r\n") {
		return ErrFieldWithSpaces
	}

	if len(user.Password) < 6 {
		return ErrShortPassword
	}

	if len(user.Password) > 72 {
		return ErrLongPassword
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}
*/
