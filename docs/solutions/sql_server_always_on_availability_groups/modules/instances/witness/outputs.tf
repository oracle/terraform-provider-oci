output "id" {
  value = "${oci_core_instance.instance.id}"
}

output "private_ip" {
  value = "${oci_core_instance.instance.*.private_ip}"
}

output "iscsi_attachment" {
  value = "${oci_core_volume_attachment.attachment.*.ipv4}"
}
