package util

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Choices(args []string, whatYouDo string) []string {

	var choices = args

	var choiceType []string

	prompt := &survey.MultiSelect{
		Message: fmt.Sprintf("Choose an %s", whatYouDo),
		Options: choices,
	}

	survey.AskOne(prompt, &choiceType)

	return choiceType

}
