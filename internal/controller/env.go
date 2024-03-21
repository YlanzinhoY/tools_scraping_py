package controller

import "github.com/ylanzinhoy/tools_scraping_py/internal/services"



func CreateEnv() {
	services.CommandRunner("Env", "Python3 -m venv env")
}