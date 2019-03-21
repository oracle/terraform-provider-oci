// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

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

output "InstanceDevices" {
  value = ["${data.oci_core_instance_devices.TFInstanceDevices.devices}"]
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

// Outputs
output "silver_policy_id" {
  value = "${data.oci_core_volume_backup_policies.test_predefined_volume_backup_policies.volume_backup_policies.0.id}"
}

output "attachment_instance_id" {
  value = "${data.oci_core_boot_volume_attachments.TFBootVolumeAttachments.*.instance_id}"
}
