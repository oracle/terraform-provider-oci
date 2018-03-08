resource "oci_core_subnet" "subnet" {
  count      = "${var.ad_count}"
  cidr_block = "${var.cidr_block["${count.index}"]}"

  // Some terraform trickery required here.
  // Due to the fact that we need to specify placement for the Witness server, we're using additional variable ${ad_placement}
  // if set to "0" it's ignored and server is placed according to ${count}
  // Otherwise subnet is placed in the specific AD

  availability_domain = "${
    var.ad_deployment == "0" ?
    lookup(data.oci_identity_availability_domains.ADs.availability_domains["${count.index}"],"name")
    :
    lookup(data.oci_identity_availability_domains.ADs.availability_domains["${var.ad_deployment}"],"name")
    }"
  display_name = "${
    var.ad_deployment == "0" ?
    "${var.label_prefix}AD${"${count.index}" + 1}-${var.dns_label}.sub"
    :
    "${var.label_prefix}AD${"${var.ad_deployment}" + 1}-${var.dns_label}.sub"
    }"
  dns_label                  = "${var.dns_label}ad${"${count.index}" + 1}"
  compartment_id             = "${var.compartment_ocid}"
  vcn_id                     = "${var.vcn_id}"
  route_table_id             = "${var.route_table_id}"
  dhcp_options_id            = "${var.dhcp_options_id}"
  prohibit_public_ip_on_vnic = "${var.private}"
  security_list_ids          = ["${concat(var.security_list_id, var.additional_security_lists_ids)}"]
  provisioner "local-exec" {
    command = "sleep 5"
  }
}
