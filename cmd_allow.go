package main

import (
	"fmt"
	"os"
)

// `direnv allow [PATH_TO_RC]`
func Allow(env Env, args []string) (err error) {
	var rcPath string
	var config *Config
	if len(args) > 1 {
		rcPath = args[2]
	} else {
		if rcPath, err = os.Getwd(); err != nil {
			return
		}
	}

	if config, err = LoadConfig(env); err != nil {
		return
	}

	rc := FindRC(rcPath, config.AllowDir())
	if rc == nil {
		return fmt.Errorf(".envrc file not found")
	}
	return rc.Allow()
}
