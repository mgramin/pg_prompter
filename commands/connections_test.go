package commands

import (
	"encoding/xml"
	"github.com/xo/dburl"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type Con struct {
	Name string `xml:"name,attr"`
	Url  string `xml:"jdbc-url"`
	User string `xml:"user-name"`
}

func TestHelloName(t *testing.T) {
	content, _ := ioutil.ReadFile(os.Getenv("HOME") + "/.pg_prompter/idea")
	for _, line := range strings.Split(string(string(content)), "\n") {
		if strings.HasPrefix(line, "#") == false && line != "" {
			bytes := []byte(line)
			var res Con
			err := xml.Unmarshal(bytes, &res)
			if err != nil {
				panic(err)
			}
			u, _ := dburl.Parse(strings.TrimPrefix(res.Url, "jdbc:"))
			println(res.Name)
			println(u.Hostname())
			//println(strings.TrimPrefix(u.Path, "/"))
		}
	}

}
