output "vcn_id" {
  value = "${oci_core_virtual_network.vcn.id}"
}

output "ig_id" {
  value = "${oci_core_internet_gateway.ig.id}"
}

output "rt_id" {
  value = "${oci_core_route_table.rt.id}"
}

output "custom_dhcp_options_id" {
  value = "${oci_core_dhcp_options.custom_dhcp_options.id}"
}

output "internet_dhcp_options_id" {
  value = "${oci_core_dhcp_options.internet_dhcp_options.id}"
}
