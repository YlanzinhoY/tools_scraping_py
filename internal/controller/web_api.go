package controller

import (
	"github.com/ylanzinhoy/tools_scraping_py/internal/services"
	"github.com/ylanzinhoy/tools_scraping_py/internal/util"
)

func WebApi() {
	api := util.Choices([]string{"flask", "fastapi"}, "webkit")

	for _, choice := range api {
		switch choice {
		case "flask":
			services.CommandRunner("flask", "pip3 install flask")
		case "fastapi":
			services.CommandRunner("fastapi", "pip3 install fastapi")
		}
	}
}
