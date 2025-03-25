package core

type Configuration struct {
	Interface   string
	Port        int
	GitCommit   string
	DatabaseDSN string
	TemplateDir string
}

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt string
}
