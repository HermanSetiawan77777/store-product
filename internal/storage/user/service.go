package user

import (
	"crypto/rand"
	"fmt"
)

var Users []*User

// normally create .env and put "KEY" for generated Token
func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
func AddNewUser(id, name, password string) *User {
	cek := checkduplicated(id)
	if !cek {
		return nil
	}
	newUser := &User{
		Id:       id,
		Name:     name,
		Password: password,
	}
	Users = append(Users, newUser)
	return newUser
}

func LoginUser(id, password string) *User {
	for _, u := range Users {
		if u.Id == id {
			if u.Password == password {
				u.LoginToken = tokenGenerator()
				return u
			}
			return nil
		}
	}
	return nil
}

func GetUserById(id string) *User {
	for _, u := range Users {
		if u.Id == id {
			return u
		}
	}
	return nil
}

func checkduplicated(newid string) bool {
	existingUser := GetUserById(newid)
	for existingUser != nil {
		return false
	}

	return true
}

func Gettokendata(token string) *User {
	for _, u := range Users {
		if u.LoginToken == token {
			return u
		}
	}
	return nil
}
