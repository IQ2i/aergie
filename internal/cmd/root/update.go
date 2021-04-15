package root

import (
	"os"

	"github.com/iq2i/aergie/internal/http"
)

func GetLatestVersion() string {
	client := http.CreateClient()
	req, err := http.CreateRequest(os.Getenv("AE_UPDATE_DOMAIN") + "/latest")
	if err != nil {
		os.Exit(0)
	}

	response, err := http.ExecRequest(client, req)
	if err != nil {
		os.Exit(0)
	}

	return response
}
