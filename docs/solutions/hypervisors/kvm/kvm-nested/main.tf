provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.ssh_private_key_path}"
  region           = "${var.region}"
}

resource "oci_core_instance" "kvm-host-instance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "${var.prefix}-kvm-host"
  image               = "${lookup(data.oci_core_images.base-image.images[0], "id")}"
  shape               = "${contains(slice(data.oci_core_shape.supported_shapes.shapes, 0,
    length(data.oci_core_shape.supported_shapes.shapes) - 1), var.instance_shape)  ? var.instance_shape : format("Only VM Shapes are supported on this demo. Selected: %s", var.instance_shape)  }"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.kvm-host-subnet.id}"
    hostname_label   = "kvm-host"
    display_name     = "kvm-host-vnic"
    assign_public_ip = true
  }

  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key_path)}"
  }

  timeouts {
    create = "10m"
  }
}

resource "oci_core_vnic_attachment" "kvm-guest-vnic-attachmnt" {
  instance_id  = "${oci_core_instance.kvm-host-instance.id}"
  display_name = "kvm-guest-vnic-attachmnt"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.kvm-host-subnet.id}"
    hostname_label   = "kvmguest"
    display_name     = "kvm-guest-vnic"
    assign_public_ip = true
  }
}

module "setup-kvm-hypervisor" {
  source                  = "./modules/kvm-installer-vm"
  private_key             = "${file(var.ssh_private_key_path)}"
  host                    = "${oci_core_instance.kvm-host-instance.public_ip}"
  qcow2_image_url         = "${var.kvm_image_url}"
  qcow2_image_target_path = "${var.kvm_image_path}"
  qcow2_image_filename    = "${var.kvm_image_name}"
  kvm_guest_domain_name   = "${var.kvm_guest_domain_name}"
  kvm_guest_memory        = "${var.kvm_guest_memory}"
  kvm_guest_vcpu          = "${var.kvm_guest_vcpu}"
  kvm_guest_os_type       = "${var.kvm_guest_os_type}"
  kvm_guest_vnc_port      = "${var.kvm_guest_vnc_port}"
  kvm_guest_vnc_pwd       = "${var.kvm_guest_vnc_pwd}"

  kvm_guest_vnic_mac_address = "${data.oci_core_vnic.kvm-guest-vnic.mac_address}"
  kvm_guest_emulation_mode   = "${var.kvm_emulation_mode}"
}
