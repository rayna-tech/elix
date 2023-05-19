package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var (
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		ReportCaller:    true,
		TimeFormat:      time.Kitchen,
	})
	Box = box.New(box.Config{Px: 2, Py: 1, Type: "Single", Color: "Yellow", TitlePos: "Top", TitleColor: "White"})
	HiddenMessage = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#212529"))
)


func ServerPort() string {
	args := os.Args[1:]
	if len(args) >= 1 {
		return args[0]
	} else {
		return "8080"
	}
}

func SocketPort() string {
	args := os.Args[1:]
	if len(args) >= 2 {
		return args[1]
	} else {
		return "8081"
	}
}

func StartUpScreen() {
	Box.Println("Elix", fmt.Sprintf("A simple yet effective go lang based key,\nvalue store database which has socket and web api support.\n\n%s", HiddenMessage.Render("Elix is not recommended for production based apps that may hold sensitive data.\nIt runs of json files to store data between startups and potental crashes.")))
}

func Included(array []string, value string) bool {
	found := false
	for _, element := range array {
		if element == value {
			found = true
			break
		}
	}
	return found
}