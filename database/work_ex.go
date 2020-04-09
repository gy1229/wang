package database

import (
	"github.com/gy1229/oa/database"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func CreateWorkEx(w *WorkEx) error {
	if err := database.DB.Create(w).Error; err != nil {
		logrus.Error("CreateEduEx err ", err.Error())
		return err
	}
	return nil
}

func UpdateWorkEx(w *WorkEx) error {
	if err := database.DB.Model(&w).Where("id = ?", w.Id).Updates(w).Error; err != nil {
		logrus.Error("UpdateUserMessage err ", err.Error())
		return err
	}
	return nil
}

func DeleteWorkEx(resume *WorkEx) error {
	if err := database.DB.Delete(resume).Error; err != nil {
		logrus.Error("DeletEduEx err ", err.Error())
		return err
	}
	return nil
}

func FindWorkEx(w *WorkEx)  error {
	if err := database.DB.Where("id = ?", w.Id).First(&w).Error; err != nil {
		logrus.Error("LoadUserMessage err ", err.Error())
		return err
	}
	return nil
}

func FindWorkExByResumeId(resumeId *int64) ([]*WorkEx, error) {
	w := make([]*WorkEx, 0)
	if err := database.DB.Model(&w).Where("resume_id = ?", resumeId).Find(&w).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return w, nil
		}
		logrus.Error("[FindFindWorkExByResumeId] err msg", err.Error())
		return nil, err
	}
	return w, nil
}
