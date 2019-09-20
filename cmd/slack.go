package cmd

import (
	"context"
	"fmt"
	"path/filepath"

	"../integrations/slack"
	"github.com/spf13/cobra"
)

var slackCMD = &cobra.Command{
	Use:   "slack",
	Short: "Allows you to interact with Slack's API",
	Run: func(cmd *cobra.Command, args []string) {
		key := interface{}(slack.KeyToken)
		ctx := context.WithValue(context.Background(), key, token)

		var outfile string
		var output interface{}

		if len(getCMD) > 0 {
			for _, get := range getCMD {
				if get == "users" {
					users, err := slack.GetUsers(ctx)
					if err != nil {
						cmd.PrintErrln(err)
					}
					outfile = "users.json"
					output = users
				}
				// if strings.Contains(getCMD, "...") {}
			}
		} // else if ...

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
	slackCMD.Flags().StringVar(&token, "token", "", "Legacy token. Info: https://api.slack.com/custom-integrations/legacy-tokens")
	slackCMD.MarkFlagRequired("token")

	rootCMD.AddCommand(slackCMD)
}
