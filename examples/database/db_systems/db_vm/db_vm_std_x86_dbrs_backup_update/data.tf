data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_core_services" "test_services" {
  filter {
    name   = "name"
    regex  = "true"
    values = [".*Oracle.*Services.*Network"]
  }
}
