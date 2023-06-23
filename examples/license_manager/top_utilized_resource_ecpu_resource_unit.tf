variable "top_utilized_resource_ecpu_resource_unit_type" {
  default = "ECPU"
}

data "oci_license_manager_top_utilized_resources" "test_top_utilized_resources" {
  #Required
  compartment_id     = var.tenancy_ocid
  resource_unit_type = var.top_utilized_resource_ecpu_resource_unit_type

  #Optional
  is_compartment_id_in_subtree = var.is_compartment_id_in_subtree
}
