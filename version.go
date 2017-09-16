package main

import (
	"log"
)

// Version injected during Makefile "release" build step, do not commit change
const version = "1.2.3"

func printVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", version)
}
