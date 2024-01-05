// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "private_endpoint_reachable_ip_private_ip" {
  default = "privateIp"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_resourcemanager_private_endpoint_reachable_ips" "test_private_endpoint_reachable_ips" {
  #Required
  private_endpoint_id = oci_resourcemanager_private_endpoint.test_private_endpoint.id
  private_ip          = var.private_endpoint_reachable_ip_private_ip
}

