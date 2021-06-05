package main

import (
	"github.com/evansmurithi/drone-github-issue/plugin"
	"github.com/urfave/cli/v2"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "api-key",
			Usage:       "api key to access github api",
			EnvVars:     []string{"PLUGIN_API_KEY", "GITHUB_RELEASE_API_KEY", "GITHUB_TOKEN"},
			Destination: &settings.APIKey,
		},
		&cli.StringFlag{
			Name:        "title",
			Usage:       "title of the github issue",
			EnvVars:     []string{"PLUGIN_TITLE", "GITHUB_ISSUE_TITLE"},
			Destination: &settings.Title,
		},
		&cli.StringFlag{
			Name:        "body",
			Usage:       "body of the github issue",
			EnvVars:     []string{"PLUGIN_BODY", "GITHUB_ISSUE_BODY"},
			Destination: &settings.Body,
		},
		&cli.StringSliceFlag{
			Name:        "assignees",
			Usage:       "list of github usernames to be assigned to the github issue",
			EnvVars:     []string{"PLUGIN_ASSIGNEES", "GITHUB_ISSUE_ASSIGNEES"},
			Destination: &settings.Assignees,
		},
		&cli.StringSliceFlag{
			Name:        "labels",
			Usage:       "list of labels to assign to the github issue",
			EnvVars:     []string{"PLUGIN_LABELS", "GITHUB_ISSUE_LABELS"},
			Destination: &settings.Labels,
		},
		&cli.StringFlag{
			Name:        "base-url",
			Usage:       "api url, needs to be changed for ghe",
			Value:       "https://api.github.com/",
			EnvVars:     []string{"PLUGIN_BASE_URL", "GITHUB_ISSUE_BASE_URL"},
			Destination: &settings.BaseURL,
		},
	}
}
