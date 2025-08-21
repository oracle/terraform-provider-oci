variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "attribute_set_type" {
  default = "IP_ADDRESS"
}

variable "attribute_set_values" {
  default = ["192.168.11.0"]
}

variable "display_name" {
  type    = string
  default = "IP addresses - AttributeSet"
}

variable "attribute_set_description" {
  default = "Attribute set for IP addresses"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_attribute_set" "test_attribute_set" {

  #Required
  attribute_set_type    = var.attribute_set_type
  attribute_set_values  = var.attribute_set_values
  compartment_id        = var.compartment_ocid
  display_name          = var.display_name

  #Optional
  description           = var.attribute_set_description

}