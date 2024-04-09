resource "oci_containerengine_cluster" "test_cluster" {
  #Required
  compartment_id     = var.compartment_ocid
  kubernetes_version = "v1.27.2"
  name               = "cluster"
  vcn_id             = oci_core_vcn.test_vcn.id
}