// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globalvar

import (
	"log"
)

const Version = "4.97.0"
const ReleaseDate = "2022-10-26"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
