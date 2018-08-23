resource "oci_core_volume" "TFBlock" {
  count               = "${var.NumInstances * var.NumIscsiVolumesPerInstance}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFBlock${count.index}"
  size_in_gbs         = "${var.DBSize}"
}

resource "oci_core_volume_attachment" "TFBlockAttach" {
  count           = "${var.NumInstances * var.NumIscsiVolumesPerInstance}"
  attachment_type = "iscsi"
  compartment_id  = "${var.compartment_ocid}"
  instance_id     = "${oci_core_instance.TFInstance.*.id[count.index / var.NumIscsiVolumesPerInstance]}"
  volume_id       = "${oci_core_volume.TFBlock.*.id[count.index]}"

  # Set this to enable CHAP authentication for an ISCSI volume attachment. The oci_core_volume_attachment resource will
  # contain the CHAP authentication details via the "chap_secret" and "chap_username" attributes.
  #use_chap = true

  # Set this to attach the volume as read-only.
  #is_read_only = true
}

resource "oci_core_volume" "TFBlockParavirtualized" {
  count               = "${var.NumInstances * var.NumParavirtualizedVolumesPerInstance}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFBlockParavirtualized${count.index}"
  size_in_gbs         = "${var.DBSize}"
}

resource "oci_core_volume_attachment" "TFBlockAttachParavirtualized" {
  count           = "${var.NumInstances * var.NumParavirtualizedVolumesPerInstance}"
  attachment_type = "paravirtualized"
  compartment_id  = "${var.compartment_ocid}"
  instance_id     = "${oci_core_instance.TFInstance.*.id[count.index / var.NumParavirtualizedVolumesPerInstance]}"
  volume_id       = "${oci_core_volume.TFBlockParavirtualized.*.id[count.index]}"

  # Set this to attach the volume as read-only.
  #is_read_only = true
}
