package server

import (
	"strings"

	"github.com/cjaewon/deploysarus/utils/color"
	"github.com/cjaewon/deploysarus/utils/commandline"
	"github.com/cjaewon/deploysarus/utils/config"
	"github.com/cjaewon/deploysarus/utils/logger.go"

	"gopkg.in/go-playground/webhooks.v5/github"
)

// Jobs defines a Jobs type
type Jobs struct {
	Sync  bool   `mapstructure:"sync"`
	Steps []Step `mapstructure:"steps"`
}

// Step defines a jobs steps type
type Step struct {
	Name string                 `mapstructure:"name"`
	Run  string                 `mapstructure:"run"`
	Uses string                 `mapstructure:"uses"`
	With map[string]interface{} `mapstructure:"with"`
}

func parseGithubEvent(platform string) []github.Event {
	on := config.GetStringMap("on")
	events := []github.Event{}

	for event := range on {
		switch event {
		case "check_run":
			events = append(events, github.CheckRunEvent)
		case "check_suit":
			events = append(events, github.CheckSuiteEvent)
		case "create":
			events = append(events, github.CreateEvent)
		case "delete":
			events = append(events, github.DeleteEvent)
		case "deployment":
			events = append(events, github.DeploymentEvent)
		case "deployment_status":
			events = append(events, github.DeploymentStatusEvent)
		case "fork":
			events = append(events, github.ForkEvent)
		case "gollum":
			events = append(events, github.GollumEvent)
		case "issue_comment":
			events = append(events, github.IssueCommentEvent)
		case "issues":
			events = append(events, github.IssuesEvent)
		case "label":
			events = append(events, github.LabelEvent)
		case "milestone":
			events = append(events, github.MilestoneEvent)
		case "page_build":
			events = append(events, github.PageBuildEvent)
		case "project":
			events = append(events, github.ProjectEvent)
		case "project_card":
			events = append(events, github.ProjectCardEvent)
		case "project_column":
			events = append(events, github.ProjectColumnEvent)
		case "public":
			events = append(events, github.PublicEvent)
		case "pull_request":
			events = append(events, github.PullRequestEvent)
		case "pull_request_review":
			events = append(events, github.PullRequestEvent)
		case "pull_request_review_comment":
			events = append(events, github.PullRequestReviewCommentEvent)
		// case "pull_request_target":
		// events = append(events, github.PullRequestTargetEvent) Not Found
		case "push":
			events = append(events, github.PushEvent)
		// case "registry_package":
		// events = append(events, github.RegistryPackage) Not Found
		case "release":
			events = append(events, github.ReleaseEvent)
		case "status":
			events = append(events, github.StatusEvent)
		case "watch":
			events = append(events, github.WatchEvent)
		// case "workflow_run":
		// events = append(events, github.WorkflowRun) Not Found
		default:
			logger.Warn("Event name did't find at webhook event list")
		}
	}

	return events
}

func runStep(step *Step) {
	if step.Name != "" {
		logger.Printlnf(color.Bold("▶ %s:"), step.Name)
	} else {
		logger.Println(color.Bold("▶ %s:"), step.Run)
	}

	multiline := strings.Split(step.Run, "\n")
	for _, line := range multiline {
		name, args, err := commandline.ParseCommandline(line)
		if err != nil {
			logger.Errorln("Commandline parsing error at %s")
			continue
		}

		logger.Printlnf("$ %s", line)
		if err := commandline.Execute(name, args...); err != nil {
			logger.Error(err)
			continue
		}
	}

	logger.Println()
}
