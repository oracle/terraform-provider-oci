resource "oci_core_instance" "instance" {
  count               = "${var.ad_count}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains["${count.index}"],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "${var.dns_label}${"${count.index}" + 1}"

  #hostname_label      = "${var.dns_label}${"${count.index}" + 1}"
  hostname_label = ""
  image          = "${var.image_id}"
  shape          = "${var.shape}"

  create_vnic_details {
    subnet_id        = "${var.subnets["${count.index}"]}"
    assign_public_ip = false
  }

  timeouts {
    create = "60m"
  }

  provisioner "local-exec" {
    command = "sleep 10"
  }
}

data "oci_core_vnic_attachments" "instancevnics" {
  count          = "${var.ad_count}"
  compartment_id = "${var.compartment_ocid}"
  instance_id    = "${oci_core_instance.instance.*.id["${count.index}"]}"
}

resource "oci_core_volume_attachment" "db_attachment" {
  count           = "${var.ad_count}"
  attachment_type = "iscsi"
  compartment_id  = "${var.compartment_ocid}"
  instance_id     = "${oci_core_instance.instance.*.id["${count.index}"]}"
  volume_id       = "${var.db_volumes["${count.index}"]}"
}

resource "oci_core_volume_attachment" "log_attachment" {
  count           = "${var.ad_count}"
  attachment_type = "iscsi"
  compartment_id  = "${var.compartment_ocid}"
  instance_id     = "${oci_core_instance.instance.*.id["${count.index}"]}"
  volume_id       = "${var.log_volumes["${count.index}"]}"
}

resource "oci_core_volume_attachment" "backup_attachment" {
  count           = "${var.ad_count}"
  attachment_type = "iscsi"
  compartment_id  = "${var.compartment_ocid}"
  instance_id     = "${oci_core_instance.instance.*.id["${count.index}"]}"
  volume_id       = "${var.backup_volumes["${count.index}"]}"
}
