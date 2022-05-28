data "oci_license_manager_top_utilized_product_licenses" "test_top_utilized_product_licenses" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  is_compartment_id_in_subtree = var.is_compartment_id_in_subtree
}

