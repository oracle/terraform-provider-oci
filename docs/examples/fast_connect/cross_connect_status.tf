data "oci_core_cross_connect_status" "cross_connect_status" {
  #Required
  cross_connect_id = "${oci_core_cross_connect.cross_connect.id}"
}

output "cross_connect_status" {
  value = {
    id                    = "${data.oci_core_cross_connect_status.cross_connect_status.id}"
    interface_state       = "${data.oci_core_cross_connect_status.cross_connect_status.interface_state}"
    light_level_ind_bm    = "${data.oci_core_cross_connect_status.cross_connect_status.light_level_ind_bm}"
    light_level_indicator = "${data.oci_core_cross_connect_status.cross_connect_status.light_level_indicator}"
  }
}
