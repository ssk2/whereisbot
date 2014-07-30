package robots

import (
	"fmt"
)

type WhereIsBot struct {
}

func init() {
}

func (w WhereIsBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	go w.DeferredAction(command)
	return "Hello Sunil"
}

func (w WhereIsBot) DeferredAction(command *SlashCommand) {
	response := new(IncomingWebhook)
	response.Channel = command.Channel_ID
	response.Username = "Where Is Bot"
	response.Text = fmt.Sprintf("@%s Hi!", command.User_Name)

	// TODO: Set icon
	response.Icon_Emoji = ":ghost:"
	response.Unfurl_Links = true
	response.Parse = "full"
	MakeIncomingWebhookCall(response)
}
