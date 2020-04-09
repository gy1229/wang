package resume_server

import (
	"fmt"
	"github.com/gy1229/wang/constant"
	"github.com/gy1229/wang/database"
	"github.com/gy1229/wang/json_struct"
	"github.com/gy1229/wang/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

func CreateResume(req *json_struct.CreateResumeRequest) (*json_struct.CreateResumeResponse, error) {
	resume := &database.Resume{}
	resumeId := util.GenId()
	resume.Id = resumeId
	for _, v := range req.WorkEx {
		workEx := &database.WorkEx{}
		workEx.Id = util.GenId()
		workEx.ResumeId = resumeId
		util.TranHttpStruct2Database(v, workEx)
		err := database.CreateWorkEx(workEx)
		if err != nil {
			logrus.Error("[CreateResume] CreateWorkEx err", err.Error())
			return nil,  err
		}
	}

	eduEx := &database.EduEx{}
	eduEx.Id = util.GenId()
	eduEx.ResumeId = resumeId
	util.TranHttpStruct2Database(req.EduEx, eduEx)
	err := database.CreateEduEx(eduEx)
	if err != nil {
		logrus.Error("[CreateResume] CreateEduEx err", err.Error())
		return nil,  err
	}

	pInfo := &database.PersonalInformation{}
	pInfo.Id = util.GenId()
	pInfo.ResumeId = resumeId
	util.TranHttpStruct2Database(req.PersonalInformation, pInfo)
	err = database.CreatePersonalInfo(pInfo)
	if err != nil {
		logrus.Error("[CreateResume] CreatePersonalInfo err", err.Error())
		return nil,  err
	}

	util.TranHttpStruct2Database(req, resume)
	err = database.CreateResume(resume)
	if err != nil {
		logrus.Error("[CreateResume] CreateResume err", err.Error())
		return nil,  err
	}
	return &json_struct.CreateResumeResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
		ResumeId: strconv.FormatInt(resumeId, 10),
	}, nil

}

func UploadResume(req *json_struct.UploadResumeRequest) (*json_struct.UploadResumeResponse, error) {
	resumeId, err := strconv.ParseInt(req.ResumeId, 10, 64)
	if err != nil {
		return nil,  err
	}
	workExs, err := database.FindWorkExByResumeId(&resumeId)
	if err != nil {
		return nil,  err
	}
	for _, w := range workExs {
		err = database.DeleteWorkEx(w)
		if err != nil {
			return nil,  err
		}
	}
	for _, w := range req.WorkEx {
		dw := &database.WorkEx{}
		dw.ResumeId = resumeId
		dw.Id = util.GenId()
		util.TranHttpStruct2Database(w, dw)
		err = database.CreateWorkEx(dw)
		if err != nil {
			return nil,  err
		}
	}
	eduExs, err := database.FindEduExByResumeId(&resumeId)
	if err != nil {
		return nil,  err
	}
	for _, w := range eduExs {
		err = database.DeletEduEx(w)
		if err != nil {
			return nil,  err
		}
	}
	de := &database.EduEx{}
	de.Id = util.GenId()
	de.ResumeId = resumeId
	util.TranHttpStruct2Database(req.EduEx, de)
	err = database.CreateEduEx(de)
	if err != nil {
		return nil,  err
	}
	pInfos, err := database.FindPersonalInfoByResumeId(&resumeId)
	if err != nil {
		return nil,  err
	}
	for _, w := range pInfos {
		err = database.DeletPersonalInfo(w)
		if err != nil {
			return nil,  err
		}
	}
	dp := &database.PersonalInformation{}
	dp.Id = util.GenId()
	dp.ResumeId = resumeId
	util.TranHttpStruct2Database(req.PersonalInformation, dp)
	err = database.CreatePersonalInfo(dp)
	if err != nil {
		return nil,  err
	}
	resume := &database.Resume{
		Id:                resumeId,
		ResumeName:        req.ResumeId,
		PersonalSkill:     req.PersonalSkill,
		PresonalAdvantage: req.PersonalAdvantage,
		Certificate:       req.Certificate,
		PersonalAddress:   req.PersonalAddress,
		Recommend:         req.Recommend,
	}
	err = database.UpdateResume(resume)
	if err != nil {
		return nil,  err
	}
	return &json_struct.UploadResumeResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func DeleteResume(req *json_struct.DeleteResumeRequest) (*json_struct.DeleteResumeResponse, error) {
	resumeId, err := strconv.ParseInt(req.ResumeId, 10, 64)
	if err != nil {
		return nil,  err
	}
	err = database.DeleteResumeById(&resumeId)
	if err != nil {
		return nil,  err
	}
	return &json_struct.DeleteResumeResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func CreateRecruitment(req *json_struct.CreateRecruitmentRequest) (*json_struct.CreateRecruitmentResponse, error) {
	recruit := &database.Recruitment{}
	recruit.Id = util.GenId()
	util.TranHttpStruct2Database(req, recruit)
	err := database.CreateRecruitment(recruit)
	if err != nil {
		return nil,  err
	}
	return &json_struct.CreateRecruitmentResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func UploadRecruitment(req *json_struct.UploadRecruitmentRequest) (*json_struct.UploadRecruitmentResponse, error) {
	recruitId, err := strconv.ParseInt(req.RecruitId, 10, 64)
	if err != nil {
		return nil,  err
	}
	recruit := &database.Recruitment{}
	recruit.Id = recruitId
	util.TranHttpStruct2Database(req, recruit)
	err = database.UpdateRecruitment(recruit)
	if err != nil {
		return nil,  err
	}
	return &json_struct.UploadRecruitmentResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func DeleteRecruitment(req *json_struct.DeleteRecruitmentRequest) (*json_struct.DeleteRecruitmentResponse, error) {
	recruitId, err := strconv.ParseInt(req.RecruitId, 10, 64)
	if err != nil {
		return nil,  err
	}
	err = database.DeleteRecruitById(&recruitId)
	if err != nil {
		return nil,  err
	}
	return &json_struct.DeleteRecruitmentResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func GetAvatar(req *json_struct.GetAvatarRequest) (*json_struct.GetAvatarResponse, error) {
	return &json_struct.GetAvatarResponse{
		ImageUrl: fmt.Sprintf("http://%s/img/%s.png",viper.GetString("FileServer.Addr"), req.ImageId),
	}, nil
}