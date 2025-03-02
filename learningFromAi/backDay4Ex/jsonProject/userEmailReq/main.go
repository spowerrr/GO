package main

import (
	"fmt"
	"os"
)

func main() {
	users := map[string]string{
		"Alice": "alice@example.com",
		"Rahul": "rahul@example.com",
	}

	// slice to store log entries
	var logs []string

	for name, email := range users {
		logentry := fmt.Sprintf("User %s with email %s requested data", name, email) // fmt.Sprintf: This formats a string, combining the name and email into a log entry (e.g., "User Alice with email alice@example.com requested data").
		// logs = append(logs, logEntry): This adds each log entry to the logs
		// slice. The append function adds the formatted log entry to the end of the slice.
		logs = append(logs, logentry)
	}

	// Open the file in append mode (create the file if it doesn't exist)
	logFile, err := os.OpenFile("user_requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error oppening file:", err)
		return
	}
	defer logFile.Close()

	// write the logs to file
	for _, log := range logs {
		_, err := logFile.WriteString(log + "\n") // append each log entry to the file
		if err != nil {
			fmt.Println("Error writing to file", err)
			return
		}

	}
	fmt.Println("All user logs saved Successfully")
}
