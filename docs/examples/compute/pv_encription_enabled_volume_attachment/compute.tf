resource "oci_core_instance" "TFInstance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFInstance"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.ExampleSubnet.id}"
    display_name     = "primaryvnic"
    assign_public_ip = true
    hostname_label   = "tfexampleinstance"
  }

  source_details {
    source_type = "image"
    source_id   = "${lookup(data.oci_core_images.TFSupportedShapeImages.images[0], "id")}"
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "60m"
  }

  is_pv_encryption_in_transit_enabled = "true"
}

resource "oci_core_volume" "TFVolume" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "display_name"
}

resource "oci_core_volume_attachment" "TFVolumeAttachment" {
  attachment_type                     = "paravirtualized"
  compartment_id                      = "${var.compartment_ocid}"
  instance_id                         = "${oci_core_instance.TFInstance.id}"
  volume_id                           = "${oci_core_volume.TFVolume.id}"
  display_name                        = "tf-vol-attach"
  is_read_only                        = true
  is_pv_encryption_in_transit_enabled = true
}
