data "oci_identity_availability_domains" "ads" {
  compartment_id = "${var.tenancy_ocid}"
}

data "oci_core_images" "base-image" {
  compartment_id           = "${var.compartment_ocid}"
  operating_system         = "Oracle Linux"
  operating_system_version = "7.3"
}

data "oci_core_vnic" "kvm-guest-vnic" {
  vnic_id = "${oci_core_vnic_attachment.kvm-guest-vnic-attachmnt.vnic_id}"
}

data "oci_core_vnic" "kvm-host-vnic" {
  vnic_id = "${oci_core_vnic_attachment.kvm-host-vnic-attachmnt.vnic_id}"
}
