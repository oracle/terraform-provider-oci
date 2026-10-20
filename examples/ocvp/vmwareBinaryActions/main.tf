// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.

variable "region" {}

variable "config_file_profile" {}

variable "ocvp_sddc_id" {}

provider "oci" {
  auth                = "SecurityToken"
  region              = var.region
  config_file_profile = var.config_file_profile
}

data "oci_ocvp_retrieve_vmware_binaries" "test_retrieve_vmware_binaries" {
  sddc_id = var.ocvp_sddc_id
}

data "oci_ocvp_generate_vmware_binary_download_info" "test_generate_vmware_binary_download_info" {
  sddc_id                  = var.ocvp_sddc_id
  vmware_binary_file_name  = data.oci_ocvp_retrieve_vmware_binaries.test_retrieve_vmware_binaries.items[0].file_name
}

output "vmware_binary_count" {
  value = length(data.oci_ocvp_retrieve_vmware_binaries.test_retrieve_vmware_binaries.items)
}

output "vmware_binary_file_name" {
  value = data.oci_ocvp_generate_vmware_binary_download_info.test_generate_vmware_binary_download_info.file_name
}

output "vmware_binary_download_url" {
  value = data.oci_ocvp_generate_vmware_binary_download_info.test_generate_vmware_binary_download_info.url
}
