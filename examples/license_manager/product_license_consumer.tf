data "oci_license_manager_product_license_consumers" "test_product_license_consumers" {
  #Required
  compartment_id     = var.tenancy_ocid
  product_license_id = oci_license_manager_product_license.test_product_license.id

  #Optional
  is_compartment_id_in_subtree = var.is_compartment_id_in_subtree
}