package main

import (
	"fmt"
	"paganotoni/jumpkit/config"
	"paganotoni/jumpkit/internal/gloves"

	"github.com/paganotoni/tailo"
)

func main() {
	err := gloves.Start(
		"cmd/app/main.go",

		// Run the tailo watcher so when changes are made to
		// the html code it rebuilds css.
		gloves.WithRunner(tailo.Watch),

		// Extensions to watch
		gloves.WatchExtension(config.GlovesExtensionsToWatch...),

		// Exclude paths from code reloading.
		gloves.ExcludePaths(config.GlovesExcludePaths...),
	)

	if err != nil {
		fmt.Println(err)
	}
}
