// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globalvar

import (
	"log"
)

const Version = "8.18.0"
const ReleaseDate = "2026-06-10"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
