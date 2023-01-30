variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "registry_display_name" {
  default = "displayName"
}

variable "registry_description" {
  default = "description"
}

variable "registry_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "registry_state" {
  default = "ACTIVE"
}

provider "oci" {
  region = var.region
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_data_connectivity_registry" "test_registry" {
  #Required
  display_name = var.registry_display_name
  compartment_id = var.compartment_ocid

  #Optional
  description = var.registry_description
  freeform_tags = var.registry_freeform_tags

  lifecycle {
    ignore_changes = [
      defined_tags]
  }
}

data "oci_data_connectivity_registries" "test_registries" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  state = var.registry_state
}

data "oci_data_connectivity_registry" "test_registry" {
  #Required
  registry_id = oci_data_connectivity_registry.test_registry.id
}