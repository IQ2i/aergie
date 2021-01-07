package build

import (
	"fmt"
	"os"
	"time"

	"github.com/iq2i/aergie/internal/logger"
)

// Version is set by build flags.
var Version = "dev"

// StringDate is set by build flags (format: Y-m-d H:i:s)
var StringDate = time.Now().Format("2006-01-02 15:04:05")

// Date is object init from StringDate
var Date time.Time

func init() {
	date, err := time.Parse("2006-01-02 15:04:05", StringDate)
	if err != nil {
		logger.Error(fmt.Errorf("Compiled time is invalid"))
		os.Exit(1)
	}
	Date = date
}
