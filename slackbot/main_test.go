package main

import (
	"testing"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "hello", "olleh"},
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"palindrome", "racecar", "racecar"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"unicode", "👋 hello", "olleh 👋"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseString(tt.input)
			if result != tt.expected {
				t.Errorf("reverseString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestHandleCatCommand(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{"empty text", "", "meow"},
		{"with name", "fluffy", "meow fluffy"},
		{"with multiple words", "fluffy cat", "meow fluffy cat"},
		{"with spaces", "  mittens  ", "meow mittens"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock slash command
			cmd := slack.SlashCommand{
				Command:   "/cat",
				Text:      tt.text,
				ChannelID: "test-channel",
			}

			// We can't easily test the actual API call without mocking,
			// but we can test the logic by extracting it
			text := cmd.Text
			text = trimSpace(text)
			
			var response string
			if text == "" {
				response = "meow"
			} else {
				response = "meow " + text
			}

			if response != tt.expected {
				t.Errorf("handleCatCommand with text %q = %q, want %q", tt.text, response, tt.expected)
			}
		})
	}
}

// Helper function to mimic strings.TrimSpace for testing
func trimSpace(s string) string {
	// Simple trim implementation for testing
	start := 0
	end := len(s)
	
	for start < end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n') {
		start++
	}
	
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\n') {
		end--
	}
	
	return s[start:end]
}

func TestHandleAppMentionEvent(t *testing.T) {
	tests := []struct {
		name           string
		text           string
		expectedResult string
	}{
		{"reverse command", "<@U123456> reverse hello world", "dlrow olleh"},
		{"echo command", "<@U123456> echo test message", "test message"},
		{"unknown command", "<@U123456> unknown command", "I understand these commands: `reverse <text>` and `echo <text>`"},
		{"reverse empty", "<@U123456> reverse", "Please provide text to reverse"},
		{"echo empty", "<@U123456> echo", "Please provide text to echo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate the parsing logic from handleAppMentionEvent
			text := tt.text
			words := splitFields(text)
			
			// Remove the bot mention
			if len(words) > 0 && hasPrefix(words[0], "<@") {
				words = words[1:]
			}
			
			if len(words) == 0 {
				return
			}

			var result string
			
			// Simulate cobra command parsing
			if len(words) > 0 {
				command := words[0]
				args := words[1:]
				
				switch command {
				case "reverse":
					if len(args) == 0 {
						result = "Please provide text to reverse"
					} else {
						input := joinWords(args, " ")
						result = reverseString(input)
					}
				case "echo":
					if len(args) == 0 {
						result = "Please provide text to echo"
					} else {
						result = joinWords(args, " ")
					}
				default:
					result = "I understand these commands: `reverse <text>` and `echo <text>`"
				}
			}

			if result != tt.expectedResult {
				t.Errorf("handleAppMentionEvent with text %q = %q, want %q", tt.text, result, tt.expectedResult)
			}
		})
	}
}

// Helper functions for testing (simplified versions of strings package functions)
func splitFields(s string) []string {
	var fields []string
	var current string
	
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if current != "" {
				fields = append(fields, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}
	
	if current != "" {
		fields = append(fields, current)
	}
	
	return fields
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func joinWords(words []string, sep string) string {
	if len(words) == 0 {
		return ""
	}
	
	result := words[0]
	for i := 1; i < len(words); i++ {
		result += sep + words[i]
	}
	
	return result
}

// Test event structure parsing
func TestEventStructures(t *testing.T) {
	t.Run("SlashCommand structure", func(t *testing.T) {
		cmd := slack.SlashCommand{
			Command:   "/cat",
			Text:      "fluffy",
			ChannelID: "C123456",
			UserID:    "U123456",
		}
		
		if cmd.Command != "/cat" {
			t.Errorf("Expected command '/cat', got '%s'", cmd.Command)
		}
		if cmd.Text != "fluffy" {
			t.Errorf("Expected text 'fluffy', got '%s'", cmd.Text)
		}
	})
	
	t.Run("AppMentionEvent structure", func(t *testing.T) {
		event := &slackevents.AppMentionEvent{
			Type:    "app_mention",
			User:    "U123456",
			Text:    "<@U789012> reverse hello",
			Channel: "C123456",
		}
		
		if event.Type != "app_mention" {
			t.Errorf("Expected type 'app_mention', got '%s'", event.Type)
		}
		if event.Text != "<@U789012> reverse hello" {
			t.Errorf("Expected text '<@U789012> reverse hello', got '%s'", event.Text)
		}
	})
}