output "data_sync_public_ip" {
  value = "${oci_core_instance.rsync_fss.public_ip}"
}

output "data_sync_private_ip" {
  value = "${oci_core_instance.rsync_fss.private_ip}"
}
