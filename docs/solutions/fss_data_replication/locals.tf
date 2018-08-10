locals {
  src_mt_private_ip_iad_ad1 = "${lookup(data.oci_core_private_ips.src_mt_private_ip_iad_ad1.private_ips[0], "ip_address")}"
}

locals {
  dst_mt_private_ip_iad_ad2 = "${lookup(data.oci_core_private_ips.dst_mt_private_ip_iad_ad2.private_ips[0], "ip_address")}"
}

locals {
  dst_mt_private_ip_phx_ad1 = "${lookup(data.oci_core_private_ips.dst_mt_private_ip_phx_ad1.private_ips[0], "ip_address")}"
}
