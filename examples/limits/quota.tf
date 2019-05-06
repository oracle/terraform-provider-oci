// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_limits_quota" "test_quota" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  description    = "Quotas for compute resources"
  name           = "TestQuotas"
  statements     = ["Set compute quotas to 0 in tenancy"]

  #Optional
  defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"

  freeform_tags = {
    "Department" = "Finance"
  }
}

data "oci_limits_quotas" "test_quotas" {
  #Required
  compartment_id = "${var.tenancy_ocid}"

  #Optional
  name  = "TestQuotas"
  state = "ACTIVE"
}
