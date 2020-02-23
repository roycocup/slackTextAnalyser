package main

import (
	"errors"
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
	// steveLockwood := "U1XQT70UF"
	// davidgarcia := "UM10263HU"
	// softwareEngineeringChan := "C2957KZML"

	api := slack.New(cnf.token, slack.OptionDebug(cnf.debug))
	// getChannels(api)
	// spew.Dump(api.GetChannelInfo("C2957KZML"))
	// spew.Dump(api.GetUserInfo("UM10263HU"))

	targetChannel := "dev-support"

	deleteFile(targetChannel + ".txt")

	chanelID, err := getChannelByName(api, targetChannel)
	checkError(err)

	params := slack.GetConversationHistoryParameters{
		ChannelID: chanelID,
		Limit:     100,
	}
	convos, _ := api.GetConversationHistory(&params)
	numMessages := len(convos.Messages)

	storedLines := []string{}
	for i := numMessages - 1; i >= 0; i-- {
		msg := convos.Messages[i].Msg.Text
		spew.Dump(msg)
		write(targetChannel+".txt", msg)
		storedLines = append(storedLines, msg)
	}

	spew.Dump(len(storedLines))

	// spew.Dump(storedLines)
	// spew.Dump(len(storedLines))

	// c := cache.New(5*time.Minute, 10*time.Minute)
	// if (c.Get("users") != null){
	// }
	// c.Set("users", getUsers(api), cache.DefaultExpiration)

}

func deleteFile(fileName string) {
	err := os.Remove(fileName)
	checkError(err)
}

func getChannelByName(api *slack.Client, chanName string) (string, error) {

	channels := getChannels(api)

	for i := range channels {
		if channels[i].Name == chanName {
			return channels[i].ID, nil
		}
	}
	return "", errors.New("Channel not found")
}

func write(filename string, msg string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	checkError(err)
	defer f.Close()
	f.WriteString("\n" + msg)
}

func getChannels(api *slack.Client) []slack.Channel {
	channels, err := api.GetChannels(true)
	checkError(err)

	return channels
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
