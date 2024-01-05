// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "fusion_environment_data_masking_activity_is_resume_data_masking" {
  default = false
}

variable "fusion_environment_data_masking_activity_state" {
  default = "AVAILABLE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_fusion_apps_fusion_environment_data_masking_activity" "test_fusion_environment_data_masking_activity" {
  #Required
  fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id

  #Optional
  is_resume_data_masking = var.fusion_environment_data_masking_activity_is_resume_data_masking
}

data "oci_fusion_apps_fusion_environment_data_masking_activities" "test_fusion_environment_data_masking_activities" {
  #Required
  fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id

  #Optional
  state = var.fusion_environment_data_masking_activity_state
}
