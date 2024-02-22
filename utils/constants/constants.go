package constants

type UserRole string

const (
	Yes           string = "YES"
	No            string = "NO"
	LocalStorage  string = "LOCAL"
	RemoteStorage string = "REMOTE"
	Completed     string = "COMPLETED"
	Success       string = "success"
	Failure       string = "failure"
	Currency      string = "TZS"
	NormalUser    string = "NORMAL_USER"
	EventPlanner  string = "EVENT_PLANNER"
	SuperAdmin    string = "SUPER_ADMIN"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\u001b[33m"
	ColorBlue   Color = "\u001b[34m"
	ColorReset  Color = "\u001b[0m"
)
