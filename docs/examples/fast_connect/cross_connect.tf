resource "oci_core_cross_connect" "cross_connect" {
  #Required
  compartment_id        = "${var.compartment_ocid}"
  location_name         = "${data.oci_core_cross_connect_locations.cross_connect_locations.cross_connect_locations.0.name}"
  port_speed_shape_name = "${data.oci_core_cross_connect_port_speed_shapes.cross_connect_port_speed_shapes.cross_connect_port_speed_shapes.0.name}"

  #Optional
  cross_connect_group_id = "${oci_core_cross_connect_group.cross_connect_group.id}"
  display_name           = "${var.cross_connect_display_name}"

  #far_cross_connect_or_cross_connect_group_id = "${oci_core_far_cross_connect_or_cross_connect_group.far_cross_connect_or_cross_connect_group.id}"
  #near_cross_connect_or_cross_connect_group_id = "${oci_core_near_cross_connect_or_cross_connect_group.near_cross_connect_or_cross_connect_group.id}"

  #Set Cross Connect to Active to provision (required to provision virtual circuits).
  #You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up
  is_active = true
}

data "oci_core_cross_connects" "cross_connects" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  cross_connect_group_id = "${oci_core_cross_connect_group.cross_connect_group.id}"
  display_name           = "${var.cross_connect_display_name}"

  #state = "${var.cross_connect_state}"
}

output "cross_connects" {
  value = "${data.oci_core_cross_connects.cross_connects.cross_connects}"
}
