package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

type config struct {
	me    string
	token string
	debug bool
}

var users = make(map[string]*slack.User)

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

	targetChannel := "dev_announcements"

	deleteFile(targetChannel + ".txt")

	chanelID, err := getChannelByName(api, targetChannel)
	checkError(err)

	params := slack.GetConversationHistoryParameters{
		ChannelID: chanelID,
		Limit:     500,
	}
	convos, _ := api.GetConversationHistory(&params)
	numMessages := len(convos.Messages)

	storedLines := []string{}
	for i := numMessages - 1; i >= 0; i-- {
		msg := convos.Messages[i].Msg
		storable := msg.Timestamp + " - " + getUser(api, msg.User) + " - " + msg.Text
		write(targetChannel+".txt", storable)
		storedLines = append(storedLines, storable)
	}

	spew.Dump(strconv.Itoa(len(storedLines)) + " Lines")

}

func getUser(api *slack.Client, userCode string) string {
	if users[userCode] == nil {
		user, _ := api.GetUserInfo(userCode)
		users[userCode] = user
	}

	return users[userCode].RealName
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
