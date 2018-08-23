data "oci_core_cross_connect_locations" "cross_connect_locations" {
  #Required
  compartment_id = "${var.compartment_ocid}"
}

output "cross_connect_locations" {
  value = "${data.oci_core_cross_connect_locations.cross_connect_locations.cross_connect_locations}"
}
