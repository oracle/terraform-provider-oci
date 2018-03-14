resource "oci_core_volume" "TFBlock" {
  count = "${var.NumInstances * var.NumVolumesPerInstance}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "TFBlock${count.index}"
  size_in_gbs = "${var.DBSize}"
}

resource "oci_core_volume_attachment" "TFBlockAttach" {
    count = "${var.NumInstances * var.NumVolumesPerInstance}"
    attachment_type = "iscsi"
    compartment_id = "${var.compartment_ocid}"
    instance_id = "${oci_core_instance.TFInstance.*.id[count.index / var.NumVolumesPerInstance]}"
    volume_id = "${oci_core_volume.TFBlock.*.id[count.index]}"
}

