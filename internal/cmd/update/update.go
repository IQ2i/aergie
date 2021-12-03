package update

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/spf13/cobra"
)

func NewUpdateCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "self-update",
		Short: "Updates aergie to the latest version",

		SilenceUsage:  true,
		SilenceErrors: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			var latestVersion = getLatestVersion()

			if version == latestVersion {
				fmt.Printf("You are already using the latest available Aergie version %s", version)
				return nil
			}

			var releaseURL = fmt.Sprintf("https://github.com/iq2i/aergie/releases/download/%s/aergie_%s_%s.gz", latestVersion, runtime.GOOS, runtime.GOARCH)

			// TODO: get archive from GitHub

			// TODO: unzip archive

			// TODO: change permission of release file to be executable

			// TODO: mv release to current aergie location

			return nil
		},
	}

	return cmd
}

func getLatestVersion() string {
	resp, err := http.Get("https://api.github.com/repos/iq2i/aergie/releases/latest")
	if err != nil {
		log.Fatalln("Could not resolve host api.github.com")
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	return fmt.Sprintf("%v", result["tag_name"])
}
