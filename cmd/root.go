package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var (
	outdir string
	export bool

	url  string
	user  string
	token  string

	getCMD  []string
	postCMD []string
)

var rootCMD = &cobra.Command{
	Use: "make",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute ...
func Execute() error {
	return rootCMD.Execute()
}

func generate(path string, artifact interface{}) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.Encode(artifact)
}
