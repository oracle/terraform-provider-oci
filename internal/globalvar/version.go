// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globalvar

import (
	"log"
)

const Version = "7.32.0"
const ReleaseDate = "2026-01-27"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
