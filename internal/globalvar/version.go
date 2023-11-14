// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globalvar

import (
	"log"
)

const Version = "5.21.0"
const ReleaseDate = "2023-11-19"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
