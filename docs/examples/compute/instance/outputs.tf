# Output the private and public IPs of the instance
output "InstancePrivateIPs" {
  value = ["${oci_core_instance.TFInstance.*.private_ip}"]
}

output "InstancePublicIPs" {
  value = ["${oci_core_instance.TFInstance.*.public_ip}"]
}

# Output the boot volume IDs of the instance
output "BootVolumeIDs" {
  value = ["${oci_core_instance.TFInstance.*.boot_volume_id}"]
}

# Output the chap secret information for ISCSI volume attachments. This can be used to output
# CHAP information for ISCSI volume attachments that have "use_chap" set to true.
#output "IscsiVolumeAttachmentChapUsernames" {
#  value = ["${oci_core_volume_attachment.TFBlockAttach.*.chap_username}"]
#}
#
#output "IscsiVolumeAttachmentChapSecrets" {
#  value = ["${oci_core_volume_attachment.TFBlockAttach.*.chap_secret}"]
#}

