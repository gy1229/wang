package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gy1229/wang/json_struct"
	"github.com/gy1229/wang/service/resume_server"
	"github.com/gy1229/wang/service/user_server"
	"github.com/gy1229/wang/util"
	"net/http"
)

func TestHanlder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Register(c *gin.Context) {
	var req json_struct.RegisterRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user_server.InsertUserMessage(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func Login(c *gin.Context) {
	var req json_struct.LoginRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user_server.LoginUser(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UpdateUserMessage(c *gin.Context) {
	var req json_struct.UpdateUserMessageRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user_server.UpdateUserMessage(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func LoadUserMessage(c *gin.Context) {
	var req json_struct.LoadUserMessageRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user_server.LoadUserMessage(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func CertainAccount(c *gin.Context) {
	var req json_struct.CertainAccountRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user_server.CertainAccount(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func CreateResume(c *gin.Context) {
	var req json_struct.CreateResumeRequest
	util.GenHandlerRequest(c, &req)
	resp, err := resume_server.CreateResume(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UploadResume(c *gin.Context) {
	var req json_struct.UploadResumeRequest
	util.GenHandlerRequest(c, &req)
	resp, err := resume_server.UploadResume(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func DeleteResume(c *gin.Context) {
	var req json_struct.DeleteResumeRequest
	util.GenHandlerRequest(c, &req)
	resp, err := resume_server.DeleteResume(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func CreateRecruitment(c *gin.Context) {
	var req json_struct.CreateRecruitmentRequest
	util.GenHandlerRequest(c, &req)
	resp, err := resume_server.CreateRecruitment(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UploadRecruitment(c *gin.Context) {
	var req json_struct.UploadRecruitmentRequest
	util.GenHandlerRequest(c, &req)
	resp, err := resume_server.UploadRecruitment(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func DeleteRecruitment(c *gin.Context) {
	var req json_struct.DeleteRecruitmentRequest
	util.GenHandlerRequest(c, &req)
	resp, err := resume_server.DeleteRecruitment(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func GetAvatar(c *gin.Context) {
	var req json_struct.GetAvatarRequest
	util.GenHandlerRequest(c, &req)
	resp, err := resume_server.GetAvatar(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}