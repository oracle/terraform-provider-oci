data "oci_license_manager_license_metric" "test_license_metrics" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  is_compartment_id_in_subtree = var.is_compartment_id_in_subtree
}