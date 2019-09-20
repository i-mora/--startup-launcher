package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var jiraCMD = &cobra.Command{
	Use:   "jira",
	Short: "Allows you to interact with Jira's API",
	Run: func(cmd *cobra.Command, args []string) {
		// key := interface{}(jira.KeyToken)
		// ctx := context.WithValue(context.Background(), key, token)

		var outfile string
		var output interface{}

		if len(getCMD) > 0 {

		} else if len(postCMD) > 0 {

		}

		if export {
			path := filepath.Join(outdir, outfile)
			generate(path, output)
		}
		fmt.Println("...done.")
	},
}

func init() {
	jiraCMD.Flags().StringArrayVar(&getCMD, "get", []string{}, "Please specify what you want (users, projects, roles etc...).")
	jiraCMD.Flags().StringArrayVar(&postCMD, "post", []string{}, "Please specify what you want to post (users, users_to_project, etc...).")

	jiraCMD.Flags().BoolVar(&export, "export", false, "Generate output files.")
	jiraCMD.Flags().StringVar(&outdir, "out-dir", "./assets", "Output directory")
	jiraCMD.Flags().StringVar(&token, "token", "", "Legacy token. Info: https://api.slack.com/custom-integrations/legacy-tokens")
	jiraCMD.MarkFlagRequired("token")

	rootCMD.AddCommand(jiraCMD)
}
