package storage

const MAX_STARTUP = 100
const MAX_USER = 100

var Startups [MAX_STARTUP]Startup
var JumlahStartup int
var Users [MAX_USER]User
var UserCount int
var CurrentUser User

func InitData() {
	InitAdmin()
}
