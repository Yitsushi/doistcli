package main

import (
	"github.com/yitsushi/doistcli/api"
	//	"log"
)

func main() {
	conf, err := parseConfiguration()

	if err != nil {
		panic(err)
	}

	api.SetToken(conf.APIToken)

	project := &api.Project{
		Name: "Test Project With API",
	}
	newProject := api.ProjectAdd{
		Project: project,
	}

	action := newProject.Prepare()

	action.Send()

	//log.Println(conf, action)
}
