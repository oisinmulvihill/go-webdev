package core

type Configuration struct {
	Interface   string
	Port        int
	GitCommit   string
	DatabaseDSN string
}

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt string
}

type Users struct {
	Users []User
}
