package user

import (
	"fmt"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/test"
	"testing"
)

func TestLoadUserMessage(t *testing.T) {
	test.InitTestConfig2()
	resp, _ := LoadUserMessage(&json_struct.LoadUserMessageRequest{UserId: "123"})
	fmt.Println(resp)
}

func TestCertainAccount(t *testing.T) {
	test.InitTestConfig2()
	resp, _ := CertainAccount(&json_struct.CertainAccountRequest{Account: "guyu"})
	fmt.Println(resp)
}
