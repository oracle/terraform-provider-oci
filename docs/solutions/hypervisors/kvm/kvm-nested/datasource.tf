data "oci_identity_availability_domains" "ads" {
  compartment_id = "${var.tenancy_ocid}"
}

data "oci_core_vnic_attachments" "kvm-host-vnics" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads.availability_domains[var.availability_domain - 1],"name")}"
  instance_id         = "${oci_core_instance.kvm-host-instance.id}"
}

data "oci_core_vnic" "kvm-guest-vnic" {
  vnic_id = "${oci_core_vnic_attachment.kvm-guest-vnic-attachmnt.vnic_id}"
}

data "oci_core_vnic" "kvm-host-vnic" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.kvm-host-vnics.vnic_attachments[0],"vnic_id")}"
}
