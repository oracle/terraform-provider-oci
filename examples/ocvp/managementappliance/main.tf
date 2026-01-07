// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.

# VCN comes with default route table, security list and DHCP options

variable "tenancy_ocid" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "config_file_profile" {}

variable "vcenter_secret_id" {}

variable "admin_secret_id" {}

variable "nsx_secret_id" {}

variable "sddc_id" {}

provider "oci" {
  region              = var.region
  tenancy_ocid        = var.tenancy_ocid
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile

}

data "oci_ocvp_sddc" "test_sddc" {
  sddc_id = "${var.sddc_id}"
}

resource "oci_ocvp_management_appliance" "test_management_appliance" {
  configuration {
    is_log_ingestion_enabled      = "false"
    is_metrics_collection_enabled = "false"
  }
  connections {
    credentials_secret_id = "${var.vcenter_secret_id}"
    type                  = "VCENTER"
  }
  connections {
    credentials_secret_id = "${var.admin_secret_id}"
    type                  = "ADMIN_VCENTER"
  }
  connections {
    credentials_secret_id = "${var.nsx_secret_id}"
    type                  = "NSX"
  }
  display_name = "displayName"
  lifecycle {
    ignore_changes = ["defined_tags"]
  }
  sddc_id = "${data.oci_ocvp_sddc.test_sddc.id}"
}

data "oci_ocvp_management_appliances" "test_management_appliances" {
  compartment_id = var.compartment_ocid
  display_name   = "displayName"

  filter {
    name = "id"
    values = [oci_ocvp_management_appliance.test_management_appliance.id]
  }

  management_appliance_id = "${oci_ocvp_management_appliance.test_management_appliance.id}"
  state                   = "ACTIVE"
}

data "oci_ocvp_management_appliance" "test_management_appliance" {
  management_appliance_id = oci_ocvp_management_appliance.test_management_appliance.id
}
