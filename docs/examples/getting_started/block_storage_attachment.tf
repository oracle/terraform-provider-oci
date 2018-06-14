# Here we are attaching our block storage to one of the instances (the one on the private subnet)
resource "oci_core_volume_attachment" "MyBlockStorageAttachment" {
  attachment_type = "iscsi"
  compartment_id = "${var.compartment_ocid}"
  instance_id = "${oci_core_instance.MyInstanceInPrivateSubnetAD1.id}"
  volume_id = "${oci_core_volume.MyBlockStorage.id}"
}
