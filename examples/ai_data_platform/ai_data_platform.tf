// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_ai_data_platform_ai_data_platform" "test_ai_data_platform" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  ai_data_platform_type  = var.ai_data_platform_ai_data_platform_type
  display_name           = var.ai_data_platform_display_name
  default_workspace_name = var.ai_data_platform_workspace_name
  freeform_tags          = var.ai_data_platform_freeform_tags
}

data "oci_ai_data_platform_ai_data_platforms" "test_ai_data_platforms" {
  #Required
  compartment_id      = var.compartment_id
}