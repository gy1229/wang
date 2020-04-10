package json_struct

type UserBase struct {
	Account    string `json:"account"  binding:"required"`
	Password   string `json:"password"  binding:"required"`
	IsPersonal string `json:"is_personal"`
}

type RegisterRequest struct {
	UserBase *UserBase `json:"user_base"`
	UserName string    `json:"user_name"  binding:"required"`
	ImageId string `json:"image_id"`
	QQ string `json:"QQ"`
	Tel string `json:"Tel"`
	Email string `json:"Email"`

}

type RegisterResponse struct {
	Base *BaseResponse `json:"base"`
}

type LoginRequest struct {
	UserBase *UserBase `json:"user_base"  binding:"required"`
}

type LoginResponse struct {
	UserId string        `json:"user_id"`
	ImageId string `json:"image_id"`
	Base   *BaseResponse `json:"base"`
}

type UpdateUserMessageRequest struct {
	UserBase *UserBase `json:"user_base"`
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	QQ string `json:"QQ"`
	Tel string `json:"Tel"`
	Email string `json:"Email"`
}

type UpdateUserMessageResponse struct {
	Base *BaseResponse `json:"base"`
}

type LoadUserMessageRequest struct {
	UserId string `json:"user_id" binding`
}

type LoadUserMessageResponse struct {
	UserBase *UserBase `json:"user_base"`
	UserName string `json:"user_name"`
	ImageId string `json:"image_id"`
	QQ string `json:"QQ"`
	Tel string `json:"Tel"`
	Email string `json:"Email"`
	Base     *BaseResponse `json:"base"`
}

type CertainAccountRequest struct {
	Account string `json:"account"`
}

type CertainAccountResponse struct {
	Base *BaseResponse `json:"base"`
}

type PersonalInformation struct {
	Name string `json:"name"`
	Educational string `json:"educational"`
	Birthday string `json:"birthday"`
	Sex string `json:"sex"`
	Status string `json:"status"`
}

type WorkEx struct {
	SubiectName string `json:"subiect_name"`
	SubjectRole string `json:"subject_role"`
	SubjectLink string `json:"subject_link"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	Company string `json:"company"`
	SubejectDescrilbe string `json:"subeject_descrilbe"`
	OwnAchievement string `json:"own_achievement"`
}

type EduEx struct {
	SchoolName string `json:"school_name"`
	EduCharacter string `json:"edu_character"`
	Educational string `json:"educational"`
	Speciality string `json:"speciality"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	SchoolEx string `json:"school_ex"`
}

type CreateResumeRequest struct {
	UserId string `json:"user_id"`
	ResumeName string `json:"resume_name"`
	PersonalInformation *PersonalInformation `json:"personal_information"`
	PersonalSkill string `json:"personal_skill"`
	PersonalAdvantage string `json:"personal_advantage"`
	WorkEx []*WorkEx `json:"work_ex"`
	EduEx *EduEx `json:"edu_ex"`
	Certificate string `json:"certificate"`
	PersonalAddress string `json:"personal_address"`
	Recommend string `json:"recommend"`
}

type CreateResumeResponse struct {
	Base *BaseResponse `json:"base"`
	ResumeId string `json:"resume_id"`
}

type UploadResumeRequest struct {
	ResumeId string `json:"resume_id"`
	UserId string `json:"user_id"`
	ResumeName string `json:"resume_name"`
	PersonalInformation *PersonalInformation `json:"personal_information"`
	PersonalSkill string `json:"personal_skill"`
	PersonalAdvantage string `json:"personal_advantage"`
	WorkEx []*WorkEx `json:"work_ex"`
	EduEx *EduEx `json:"edu_ex"`
	Certificate string `json:"certificate"`
	PersonalAddress string `json:"personal_address"`
	Recommend string `json:"recommend"`
}

type UploadResumeResponse struct {
	Base *BaseResponse `json:"base"`
}

type DeleteResumeRequest struct {
	ResumeId string `json:"resume_id"`
	UserId string `json:"user_id"`
}

type DeleteResumeResponse struct {
	Base *BaseResponse `json:"base"`
}

type CreateRecruitmentRequest struct {
	UserId string `json:"user_id"`
	Trade string `json:"trade"`
	Position string `json:"position"`
	Status string `json:"status"`
	Location string `json:"location"`
	Pay string `json:"pay"`
	NeedSkills string `json:"need_skills"`
	AboutUs string `json:"about_us"`
}

type CreateRecruitmentResponse struct {
	Base *BaseResponse `json:"base"`
	RecruitId string `json:"recruit_id"`
}

type UploadRecruitmentRequest struct {
	UserId string `json:"user_id"`
	Trade string `json:"trade"`
	Position string `json:"position"`
	Status string `json:"status"`
	Location string `json:"location"`
	Pay string `json:"pay"`
	NeedSkills string `json:"need_skills"`
	AboutUs string `json:"about_us"`
	RecruitId string `json:"recruit_id"`
}

type UploadRecruitmentResponse struct {
	Base *BaseResponse `json:"base"`
}

type DeleteRecruitmentRequest struct {
	UserId string `json:"user_id"`
	RecruitId string `json:"recruit_id"`
}
type DeleteRecruitmentResponse struct {
	Base *BaseResponse `json:"base"`
}

type GetAvatarRequest struct {
	ImageId string `json:"image_id"`
}

type GetAvatarResponse struct {
	ImageUrl string `json:"image_url"`
}

type UploadAvatarRequest struct {

}

type UploadAvatarResponse struct {
	Base *BaseResponse `json:"base"`
	ImageId string `json:"image_id"`
}