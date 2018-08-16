variable "cross_connect_group_display_name" {
  default = "displayName"
}

variable "cross_connect_group_state" {
  default = "AVAILABLE"
}

resource "oci_core_cross_connect_group" "test_cross_connect_group" {
  #Required
  compartment_id = "${var.compartment_id}"

  #Optional
  display_name = "${var.cross_connect_group_display_name}"
}

data "oci_core_cross_connect_groups" "test_cross_connect_groups" {
  #Required
  compartment_id = "${var.compartment_id}"

  #Optional
  display_name = "${var.cross_connect_group_display_name}"

  #state = "${var.cross_connect_group_state}"
}

output "cross_connect_groups" {
  value = "${data.oci_core_cross_connect_groups.test_cross_connect_groups.cross_connect_groups}"
}
