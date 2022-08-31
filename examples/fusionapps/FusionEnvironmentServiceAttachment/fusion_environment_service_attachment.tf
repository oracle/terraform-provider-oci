// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "fusion_environment_service_attachment_display_name" {
  default = "displayName"
}

variable "fusion_environment_service_attachment_service_instance_type" {
  default = "DIGITAL_ASSISTANT"
}

variable "fusion_environment_service_attachment_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_fusion_apps_fusion_environment_service_attachments" "test_fusion_environment_service_attachments" {
  #Required
  fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id

  #Optional
  display_name          = var.fusion_environment_service_attachment_display_name
  service_instance_type = var.fusion_environment_service_attachment_service_instance_type
  state                 = var.fusion_environment_service_attachment_state
}
