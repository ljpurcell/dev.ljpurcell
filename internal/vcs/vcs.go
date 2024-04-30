package vcs

import (
	"fmt"
	"runtime/debug"
)

func Version() string {
	var revision, time string
	var modified bool

	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				revision = s.Value
			case "vcs.time":
				time = s.Value
			case "vcs.modified":
				if s.Value == "true" {
					modified = true
				}
			}
		}
	}

	if modified {
		return fmt.Sprintf("%s-dirty at %s", revision, time)
	}

	return fmt.Sprintf("%s at %s", revision, time)
}
