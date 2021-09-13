package db

type Params struct {
	Host     string
	Port     int
	User     string
	Dbname   string
	Password string
}

var CurrentParams Params
