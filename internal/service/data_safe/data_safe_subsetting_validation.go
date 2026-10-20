// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validateSubsettingParallelDegree = validation.StringMatch(
	regexp.MustCompile(`(?i)^(NONE|DEFAULT|[1-9][0-9]*)$`),
	"must be NONE, DEFAULT, or a positive integer",
)
