// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"log"
)

const Version = "3.80.0"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
