package robots

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WhereIsBot struct {
}

var userUrlMap = make(map[string]string)

func init() {
	for source := range Config.Sources {
		name := Config.Sources[source].Name
		url := Config.Sources[source].URL
		userUrlMap[name] = url
	}
}

func (w WhereIsBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	url, contains := userUrlMap[command.Text]
	if contains {
		userLocationMap := getAndCreateMap(url)
		location := lookupLocation(userLocationMap)
		return command.Text + " says: " + location
	} else {
		return "Sorry, no such user"
	}
}

func getAndCreateMap(url string) map[string]string {
	var locations = make(map[string]string)

	resp, err := http.Get(url)
	if err != nil {
		return locations
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	lines := strings.Split(string(body), "\n")
	for index := range lines {
		fields := strings.Split(lines[index], "\t")
		locations[fields[0]] = fields[1]
	}
	return locations
}

func lookupLocation(locations map[string]string) string {
	date := today()
	// location := locations[date]
	location, contains := locations[date]
	if contains {
		return location
	} else {
		return "Lost!"
	}
}

func today() string {
	local, _ := time.LoadLocation("America/Los_Angeles")
	return time.Now().In(local).Format("02/01/2006")
}
