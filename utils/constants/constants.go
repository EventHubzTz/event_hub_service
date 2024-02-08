package constants

type UserRole string

const (
	AdminUser              UserRole = "admin"
	NormalUser             UserRole = "user"
	RootChatCollection     string   = "chats_list"
	ChatMessagesCollection string   = "messages"
	LocalStorage           string   = "LOCAL"
	RemoteStorage          string   = "REMOTE"
	UsersMessageToken      string   = "users_message_token"
	OnlineUserCollection   string   = "afya_app_online_user"
)

const (
	Yes = "YES"
	No  = "NO"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)
