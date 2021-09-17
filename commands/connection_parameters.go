package commands

var ConOpts struct {
	Host     string `short:"H" long:"host" help:"Database server host or socket directory." default:"localhost"`
	Port     int    `short:"p" long:"port" help:"Database server port." default:"5432"`
	Username string `short:"U" long:"username" help:"Database user name." default:"postgres"`
	Database string `short:"d" long:"dbname" help:"Database name to connect to." default:"postgres"`
}
