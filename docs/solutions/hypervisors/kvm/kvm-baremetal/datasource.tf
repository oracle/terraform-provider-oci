data "oci_identity_availability_domains" "ads" {
  compartment_id = "${var.tenancy_ocid}"
}

data "oci_core_vnic" "kvm-guest-vnic" {
  vnic_id = "${oci_core_vnic_attachment.kvm-guest-vnic-attachmnt.vnic_id}"
}

data "oci_core_vnic" "kvm-host-vnic" {
  vnic_id = "${oci_core_vnic_attachment.kvm-host-vnic-attachmnt.vnic_id}"
}
