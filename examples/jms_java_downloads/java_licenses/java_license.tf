// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "region" {}

variable "java_license_display_name" {
  default = "Oracle Technology Network"
}

variable "java_license_license_type" {
  default = "OTN"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  region           = var.region
}

data "oci_jms_java_downloads_java_licenses" "test_java_licenses" {

  #Optional
  display_name = var.java_license_display_name
  license_type = var.java_license_license_type
}

