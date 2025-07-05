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
	"github.com/spf13/cobra"
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
			case socketmode.EventTypeSlashCommand:
				handleSlashCommand(evt, client, api)
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
						handleMessageEvent(ev, api)
					case *slackevents.AppMentionEvent:
						handleAppMentionEvent(ev, api)
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

func handleSlashCommand(evt socketmode.Event, client *socketmode.Client, api *slack.Client) {
	cmd, ok := evt.Data.(slack.SlashCommand)
	if !ok {
		fmt.Printf("Ignored slash command: %+v\n", evt)
		return
	}

	client.Ack(*evt.Request)

	switch cmd.Command {
	case "/cat":
		handleCatCommand(cmd, api)
	default:
		fmt.Printf("Unknown slash command: %s\n", cmd.Command)
	}
}

func handleCatCommand(cmd slack.SlashCommand, api *slack.Client) {
	text := strings.TrimSpace(cmd.Text)
	
	var response string
	if text == "" {
		response = "meow"
	} else {
		response = fmt.Sprintf("meow %s", text)
	}

	_, _, err := api.PostMessage(cmd.ChannelID, slack.MsgOptionText(response, false))
	if err != nil {
		fmt.Printf("failed posting cat command response: %v\n", err)
	}
}

func handleMessageEvent(ev *slackevents.MessageEvent, api *slack.Client) {
	// Handle regular messages (keep existing hello functionality)
	if ev.User != "" && strings.Contains(ev.Text, "hello") {
		_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText("Hello! How can I help you?", false))
		if err != nil {
			fmt.Printf("failed posting message: %v\n", err)
		}
	}
}

func handleAppMentionEvent(ev *slackevents.AppMentionEvent, api *slack.Client) {
	// Parse the mention text using cobra
	text := strings.TrimSpace(ev.Text)
	
	// Remove the bot mention from the text
	words := strings.Fields(text)
	if len(words) > 0 && strings.HasPrefix(words[0], "<@") {
		words = words[1:] // Remove the mention
	}
	
	if len(words) == 0 {
		return
	}

	// Create a cobra command to parse the input
	var response string
	
	rootCmd := &cobra.Command{
		Use:   "catbot",
		Short: "CatBot commands",
		Run: func(cmd *cobra.Command, args []string) {
			response = "I understand these commands: `reverse <text>` and `echo <text>`"
		},
	}

	reverseCmd := &cobra.Command{
		Use:   "reverse [text...]",
		Short: "Reverse the given text",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				response = "Please provide text to reverse"
				return
			}
			input := strings.Join(args, " ")
			response = reverseString(input)
		},
	}

	echoCmd := &cobra.Command{
		Use:   "echo [text...]",
		Short: "Echo the given text",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				response = "Please provide text to echo"
				return
			}
			response = strings.Join(args, " ")
		},
	}

	rootCmd.AddCommand(reverseCmd, echoCmd)

	// Set the args and execute
	rootCmd.SetArgs(words)
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	
	err := rootCmd.Execute()
	if err != nil {
		response = "I understand these commands: `reverse <text>` and `echo <text>`"
	}

	if response != "" {
		_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText(response, false))
		if err != nil {
			fmt.Printf("failed posting mention response: %v\n", err)
		}
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}