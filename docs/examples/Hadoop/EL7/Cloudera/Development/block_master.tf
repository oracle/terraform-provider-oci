###
### Block Volumes for Master & Utility Nodes - used to store CDH parcels & data
###

resource "oci_core_volume" "MasterVolume1" {
  count="${var.MasterNodeCount}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "CDH Master ${format("%01d", count.index+1)} Volume 1"
  size_in_gbs = "${var.blocksize_in_gbs}"
}

resource "oci_core_volume_attachment" "MasterAttachment1" {
  count="${var.MasterNodeCount}"
  attachment_type = "iscsi"
  compartment_id = "${var.compartment_ocid}"
  instance_id = "${oci_core_instance.MasterNode.*.id[count.index]}"
  volume_id = "${oci_core_volume.MasterVolume1.*.id[count.index]}"
}

resource "oci_core_volume" "UtilityVolume" {
  count="1"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "CDH Utility Volume"
  size_in_gbs = "${var.blocksize_in_gbs}"
}

resource "oci_core_volume_attachment" "UtilityAttachment" {
  count="1"
  attachment_type = "iscsi"
  compartment_id = "${var.compartment_ocid}"
  instance_id = "${oci_core_instance.UtilityNode.id}"
  volume_id = "${oci_core_volume.UtilityVolume.id}"
}

