output "id" {
  value = ["${oci_core_instance.instance.*.id}"]
}

output "vnic_ids" {
  // We need to flatten the list, sicnce it was returned with count and contains the maps
  value = "${flatten(data.oci_core_vnic_attachments.instancevnics.*.vnic_attachments)}"
}

output "private_ip" {
  value = "${oci_core_instance.instance.*.private_ip}"
}

output "iscsi_attachment_db" {
  value = "${oci_core_volume_attachment.db_attachment.*.ipv4}"
}

output "iscsi_attachment_log" {
  value = "${oci_core_volume_attachment.log_attachment.*.ipv4}"
}

output "iscsi_attachment_backup" {
  value = "${oci_core_volume_attachment.backup_attachment.*.ipv4}"
}
