// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "java_release_family_version" {
  default = "17"
}
variable "java_release_release_version" {
  default = "17.0.2"
}
variable "java_release_jre_security_status" {}
variable "java_release_license_type" {
  default = "OTN"
}
variable "java_release_release_type" {
  default = "CPU"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_jms_java_releases" "test_java_releases" {
  family_version      = var.java_release_family_version
  jre_security_status = var.java_release_jre_security_status
  license_type        = var.java_release_license_type
  release_type        = var.java_release_release_type
  release_version     = var.java_release_release_version
}

data "oci_jms_java_release" "test_java_release" {
  release_version     = var.java_release_release_version
}

