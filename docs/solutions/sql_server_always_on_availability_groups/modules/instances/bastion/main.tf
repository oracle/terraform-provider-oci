resource "oci_core_instance" "instance" {
  count               = "${var.ad_count}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains["${count.index}"],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "${var.dns_label}${"${count.index}" + 1}"
  hostname_label      = "${var.dns_label}${"${count.index}" + 1}"
  image               = "${var.image_id}"
  shape               = "${var.shape}"

  create_vnic_details {
    subnet_id              = "${var.subnets["${count.index}"]}"
    skip_source_dest_check = true
    assign_public_ip       = true
  }

  timeouts {
    create = "60m"
  }

  provisioner "local-exec" {
    command = "sleep 10"
  }
}
