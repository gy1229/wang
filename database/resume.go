package database

import (
	"github.com/sirupsen/logrus"
)

func CreateResume(resume *Resume) error {
	if err := DB.Create(resume).Error; err != nil {
		logrus.Error("CreateEduEx err ", err.Error())
		return err
	}
	return nil
}

func UpdateResume(resume *Resume) error {
	if err := DB.Model(&resume).Where("id = ?", resume.Id).Updates(resume).Error; err != nil {
		logrus.Error("UpdateUserMessage err ", err.Error())
		return err
	}
	return nil
}

func DeleteResumeById(id *int64) error {
	if err := DB.Where("id = ?", id).Delete(Resume{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteResume(resume *Resume) error {
	if err := DB.Delete(resume).Error; err != nil {
		logrus.Error("DeletEduEx err ", err.Error())
		return err
	}
	return nil
	return nil
}

func FindResume(resume *Resume) error {
	if err := DB.Where("id = ?", resume.Id).First(&resume).Error; err != nil {
		logrus.Error("LoadUserMessage err ", err.Error())
		return err
	}
	return nil
}