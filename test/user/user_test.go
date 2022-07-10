package user_test

import (
	"testing"

	"github.com/hermansetiawan77777/technical-warungpintar/internal/storage/user"
)

func TestCreateUser(t *testing.T) {
	payload := &user.User{
		Id:       "Hesrti",
		Name:     "Herstian",
		Password: "123",
	}
	newuser := user.AddNewUser(payload.Id, payload.Name, payload.Password)
	if newuser == nil {
		t.Errorf("Failed Crete new user")
		return
	}
}

//please run add function "NewUser" first before "LoginUser"
func TestLoginUser(t *testing.T) {
	loginuser := user.LoginUser("Hesrti", "123")
	if loginuser == nil {
		t.Errorf("Wrong ID or Password !")
		return
	}
}
