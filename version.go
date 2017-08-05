package main

import (
	"log"
	"strconv"
	"strings"
)

// Current public release available on github, overwritten as build step, do not alter, see Makefile
const version = "v0.0.0"

func printVersion() {
	// increment build number programatically since this will be the next public release
	strs := strings.Split(version, ".")
	build, _ := strconv.ParseInt(strs[2], 10, 64)
	strs[2] = strconv.FormatInt(build+1, 10)
	log.Printf("[INFO] terraform-provider-oci %s\n", strings.Join(strs, "."))
}
