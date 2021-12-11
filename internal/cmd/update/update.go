package update

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/iq2i/aergie/internal/file"
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

			var source = fmt.Sprintf("/tmp/ae_%s.gz", time.Now().Format("20060102150405"))
			if err := file.Download(source, releaseURL); err != nil {
				log.Fatal(err)
			}
			defer clean(source)

			var dest = strings.TrimSuffix(source, ".gz")
			if err := file.Uncompress(source, dest); err != nil {
				log.Fatal(err)
			}
			defer clean(dest)

			if err := os.Chmod(dest, 0755); err != nil {
				log.Fatal(err)
			}

			var executable, err = os.Executable()
			if err != nil {
				log.Fatal(err)
			}

			if err := os.Rename(dest, executable); err != nil {
				log.Fatal(err)
			}

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

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%v", result["tag_name"])
}

func clean(path string) {
	if err := file.Remove(path); err != nil {
		log.Fatal(err)
	}
}
