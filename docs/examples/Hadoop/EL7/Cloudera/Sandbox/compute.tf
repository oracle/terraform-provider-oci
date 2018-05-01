resource "oci_core_instance" "Sandbox" {
  count = "1"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "CDH Sandbox"
  hostname_label      = "CDH-Sandbox"
  image		      = "${var.image_ocid}"
  shape               = "VM.Standard2.8"
  subnet_id	      = "${oci_core_subnet.public.*.id[var.AD - 1]}"

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file("scripts/boot.sh"))}"
  }

  timeouts {
    create = "30m"
  }
}
