package version

import (
	"flag"
	"fmt"
	"os"
)

var (
	commitID  = "%COMMITID%"
	buildTime = "%BUILDID%"
)

func version() {
	buildVer := false
	flag.BoolVar(&buildVer, "version", false, "print build version and then exit")
	flag.Parse()

	if buildVer {
		fmt.Fprintf(os.Stdout, "BuildTime: %s\r\n", buildTime)
		fmt.Fprintf(os.Stdout, "CommitID: %s\r\n", commitID)
		os.Exit(0)
	}

}
