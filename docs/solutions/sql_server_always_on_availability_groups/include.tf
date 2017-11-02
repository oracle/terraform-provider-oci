// Values below are retrieved from terraform.tfvars file
// You probably don't need to change anything here.

variable "tenancy_id" {}

variable "compartment_id" {}
variable "user_id" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "private_key_password" {}
variable "region" {}

variable "home_region" {}

//True by default. Solve inconsistency problems but creates performance issue on destroy.
variable "disable_auto_retries" {
  default = "false"
}

variable "label_prefix" {
  description = "To create unique identifier for multiple clusters in a compartment."
  type        = "string"
  default     = ""
}

// AD's lookup

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_id}"
}

// Provider declaration

provider "oci" {
  tenancy_ocid         = "${var.tenancy_id}"
  user_ocid            = "${var.user_id}"
  fingerprint          = "${var.fingerprint}"
  private_key_path     = "${var.private_key_path}"
  private_key_password = "${var.private_key_password}"
  region               = "${var.region}"
  disable_auto_retries = "${var.disable_auto_retries}"
}

// Home provider declaration

provider "oci" {
  alias                = "home"
  tenancy_ocid         = "${var.tenancy_id}"
  user_ocid            = "${var.user_id}"
  fingerprint          = "${var.fingerprint}"
  private_key_path     = "${var.private_key_path}"
  private_key_password = "${var.private_key_password}"
  region               = "${var.home_region}"
  disable_auto_retries = "${var.disable_auto_retries}"
}
