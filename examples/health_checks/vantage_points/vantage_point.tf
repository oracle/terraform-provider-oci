// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

# variable "vantage_point_display_name" {
#   default = "displayName"
# }

variable "vantage_point_name" {
  default = "aws-bom"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_health_checks_vantage_points" "test_vantage_points" {
  #Optional  # display_name = var.vantage_point_display_name  # name         = var.vantage_point_name
}

output "vantage_points" {
  value = data.oci_health_checks_vantage_points.test_vantage_points.health_checks_vantage_points
}

