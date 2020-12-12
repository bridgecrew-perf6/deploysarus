package server

import (
	"fmt"
	"net/http"
	"os"

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

	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PushEvent)
		if err != nil {
			if err != github.ErrEventNotFound {
				logger.Error(err)
			}
		}

		switch payload.(type) {
		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			fmt.Printf("%+v", pullRequest)
		}
	}
}
