// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"log"
)

const Version = "3.53.0"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
