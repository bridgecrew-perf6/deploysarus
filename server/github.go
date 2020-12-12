package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cjaewon/deploysarus/utils/commandline"
	"github.com/cjaewon/deploysarus/utils/config"
	"github.com/cjaewon/deploysarus/utils/logger.go"
	"gopkg.in/go-playground/webhooks.v5/github"
)

func githubHandler() func(w http.ResponseWriter, r *http.Request) {
	secretKey := config.GetString("secret_key")
	if secretKey == "" {
		logger.ErrorlnFatal("Cant't find secret_key from config file or env")
		os.Exit(1)
	}

	hook, err := github.New(github.Options.Secret(secretKey))
	if err != nil {
		logger.ErrorlnFatal(err)
	}

	events := parseGithubEvent("github")

	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, events...)
		if err != nil {
			if err != github.ErrEventNotFound {
				logger.Error(err)
			}
		}

		switch payload.(type) {
		case github.PushPayload:
			// push := payload.(github.PushPayload)
			// fmt.Printf("%+v", push)
			var steps []Step

			trigger := config.GetString("on.push.trigger")
			err := config.UnmarshalKey(fmt.Sprintf("jobs.%s.steps", trigger), &steps)

			if err != nil {
				logger.ErrorlnfFatal("%s Steps parsing error, %v", trigger, err)
			}

			for _, step := range steps {
				if step.Name != "" {
					logger.Printlnf("Start running %s step", step.Name)
				} else {
					logger.Println("Start running %s step", step.Run)
				}

				multiline := strings.Split(step.Run, "\n")
				for _, line := range multiline {
					name, args, err := commandline.ParseCommandline(line)
					if err != nil {
						logger.Errorln("Commandline parsing error at %s")
						continue
					}

					if err := commandline.Execute(name, args...); err != nil {
						logger.Error()
						continue
					}
				}
			}
		}
	}
}
