package robots

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WhereIsBot struct {
}

func init() {
}

func (w WhereIsBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	// command.Command
	latestMap := getAndCreateMap()
	location := lookupLocation(latestMap)
	return location
}

func getAndCreateMap() map[string]string {
	var locations = make(map[string]string)

	resp, err := http.Get(Config.Source)
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
	return time.Now().Format("02/01/2006")
}
