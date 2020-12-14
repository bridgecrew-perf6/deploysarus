package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cjaewon/deploysarus/utils"
	"github.com/cjaewon/deploysarus/utils/color"
	"github.com/cjaewon/deploysarus/utils/config"
	"github.com/cjaewon/deploysarus/utils/logger.go"
	"gopkg.in/go-playground/webhooks.v5/github"
)

func githubHandler() func(w http.ResponseWriter, r *http.Request) {
	// check post when no secret key
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
				logger.Errorln(err)
			}

			return
		}

		switch payload.(type) {
		case github.PushPayload:
			var push github.PushPayload = payload.(github.PushPayload)
			var jobs Jobs

			trigger := config.GetString("on.push.trigger")
			branches := config.GetStringSlice("on.push.branches")

			if trigger != "" {
				logger.Warnln("Can't find trigger of push event")
				return
			}

			if len(branches) != 0 && !utils.Contain(branches, push.Ref) {
				return
			}

			if err := config.UnmarshalKey(fmt.Sprintf("jobs.%s", trigger), &jobs); err != nil {
				logger.ErrorlnfFatal("%s jobs parsing error, %v", trigger, err)
			}

			logger.Printlnf(color.Bold("Starting %s ..."), trigger)

			for _, step := range jobs.Steps {
				runStep(&step)
			}
		}
	}
}
