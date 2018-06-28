output "src_mount_target_private_ip_iad_ad1" {
  value = "${lookup(data.oci_core_private_ips.src_mt_private_ip_iad_ad1.private_ips[0], "ip_address")}"
}

output "dst_mount_target_private_ip_iad_ad2" {
  value = "${lookup(data.oci_core_private_ips.dst_mt_private_ip_iad_ad2.private_ips[0], "ip_address")}"
}

output "dst_mount_target_private_ip_phx_ad1" {
  value = "${lookup(data.oci_core_private_ips.dst_mt_private_ip_phx_ad1.private_ips[0], "ip_address")}"
}
