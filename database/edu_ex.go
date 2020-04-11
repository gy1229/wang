package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func CreateEduEx(eduEx *EduEx) error {
	if err := DB.Create(eduEx).Error; err != nil {
		logrus.Error("CreateEduEx err ", err.Error())
		return err
	}
	return nil
}

func UpdateEduEx(eduEx *EduEx) error {
	if err := DB.Model(&eduEx).Where("id = ?", eduEx.Id).Updates(eduEx).Error; err != nil {
		logrus.Error("UpdateEduEx err ", err.Error())
		return err
	}
	return nil
}

func DeletEduEx(eduEx *EduEx) error {
	if err := DB.Delete(eduEx).Error; err != nil {
		logrus.Error("DeletEduEx err ", err.Error())
		return err
	}
	return nil
}

func FindEduEx(eduEx *EduEx) error {
	if err := DB.Where("id = ?", eduEx.Id).First(&eduEx).Error; err != nil {
		logrus.Error("FindEduEx err ", err.Error())
		return err
	}
	return nil
}

func FindEduExByResumeId(resumeId *int64) ([]*EduEx, error) {
	eduEx := make([]*EduEx, 0)
	if err := DB.Model(&eduEx).Where("resume_id = ?", resumeId).Find(&eduEx).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return eduEx, nil
		}
		logrus.Error("[FindEduExByResumeId] err msg", err.Error())
		return nil, err
	}
	return eduEx, nil
}