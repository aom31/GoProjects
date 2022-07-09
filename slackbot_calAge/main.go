package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvents){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Events)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3781298018516-3764303209335-BziEriDLVXSVp2K4mtIML5lW")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03P9FQ8JF3-3778877581058-596fcf1fe2d8dda3b47819d8806ffa818919af33c1b81ad989afcb53d11220b5")

	//create bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	//stop bot 
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}


}
