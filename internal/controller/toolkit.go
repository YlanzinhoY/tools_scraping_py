package controller

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/ylanzinhoy/tools_scraping_py/internal/services"
)

func ToolkitController() {
	var choices = []string{"beautifulsoup", "requests", "schedule", "selenium", "openpyxl"}

	prompt := &survey.MultiSelect{
		Message: "Choose a lib",
		Options: choices,
	}

	var choiceType []string

	survey.AskOne(prompt, &choiceType)

	for _, choice := range choiceType {
		switch choice {
		case "beautifulsoup":
			services.CommandRunner("beautiful soup", "pip3 install beautifulsoup4")

		case "requests":
			services.CommandRunner("requests", "pip3 install requests")

		case "schedule":
			services.CommandRunner("requests", "pip3 install schedule")

		case "selenium":
			services.CommandRunner("selenium", "pip3 install selenium")
		case "openpyxl":
			services.CommandRunner("openpyxl", "pip3 install openpyxl")

		}
	}
}
