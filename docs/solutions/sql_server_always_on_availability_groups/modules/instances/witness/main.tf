resource "oci_core_instance" "instance" {
  count               = "1"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains["${var.ad_deployment}"],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "${var.dns_label}"
  hostname_label      = "${var.dns_label}"
  image               = "${var.image_id}"
  shape               = "${var.shape}"

  create_vnic_details {
    subnet_id        = "${var.subnets["0"]}"
    assign_public_ip = false
  }

  timeouts {
    create = "60m"
  }

  provisioner "local-exec" {
    command = "sleep 10"
  }
}

resource "oci_core_volume_attachment" "attachment" {
  count           = "1"
  attachment_type = "iscsi"
  compartment_id  = "${var.compartment_ocid}"
  instance_id     = "${oci_core_instance.instance.*.id["${count.index}"]}"
  volume_id       = "${var.witness_volumes["${count.index}"]}"
}
