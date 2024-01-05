// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_limits_quota" "quota_rd" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  description    = "Quotas for Resource Discorvery"
  name           = "TestQuotasRD"
  statements     = ["Set notifications quota topic-count to 99 in tenancy"]

  freeform_tags = {
    "Department" = "Finance"
  }
}
