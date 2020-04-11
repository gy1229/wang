package database

import (
	"github.com/sirupsen/logrus"
)

func CreateRecruitment(r *Recruitment) error {
	if err := DB.Create(r).Error; err != nil {
		logrus.Error("CreateEduEx err ", err.Error())
		return err
	}
	return nil
}

func UpdateRecruitment(r *Recruitment) error {
	if err := DB.Model(&r).Where("id = ?", r.Id).Updates(r).Error; err != nil {
		logrus.Error("UpdateUserMessage err ", err.Error())
		return err
	}
	return nil
}

func DeleteRecruitment(r *Recruitment) error {
	if err := DB.Delete(r).Error; err != nil {
		logrus.Error("DeletEduEx err ", err.Error())
		return err
	}
	return nil
}

func DeleteRecruitById(id *int64) error {
	if err := DB.Where("id = ?", id).Delete(Recruitment{}).Error; err != nil {
		return err
	}
	return nil
}

func FindRecruitment(r *Recruitment) error {
	if err := DB.Where("id = ?", r.Id).First(&r).Error; err != nil {
		logrus.Error("LoadUserMessage err ", err.Error())
		return err
	}
	return nil
}
