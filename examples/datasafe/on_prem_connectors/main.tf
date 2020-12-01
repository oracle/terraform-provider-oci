variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_id" {
}


variable "on_prem_connector_defined_tags_value" {
  default = "value"
}

variable "on_prem_connector_description" {
  default = "description"
}

variable "on_prem_connector_display_name" {
  default = "displayName"
}

variable "on_prem_connector_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "on_prem_connector_on_prem_connector_lifecycle_state" {
  default = "INACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_on_prem_connector" "test_on_prem_connector" {
  compartment_id = var.compartment_id
  description   = var.on_prem_connector_description
  display_name  = var.on_prem_connector_display_name
  freeform_tags = var.on_prem_connector_freeform_tags
}

data "oci_data_safe_on_prem_connectors" "test_on_prem_connectors" {
  compartment_id = var.compartment_id
  display_name = var.on_prem_connector_display_name
  on_prem_connector_id = oci_data_safe_on_prem_connector.test_on_prem_connector.id
  on_prem_connector_lifecycle_state = var.on_prem_connector_on_prem_connector_lifecycle_state
}
