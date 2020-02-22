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
	// getChannels(api)
	spew.Dump(api.GetChannelInfo("CEXG0RH7W"))
	// spew.Dump(api.GetUserInfo("U1XQT70UF"))

	// c := cache.New(5*time.Minute, 10*time.Minute)
	// if (c.Get("users") != null){
	// }
	// c.Set("users", getUsers(api), cache.DefaultExpiration)

}

func getChannels(api *slack.Client) {
	channels, err := api.GetChannels(false)
	checkError(err)

	for _, channel := range channels {
		spew.Dump(channel.ID + " - " + channel.Name)
	}
	spew.Dump(len(channels))
}

func getGroups(api *slack.Client) {
	groups, err := api.GetGroups(true)
	checkError(err)

	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}
}

func getUsers(api *slack.Client) (user []slack.User) {
	users, _ := api.GetUsers()
	return users
}

func foreach(objs []interface{}) {
	for _, obj := range objs {
		spew.Dump(obj)
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
