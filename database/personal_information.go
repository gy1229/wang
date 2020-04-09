package database

import (
	"github.com/gy1229/oa/database"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func CreatePersonalInfo(p *PersonalInformation) error {
	if err := database.DB.Create(p).Error; err != nil {
		logrus.Error("CreateEduEx err ", err.Error())
		return err
	}
	return nil
}

func UpdatePersonalInfo(p *PersonalInformation) error {
	if err := database.DB.Model(&p).Where("id = ?", p.Id).Updates(p).Error; err != nil {
		logrus.Error("UpdateUserMessage err ", err.Error())
		return err
	}
	return nil
}

func DeletPersonalInfo(p *PersonalInformation) error {
	if err := database.DB.Delete(p).Error; err != nil {
		logrus.Error("DeletEduEx err ", err.Error())
		return err
	}
	return nil
}

func FindPersonalInfo(p *PersonalInformation) error {
	if err := database.DB.Where("id = ?", p.Id).First(&p).Error; err != nil {
		logrus.Error("LoadUserMessage err ", err.Error())
		return err
	}
	return nil
}

func FindPersonalInfoByResumeId(resumeId *int64)  ([]*PersonalInformation, error) {
	p := make([]*PersonalInformation, 0)
	if err := database.DB.Model(&p).Where("resume_id = ?", resumeId).Find(&p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return p, nil
		}
		logrus.Error("[FindPersonalInfoByResumeId] err msg", err.Error())
		return nil, err
	}
	return p, nil
}
