package server

import (
	"github.com/cjaewon/deploysarus/utils/config"
	"github.com/cjaewon/deploysarus/utils/logger.go"
	"gopkg.in/go-playground/webhooks.v5/github"
)

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
