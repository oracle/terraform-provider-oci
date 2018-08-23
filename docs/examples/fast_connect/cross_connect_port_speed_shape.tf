data "oci_core_cross_connect_port_speed_shapes" "cross_connect_port_speed_shapes" {
  #Required
  compartment_id = "${var.compartment_ocid}"
}

output "cross_connect_port_speed_shapes" {
  value = "${data.oci_core_cross_connect_port_speed_shapes.cross_connect_port_speed_shapes.cross_connect_port_speed_shapes}"
}
