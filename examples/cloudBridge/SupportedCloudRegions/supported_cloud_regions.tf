variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "asset_source_type" {
  default = "AWS"
}

variable "name_contains" {
  default = "us-east-1"
}

data "oci_cloud_bridge_supported_cloud_regions" "test_supported_cloud_regions" {
  asset_source_type = var.asset_source_type
  name_contains     = var.name_contains
}