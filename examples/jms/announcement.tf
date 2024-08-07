// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "summary_contains" {
  default = "example summary value"
}

data "oci_jms_announcements" "test_jms_announcements" {
	#Optional
	summary_contains = var.summary_contains
	time_start = var.time_start
	time_end = var.time_end
}
