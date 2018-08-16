/*
 * This example file shows how to configure multiple oci providers to target different regions. 
 */

// These variables would commonly be defined as environment variablbe or sourced in a .env file 
variable "tenancy_ocid" {}

variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

// This provider has no alias and consequently will be used by resources that do not specify an alias 
provider "oci" {
  region           = "us-phoenix-1"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

// This provider defines an alias and is targetable by resources by including `provider = "oci.iad"`. 
provider "oci" {
  region           = "us-ashburn-1"
  alias            = "iad"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}
