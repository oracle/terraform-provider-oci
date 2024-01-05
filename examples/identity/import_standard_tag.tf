// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "test_tag_namespace_name" {
  default = "Oracle-Standard"
}

resource "oci_identity_import_standard_tags_management" "test_import_standard_tags_management" {
  #Required
  compartment_id              = var.compartment_id
  standard_tag_namespace_name = var.test_tag_namespace_name
}

