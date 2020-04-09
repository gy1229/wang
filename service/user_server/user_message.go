package user_server

import (
	"fmt"
	"github.com/gy1229/wang/constant"
	"github.com/gy1229/wang/database"
	"github.com/gy1229/wang/json_struct"
	"github.com/gy1229/wang/util"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strconv"
)

func LoadUserMessage(req *json_struct.LoadUserMessageRequest) (*json_struct.LoadUserMessageResponse, error) {
	user := database.User{}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("LoadUserMessage err ", err.Error())
		return nil, err
	}
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		logrus.Error("LoadUserMessage err ", err.Error())
		return nil, err
	}
	return &json_struct.LoadUserMessageResponse{
		UserBase: &json_struct.UserBase{
			Account:  user.Account,
			Password: user.Password,
		},
		UserName: user.UserName,
		ImageId: fmt.Sprintf("%d", user.ImageId),
		QQ: user.QqNumber,
		Tel: user.TelNumber,
		Email: user.Email,
		Base:     &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func UpdateUserMessage(req *json_struct.UpdateUserMessageRequest) (*json_struct.UpdateUserMessageResponse, error) {
	user := database.User{
		Password: req.UserBase.Password,
		UserName: req.UserName,
		IsPersonnel: req.UserBase.AvatarId,
		QqNumber:   req.QQ,
		TelNumber:  req.Tel,
		Email: req.Email,
	}
	if err := database.DB.Model(&user).Where("account = ?", req.UserBase.Account).Updates(user).Error; err != nil {
		logrus.Error("UpdateUserMessage err ", err.Error())
		return nil, err
	}
	return &json_struct.UpdateUserMessageResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func InsertUserMessage(req *json_struct.RegisterRequest) (*json_struct.RegisterResponse, error) {
	imageId, _ := strconv.ParseInt(req.ImageId, 10, 64)
	user := database.User{
		Id:         util.GenId(),
		Account:    req.UserBase.Account,
		UserName:   req.UserName,
		Password:   req.UserBase.Password,
		IsPersonnel: req.UserBase.AvatarId,
		ImageId:    imageId,
		QqNumber:   req.QQ,
		TelNumber:  req.Tel,
		Email: req.Email,
	}
	if err := database.DB.Create(user).Error; err != nil {
		logrus.Error("InsertUserMessage err ", err.Error())
		return nil, err
	}
	return &json_struct.RegisterResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func LoginUser(req *json_struct.LoginRequest) (*json_struct.LoginResponse, error) {
	user := database.User{
		Account: req.UserBase.Account,
	}
	if err := database.DB.Model(&user).Where("account = ?", user.Account).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &json_struct.LoginResponse{
				Base: &json_struct.BaseResponse{Body: constant.LoginFailAccount},
			}, nil
		}
		logrus.Error("UpdateUserMessage err ", err.Error())
		return nil, err
	}
	if user.Password == req.UserBase.Password {
		return &json_struct.LoginResponse{
			UserId: strconv.FormatInt(user.Id, 10),
			ImageId: strconv.FormatInt(user.ImageId, 10),
			Base:   &json_struct.BaseResponse{Body: constant.SUCCESS},
		}, nil
	}
	return &json_struct.LoginResponse{
		Base: &json_struct.BaseResponse{Body: constant.LoginFailPassword},
	}, nil
}

func CertainAccount(req *json_struct.CertainAccountRequest) (*json_struct.CertainAccountResponse, error) {
	user := database.User{
		Account: req.Account,
	}
	if err := database.DB.Model(&user).Where("account = ?", user.Account).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &json_struct.CertainAccountResponse{
				Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
			}, nil
		}
		logrus.Error("UpdateUserMessage err ", err.Error())
		return nil, err
	}
	return &json_struct.CertainAccountResponse{
		Base: &json_struct.BaseResponse{Body: constant.RegisterAccountExit},
	}, nil
}
