// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globalvar

import (
	"log"
)

<<<<<<< HEAD
const Version = "6.22.0"

const ReleaseDate = "2025-01-19"
=======
const Version = "6.21.0"

const ReleaseDate = "2024-12-22"
>>>>>>> 57b36061333 (Finalize changelog and release for version v6.21.0)

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
