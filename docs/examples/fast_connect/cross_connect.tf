variable "cross_connect_display_name" { default = "displayName" }
variable "cross_connect_location_name" { default = "Equinix DC6, Ashburn, VA" }
variable "cross_connect_port_speed_shape_name" { default = "10 Gbps" }
variable "cross_connect_state" { default = "AVAILABLE" }

resource "oci_core_cross_connect" "test_cross_connect" {
	#Required
	compartment_id = "${var.compartment_id}"
	location_name = "${var.cross_connect_location_name}"
	port_speed_shape_name = "${var.cross_connect_port_speed_shape_name}"

	#Optional
	cross_connect_group_id = "${oci_core_cross_connect_group.test_cross_connect_group.id}"
	display_name = "${var.cross_connect_display_name}"
	#far_cross_connect_or_cross_connect_group_id = "${oci_core_far_cross_connect_or_cross_connect_group.test_far_cross_connect_or_cross_connect_group.id}"
	#near_cross_connect_or_cross_connect_group_id = "${oci_core_near_cross_connect_or_cross_connect_group.test_near_cross_connect_or_cross_connect_group.id}"

	#Set Cross Connect to Active to provision (required to provision virtual circuits).
	#You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up
	is_active = true
}

data "oci_core_cross_connects" "test_cross_connects" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	cross_connect_group_id = "${oci_core_cross_connect_group.test_cross_connect_group.id}"
	display_name = "${var.cross_connect_display_name}"
	#state = "${var.cross_connect_state}"
}

output "cross_connects" {
	value = "${data.oci_core_cross_connects.test_cross_connects.cross_connects}"
}
