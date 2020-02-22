package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

type config struct {
	me    string
	token string
	debug bool
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cnf := config{
		me:    os.Getenv("ME"),
		token: os.Getenv("TOKEN"),
		debug: true,
	}

	api := slack.New(cnf.token, slack.OptionDebug(cnf.debug))
	getUsers(api)

	// spew.Dump(api.GetUserInfo("UH8PLH527"))

	// groups, err := api.GetGroups(false)
	// checkError(err)
	// for _, group := range groups {
	// 	fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	// }

	// channels, err := api.GetChannels(true)
	// checkError(err)
	// for _, channel := range channels {
	// 	spew.Dump(channel.ID + " - " + channel.Name)
	// }

}

func getUsers(api *slack.Client) {
	users, _ := api.GetUsers()
	for _, user := range users {
		spew.Dump(user.RealName + " " + user.ID)
	}
}

func getUserInfo(api *slack.Client, user string) {
	params := slack.GetConversationHistoryParameters{
		Limit:     10,
		ChannelID: "CJPB48NFR",
	}
	c, _ := api.GetConversationHistory(&params)
	spew.Dump(c)
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
