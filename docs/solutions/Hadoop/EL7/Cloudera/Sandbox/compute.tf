resource "oci_core_instance" "Sandbox" {
  count = "1"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "CDH Sandbox"
  hostname_label      = "CDH-Sandbox"
  shape               = "VM.Standard2.8"
  subnet_id	      = "${oci_core_subnet.public.*.id[var.AD - 1]}"

  source_details {
    source_type = "image"
    source_id = "${var.InstanceImageOCID[var.region]}"
    boot_volume_size_in_gbs = "${var.boot_volume_size}"
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file("scripts/boot.sh"))}"
  }

  timeouts {
    create = "30m"
  }
}
