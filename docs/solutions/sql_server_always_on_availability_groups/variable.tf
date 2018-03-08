provider "oci" {
  tenancy_ocid         = "${var.tenancy_ocid}"
  user_ocid            = "${var.user_ocid}"
  fingerprint          = "${var.fingerprint}"
  private_key_path     = "${var.private_key_path}"
  private_key_password = "${var.private_key_password}"
  region               = "${var.region}"
}

provider "oci" {
  alias		       = "home"
  tenancy_ocid         = "${var.tenancy_ocid}"
  user_ocid            = "${var.user_ocid}"
  fingerprint          = "${var.fingerprint}"
  private_key_path     = "${var.private_key_path}"
  private_key_password = "${var.private_key_password}"
  region               = "${var.home_region == "" ? var.region : var.home_region}"
}

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "private_key_password" {}
variable "region" {}

variable "home_region" {
	default = ""
}
variable "disable_auto_retries" {
  default = "false"
}

# This set of locals is primarily used to define the subnet structure to 
# deployed as part of SQL AO.  The subnet definition is optional - when the 
# subnet is deployed and no definition is present, a range will be selected
# automatically.  
#
# The subnet is allocated in three blocks:
# DMZ - Used to isolate the deployment
# Active Directory - Layer containing all AD required components
# SQL Servers - The servers themselves

locals {

  DMZ_prefix   = "${cidrsubnet("${var.vcn_cidr_block}", 2, 0)}"
  ADMIN_prefix = "${cidrsubnet("${var.vcn_cidr_block}", 2, 1)}"
  SQL_prefix   = "${cidrsubnet("${var.vcn_cidr_block}", 2, 2)}"

  DMZ_subnets_cidrs = [
    "${cidrsubnet("${local.DMZ_prefix}", 2, 0)}",
    "${cidrsubnet("${local.DMZ_prefix}", 2, 1)}",
    "${cidrsubnet("${local.DMZ_prefix}", 2, 2)}",
  ]

  ADMIN_subnets_cidrs = [
    "${cidrsubnet("${local.ADMIN_prefix}", 2, 0)}",
    "${cidrsubnet("${local.ADMIN_prefix}", 2, 1)}",
    "${cidrsubnet("${local.ADMIN_prefix}", 2, 2)}",
  ]

  SQL_subnets_cidrs = [
    "${cidrsubnet("${local.SQL_prefix}", 2, 0)}",
    "${cidrsubnet("${local.SQL_prefix}", 2, 1)}",
    "${cidrsubnet("${local.SQL_prefix}", 2, 2)}",
  ]

  Witness_subnets_cidrs = [
    "${cidrsubnet("${local.SQL_prefix}", 2, 3)}",
  ]
}
