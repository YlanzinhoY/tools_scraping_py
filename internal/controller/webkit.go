package controller

import (
	"github.com/ylanzinhoy/tools_scraping_py/internal/services"
	"github.com/ylanzinhoy/tools_scraping_py/internal/util"
)

func WebApi() {
	api := util.Choices([]string{"flask", "fastapi", "SQLAlchemy", "django", "django_rest_framework"}, "webkit")

	for _, choice := range api {
		switch choice {
		case "flask":
			services.CommandRunner("Flask", "pip3 install flask")
		case "fastapi":
			services.CommandRunner("Fastapi", "pip3 install fastapi")
		case "SQLAlchemy":
			services.CommandRunner("SQLAlchemy", "pip3 install SQLAlchemy")
		case "django":
			services.CommandRunner("django", "pip install django")
		case "django_rest_framework":
			services.CommandRunner("django_rest_framework", "pip install django djangorestframework")
		case "postgres_driver":
			services.CommandRunner("postgres", "pip install psycopg-binary")
		}
	}
}
