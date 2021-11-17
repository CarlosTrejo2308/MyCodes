package main

import (
	"fmt"
	"strings"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	message := strings.Split(line, ":")[1]

	return strings.TrimSpace(message)
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	message := Message(line)
	return len([]rune(message))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	message := strings.Split(line, ":")[0]
	message = message[1 : len(message)-1]

	return strings.ToLower(message)
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", Message(line), LogLevel(line))
}
