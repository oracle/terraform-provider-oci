// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globalvar

import (
	"log"
)

const Version = "5.32.0"
const ReleaseDate = "2024-03-06"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
