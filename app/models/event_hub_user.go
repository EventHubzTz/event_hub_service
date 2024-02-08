package models

type EventHubUser struct {
	ID
	FirstName   string `json:"first_name" gorm:"not null;size:50" validate:"min=3"`
	LastName    string `json:"last_name" gorm:"not null;size:50" validate:"min=3"`
	PhoneNumber string `json:"phone_no" gorm:"not null;unique;size:20" validate:"required,min=9,unique=users_tests.phone_number"`
	Gender      string `json:"gender" gorm:"not null;type:enum('MALE','FEMALE');default:'MALE'"`
	Role        string `json:"role" gorm:"not null;type:enum('AGENT','ADIMINISTRATOR');default:'AGENT'"`
	Password    string `json:"password" gorm:"null" `
	Active      bool   `json:"active" gorm:"not null;default:true"`
	Timestamp
}

type KataTiketiUserDTO struct {
	ID
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_no"`
	Gender            string `json:"gender"`
	Role              string `json:"role"`
	AgentBusStandName string `json:"agent_bus_stand_name"`
	Active            bool   `json:"active"`
	TimestampString
}

func (EventHubUser) TableName() string {
	return tablePrefix + "users"
}
