// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file
variable "tenancy_ocid" {}

variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

variable "email_notification" {}

variable "idcs_access_token" {}

provider "oci" {
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

resource "oci_analytics_analytics_instance" "test_oce_instance_public" {
  compartment_id     = "${var.compartment_ocid}"
  description        = "OAC instance"
  email_notification = "${var.email_notification}"
  feature_set        = "ENTERPRISE_ANALYTICS"
  license_type       = "LICENSE_INCLUDED"

  capacity {
    capacity_type  = "OLPU_COUNT"
    capacity_value = 2
  }

  name              = "testoacinstance1"
  freeform_tags     = "${map("freeformkey", "freeformvalue")}"
  state             = "ACTIVE"
  idcs_access_token = "${var.idcs_access_token}"

  # Optional
  network_endpoint_details {
    #Required
    network_endpoint_type = "PUBLIC"

    #Optional
    whitelisted_ips = ["${oci_core_vcn.test_vcn.cidr_block}"]

    whitelisted_vcns = [{
      #Optional
      id              = "${oci_core_vcn.test_vcn.id}"
      whitelisted_ips = ["${oci_core_vcn.test_vcn.cidr_block}"]
    }]
  }
}

resource "oci_analytics_analytics_instance" "test_oce_instance_private" {
  compartment_id     = "${var.compartment_ocid}"
  description        = "OAC instance"
  email_notification = "${var.email_notification}"
  feature_set        = "ENTERPRISE_ANALYTICS"
  license_type       = "LICENSE_INCLUDED"

  capacity {
    capacity_type  = "OLPU_COUNT"
    capacity_value = 2
  }

  name              = "testoacinstance2"
  freeform_tags     = "${map("freeformkey", "freeformvalue")}"
  state             = "ACTIVE"
  idcs_access_token = "${var.idcs_access_token}"

  # Optional
  network_endpoint_details {
    #Required
    network_endpoint_type = "PRIVATE"

    #Optional
    subnet_id = "${oci_core_subnet.test_subnet.id}"
    vcn_id    = "${oci_core_vcn.test_vcn.id}"
  }
}

data "oci_analytics_analytics_instances" "test_instance" {
  compartment_id = "${var.compartment_ocid}"
}

output "test" {
  value = "${data.oci_analytics_analytics_instances.test_instance.id}"
}
