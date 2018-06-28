output "data_sync_src_public_ip" {
  value = "${oci_core_instance.fss_client_instance_src.public_ip}"
}

output "data_sync_src_private_ip" {
  value = "${oci_core_instance.fss_client_instance_src.private_ip}"
}

output "data_sync_dst_public_ip" {
  value = "${oci_core_instance.fss_client_instance_dst.public_ip}"
}

output "data_sync_dst_private_ip" {
  value = "${oci_core_instance.fss_client_instance_dst.private_ip}"
}
