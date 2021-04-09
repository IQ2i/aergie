package root

import (
	"github.com/iq2i/aergie/internal/http"
)

func GetLatestVersion() string {
	client := http.CreateClient()
	req := http.CreateRequest("https://get.aergie.com/latest")

	return http.ExecRequest(client, req)
}
