// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globalvar

import (
	"log"
)

const Version = "8.10.0"
const ReleaseDate = "2026-04-15"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
