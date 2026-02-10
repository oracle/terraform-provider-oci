# Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  tenancy_ocid        = var.compartment_id
}