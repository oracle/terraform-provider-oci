variable "configuration_email_ids" {
  default = ["testuser@oracle.com"]
}

resource "oci_license_manager_configuration" "test_configuration" {
  #Required
  compartment_id = var.tenancy_ocid
  #Optional
  email_ids      = var.configuration_email_ids
}

data "oci_license_manager_configuration" "test_configurations" {
  #Required
  compartment_id = var.tenancy_ocid
}