package util

import (
	"fmt"
	"github.com/gy1229/wang/database"
	"github.com/gy1229/wang/json_struct"
	"log"
	"testing"
)

func TestGenDefaultResp(t *testing.T) {
	GenDefaultResp("cc")
}

func TestTranformStruct2GinH(t *testing.T) {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	req := json_struct.RegisterRequest{
		UserBase: &json_struct.UserBase{
			Account:  "123",
			Password: "1qwe",
		},
		UserName: "asd",
	}
	TranformStruct2GinH(req)
}

func TestGenId(t *testing.T) {
	InitID()
	fmt.Println(GenId())
}


func TestTranHttpStruct2Database(t *testing.T) {
	e := json_struct.EduEx{
		SchoolName:   "SchoolName",
		EduCharacter: "SchoolName",
		Educational:  "we",
		Speciality:   "dsa",
		StartTime:    "das",
		EndTime:      "da",
		SchoolEx:     "SchoolEx",
	}
	b := &database.EduEx{}
	b.Id = 1
	TranHttpStruct2Database(e, b)
	fmt.Println( b)
}