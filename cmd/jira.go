package cmd

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"../integrations/jira"
	"github.com/spf13/cobra"
)

var jiraCMD = &cobra.Command{
	Use:   "jira",
	Short: "Allows you to interact with Jira's API",
	Run: func(cmd *cobra.Command, args []string) {
		a := jira.API{
			Token:  token,
			User:   user,
			URL:    url,
			Ctx:    context.Background(),
			Client: http.Client{},
		}

		var outfile string
		var output interface{}

		if len(getCMD) > 0 {
			for _, get := range getCMD {
				switch get {
				case "projects":
					projects, err := a.GETprojects()
					if err != nil {
						cmd.PrintErrln(err)
					}
					outfile = "jira_projects.json"
					output = projects
				case "roles":
					roles, err := a.GETroles()
					if err != nil {
						cmd.PrintErrln(err)
					}
					outfile = "jira_roles.json"
					output = roles
				case "users":
					users, err := a.GETusers()
					if err != nil {
						cmd.PrintErrln(err)
					}
					outfile = "jira_users.json"
					output = users
				}
			}
		} else if len(postCMD) > 0 {
			for _, post := range postCMD {
				switch post {
				case "issue":
				case "project": // join user to a project
				case "user":
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
	jiraCMD.Flags().StringArrayVar(&getCMD, "get", []string{}, "Please specify what you want (users, projects, roles etc...).")
	jiraCMD.Flags().StringArrayVar(&postCMD, "post", []string{}, "Please specify what you want to post (users, users_to_project, etc...).")

	jiraCMD.Flags().BoolVar(&export, "export", false, "Generate output files.")
	jiraCMD.Flags().StringVar(&outdir, "out-dir", "./assets", "Output directory")

	jiraCMD.Flags().StringVar(&url, "url", "", "Site API url. Ex: https://<site-name>.atlassian.net/rest/api/3")
	jiraCMD.Flags().StringVar(&user, "user", "", "User's token email")
	jiraCMD.Flags().StringVar(&token, "token", "", "Legacy token. Info: https://api.slack.com/custom-integrations/legacy-tokens")
	jiraCMD.MarkFlagRequired("url")
	jiraCMD.MarkFlagRequired("user")
	jiraCMD.MarkFlagRequired("token")

	rootCMD.AddCommand(jiraCMD)
}
