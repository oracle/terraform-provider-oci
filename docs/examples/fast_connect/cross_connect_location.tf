data "oci_core_cross_connect_locations" "cross_connect_locations" {
  #Required
  compartment_id = "${var.compartment_id}"
}

output "cross_connect_locations" {
  value = "${data.oci_core_cross_connect_locations.cross_connect_locations.cross_connect_locations}"
}
