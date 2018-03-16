resource "oci_core_private_ip" "ip2" {
  // Terraform sometimes can be stubborn and doesn't like referencing datasource  // as a count value. We're going to use ad_count value to count the number of hosts  // For larger number of hosts we could introduce a variable ${var.hosts_num} and use  // math interpolation to calculate number of additional addresses.

  count   = "${var.ad_count}"
  vnic_id = "${lookup(var.vnic_ids["${count.index}"],"vnic_id")}"

  display_name   = "${var.dns_label}${"${count.index}" + 1}-node2"
  hostname_label = "${var.dns_label}${"${count.index}" + 1}-node2"
}

resource "oci_core_private_ip" "ip3" {
  count   = "${var.ad_count}"
  vnic_id = "${lookup(var.vnic_ids["${count.index}"],"vnic_id")}"

  display_name   = "${var.dns_label}${"${count.index}" + 1}-node3"
  hostname_label = "${var.dns_label}${"${count.index}" + 1}-node3"
}
