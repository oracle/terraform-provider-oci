# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Gets the boot volume attachments for each instance
data "oci_core_boot_volume_attachments" "TFBootVolumeAttachments" {
  depends_on          = ["oci_core_instance.TFInstance"]
  count               = "${var.NumInstances}"
  availability_domain = "${oci_core_instance.TFInstance.*.availability_domain[count.index]}"
  compartment_id      = "${var.compartment_ocid}"

  instance_id = "${oci_core_instance.TFInstance.*.id[count.index]}"
}
