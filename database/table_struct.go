package database

type User struct {
	Id         int64 `json:"id"`
	Account    string    `json:"account"`
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	ImageId    int64     `json:"image_id"`
	QqNumber 	string `json:"qq_number"`
	IsPersonnel string `json:"is_personnel"`
	TelNumber string `json:"tel_number"`
	Email string `json:"email"`
}
type EduEx struct {
	Id int64 `json:"id"`
	SchoolName string `json:"school_name"`
	EduCharacter string `json:"edu_character"`
	Educational string `json:"educational"`
	Speciality string `json:"speciality"`
	StartTime string  `json:"start_time"`
	EndTime string  `json:"end_time"`
	SchoolEx string `json:"school_ex"`
	ResumeId int64 `json:"resume_id"`
}

type PersonalInformation struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Sex string `json:"sex"`
	Status string `json:"status"`
	ResumeId int64 `json:"resume_id"`
}

type Recruitment struct {
	Id int64 `json:"id"`
	Trade string `json:"trade"`
	Position string `json:"position"`
	Status string `json:"status"`
	Location string `json:"location"`
	Pay string `json:"pay"`
	NeedSkills string `json:"need_skills"`
	AboutUs string `json:"about_us"`
}

type Resume struct {
	Id int64 `json:"id"`
	ResumeName string `json:"resume_name"`
	PersonalSkill string `json:"personal_skill"`
	PresonalAdvantage string `json:"presonal_information"`
	Certificate string `json:"certificate"`
	PersonalAddress string `json:"personal_address"`
	Recommend string `json:"recommend"`
}

type WorkEx struct {
	Id int64 `json:"id"`
	Company string `json:"company"`
	SubjectRole string `json:"subject_role"`
	SubjectLink string `json:"subject_link"`
	StartTime string`json:"start_time"`
	EndTime string`json:"end_time"`
	SubjectDescribe string `json:"subject_describe"`
	OwnAchievement string `json:"own_achievement"`
	ResumeId int64 `json:"resume_id"`
}