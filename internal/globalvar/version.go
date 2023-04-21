// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globalvar

import (
	"log"
)

const Version = "4.118.0"
const ReleaseDate = "2023-04-26"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
