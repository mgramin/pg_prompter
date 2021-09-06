package main

type DbParams struct {
	host string
	port int
	user string
	dbname string
	password string
}

var dbParams DbParams