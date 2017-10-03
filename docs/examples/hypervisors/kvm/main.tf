provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.ssh_private_key_path}"
  region           = "${var.region}"
}

resource "oci_core_instance" "KVM-HOST" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "KVM-HOST-${var.customer_name}"
  image               = "${lookup(data.oci_core_images.BaseImage.images[0], "id")}"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.frontend-subnet.id}"
    hostname_label   = "bmcs-host"
    display_name     = "bmcs-host-vnic"
    assign_public_ip = true
  }

  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key_path)}"
  }

  timeouts {
    create = "10m"
  }
}

resource "oci_core_vnic_attachment" "kvm-mgmt-vnic-attachmnt" {
  instance_id  = "${oci_core_instance.KVM-HOST.id}"
  display_name = "kvm-mgmt-vnic-attachmnt"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.frontend-subnet.id}"
    hostname_label   = "kvmmanagement"
    display_name     = "kvm-mgmt-vnic"
    assign_public_ip = true
  }
}

resource "oci_core_vnic_attachment" "frontend-vnic-attachmnt" {
  instance_id  = "${oci_core_instance.KVM-HOST.id}"
  display_name = "frontend-vnic-attachmnt"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.frontend-subnet.id}"
    hostname_label   = "kvmfrontend"
    display_name     = "kvm-frontend-vnic"
    assign_public_ip = true
  }
}

resource "oci_core_vnic_attachment" "backend-vnic-attachmnt" {
  instance_id  = "${oci_core_instance.KVM-HOST.id}"
  display_name = "backend-vnic-attachmnt"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.backend-subnet.id}"
    hostname_label   = "kvmbackend"
    display_name     = "kvm-backend-vnic"
    assign_public_ip = false
  }
}

module "setup-kvm-hypervisor" {
  source                  = "./modules/kvm-hypervisor"
  private_key             = "${file(var.ssh_private_key_path)}"
  host                    = "${oci_core_instance.KVM-HOST.public_ip}"
  qcow2_image_url         = "${var.kvm_image_url}"
  qcow2_image_target_path = "${var.kvm_image_path}"
  qcow2_image_filename    = "${var.kvm_image_name}"
  kvm_guest_domain_name   = "${var.kvm_guest_domain_name}"
  kvm_guest_memory        = "${var.kvm_guest_memory}"
  kvm_guest_vcpu          = "${var.kvm_guest_vcpu}"
  kvm_guest_os_type       = "${var.kvm_guest_os_type}"
  kvm_guest_vnc_port      = "${var.kvm_guest_vnc_port}"
  kvm_guest_vnc_pwd       = "${var.kvm_guest_vnc_pwd}"

  kvm_guest_vnic_mac_address = "${data.oci_core_vnic.KVM-mgmt-vnic.mac_address}"
  kvm_guest_vnic_id          = "${data.oci_core_vnic.KVM-mgmt-vnic.id}"
  kvm_guest_emulation_mode   = "${var.kvm_emulation_mode}"
}
