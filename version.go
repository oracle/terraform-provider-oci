package main

import (
	"log"
)

// Version injected during Makefile "release" build step, do not commit change
const Version = "0.0.0"

func printVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
