data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

data "oci_core_images" "BaseImage" {
  compartment_id           = "${var.compartment_ocid}"
  operating_system         = "Oracle Linux"
  operating_system_version = "7.3"
}

data "oci_core_vnic" "KVM-mgmt-vnic" {
  vnic_id = "${oci_core_vnic_attachment.kvm-mgmt-vnic-attachmnt.vnic_id}"
}

data "oci_core_vnic" "KVM-frontend-vnic" {
  vnic_id = "${oci_core_vnic_attachment.frontend-vnic-attachmnt.vnic_id}"
}

data "oci_core_vnic" "KVM-backend-vnic" {
  vnic_id = "${oci_core_vnic_attachment.backend-vnic-attachmnt.vnic_id}"
}
