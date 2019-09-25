package cmd

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"../integrations/slack"
	"github.com/spf13/cobra"
)

var slackCMD = &cobra.Command{
	Use:   "slack",
	Short: "Allows you to interact with Slack's API",
	Run: func(cmd *cobra.Command, args []string) {
		a := slack.API{
			Token:  token,
			URL:    url,
			Ctx:    context.Background(),
			Client: http.Client{},
		}

		var outfile string
		var output interface{}

		if len(getCMD) > 0 {
			for _, get := range getCMD {
				switch get {
				case "users":
					users, err := a.GETusers()
					if err != nil {
						cmd.PrintErrln(err)
					}
					outfile = "slack_users.json"
					output = users
				}
			}
		}

		if export {
			path := filepath.Join(outdir, outfile)
			generate(path, output)
		}
		fmt.Println("...done.")
	},
}

func init() {
	slackCMD.Flags().StringArrayVar(&getCMD, "get", []string{}, "Please specify what you want (users, etc...).")

	slackCMD.Flags().BoolVar(&export, "export", false, "Generate output files.")
	slackCMD.Flags().StringVar(&outdir, "out-dir", "./assets", "Output directory")

	slackCMD.Flags().StringVar(&url, "url", "https://slack.com/api", "API url")
	slackCMD.Flags().StringVar(&token, "token", "", "Legacy token. Info: https://api.slack.com/custom-integrations/legacy-tokens")
	slackCMD.MarkFlagRequired("token")

	rootCMD.AddCommand(slackCMD)
}
