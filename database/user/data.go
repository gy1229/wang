package data_user

import (
	"github.com/gy1229/wang/database"
	"github.com/sirupsen/logrus"
)

func GetUserNameById(id int64) (string, error) {
	user := &database.User{
		Id: id,
	}
	if err := database.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		logrus.Error("GetUserNameById err ", err.Error())
		return "", err
	}
	return user.UserName, nil
}
