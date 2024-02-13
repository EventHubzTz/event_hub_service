package models

type EventHubUser struct {
	ID
	FirstName     string `json:"first_name" gorm:"not null;size:50" validate:"min=3"`
	LastName      string `json:"last_name" gorm:"not null;size:50" validate:"min=3"`
	Email         string `json:"email" gorm:"not null;unique;size:50" validate:"required,email,min=3,unique=users_tests.email"`
	PhoneNumber   string `json:"phone_no" gorm:"not null;unique;size:20" validate:"required,min=9,unique=users_tests.phone_number"`
	Gender        string `json:"gender" gorm:"not null;type:enum('MALE','FEMALE');default:'MALE'"`
	ProfileImage  string `json:"profile_image" gorm:"null;size:200;default:/users/profileImages/male.png"`
	Role          string `json:"role" gorm:"not null;type:enum('NORMAL_USER','EVENT_PLANNER','SUPER_ADMIN');default:'NORMAL_USER'"`
	Password      string `json:"password" gorm:"not null"`
	ImageStorage  string `json:"image_storage" gorm:"not null;type:enum('LOCAL','REMOTE');default:'LOCAL'"`
	IsValidEmail  bool   `json:"is_valid_email" gorm:"not null;default:false"`
	IsValidNumber bool   `json:"is_valid_number" gorm:"not null;default:false"`
	Active        bool   `json:"active" gorm:"not null;default:true"`
	Timestamp
}

type EventHubUserDTO struct {
	ID
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_no"`
	Gender       string `json:"gender"`
	ProfileImage string `json:"profile_image"`
	Role         string `json:"role"`
	Active       bool   `json:"active"`
	TimestampString
}

func (EventHubUser) TableName() string {
	return tablePrefix + "users"
}
