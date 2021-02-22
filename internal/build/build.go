package build

import (
	"time"
)

// Version is set by build flags.
var Version = "dev"

// Date is is set by build flags (format: Y-m-d H:i:s)
var Date = time.Now().Format("2006-01-02 15:04:05")
