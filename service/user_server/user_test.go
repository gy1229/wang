package user_server

import (
	"fmt"
	"github.com/gy1229/wang/json_struct"
	"github.com/gy1229/wang/test"
	"testing"
)

func TestLoadUserMessage(t *testing.T) {
	test.InitTestConfig2()
	resp, _ := InsertUserMessage(&json_struct.RegisterRequest{UserBase:&json_struct.UserBase{
		AvatarId:"1",
		Account:  "wang",
		Password: "wang",
	},
	UserName:"hell",
	ImageId:"123",
	QQ:"123",
	Tel:"32",
	Email:"wang",
	})
	fmt.Println(resp)
}

func TestCertainAccount(t *testing.T) {
	test.InitTestConfig2()
	resp, _ := CertainAccount(&json_struct.CertainAccountRequest{Account: "guyu"})
	fmt.Println(resp)
}
