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
			services.CommandRunner("beautiful soup", "pip install beautifulsoup4")

		case "requests":
			services.CommandRunner("requests", "pip install requests")

		case "schedule":
			services.CommandRunner("requests", "pip install schedule")
		case "selenium":
			services.CommandRunner("selenium", "pip install selenium")
		case "openpyxl":
			services.CommandRunner("openpyxl", "pip install openpyxl")
		}
	}
}
