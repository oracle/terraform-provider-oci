// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "log_saved_search_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

resource "oci_logging_log_saved_search" "test_log_saved_search" {
  #Required
  compartment_id = var.compartment_ocid
  name           = "exampleLogSavedSearch"
  query          = "exampleQuery"

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = var.log_group_defined_tags_value
  }
  description = "description"

  freeform_tags = var.log_saved_search_freeform_tags
}

data "oci_logging_log_saved_searches" "test_log_saved_searches" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  name           = "exampleLogSavedSearch"
  log_saved_search_id = oci_logging_log_saved_search.test_log_saved_search.id
}
