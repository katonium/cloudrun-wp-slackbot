package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	token := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	if token == "" {
		log.Fatal("SLACK_BOT_TOKEN environment variable is required")
	}
	if appToken == "" {
		log.Fatal("SLACK_APP_TOKEN environment variable is required")
	}

	api := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))

	client := socketmode.New(api, socketmode.OptionDebug(true))

	go func() {
		for evt := range client.Events {
			switch evt.Type {
			case socketmode.EventTypeConnecting:
				fmt.Println("Connecting to Slack with Socket Mode...")
			case socketmode.EventTypeConnectionError:
				fmt.Println("Connection failed. Retrying later...")
			case socketmode.EventTypeConnected:
				fmt.Println("Connected to Slack with Socket Mode.")
			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					fmt.Printf("Ignored %+v\n", evt)
					continue
				}

				client.Ack(*evt.Request)

				switch eventsAPIEvent.Type {
				case slackevents.CallbackEvent:
					innerEvent := eventsAPIEvent.InnerEvent
					switch ev := innerEvent.Data.(type) {
					case *slackevents.MessageEvent:
						if ev.User != "" && strings.Contains(ev.Text, "hello") {
							_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText("Hello! How can I help you?", false))
							if err != nil {
								fmt.Printf("failed posting message: %v", err)
							}
						}
					}
				default:
					client.Debugf("unsupported Events API event received")
				}
			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type received: %s\n", evt.Type)
			}
		}
	}()

	ctx := context.Background()
	err := client.RunContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
}