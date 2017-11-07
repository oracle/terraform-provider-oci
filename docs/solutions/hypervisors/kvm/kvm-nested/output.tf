output "KVM_HOST_PUBLIC_IP" {
  value = ["${oci_core_instance.kvm-host-instance.*.public_ip}"]
}

output "KVM_HOST_PRIVATE_IP" {
  value = ["${oci_core_instance.kvm-host-instance.*.private_ip}"]
}

output "KVM_GUEST_PUBLIC_IP" {
  value = ["${data.oci_core_vnic.kvm-guest-vnic.*.public_ip_address}"]
}

output "KVM_GUEST_PRIVATE_IP" {
  value = ["${data.oci_core_vnic.kvm-guest-vnic.*.private_ip_address}"]
}

output "KVM_GUEST_MAC_ADDRESS" {
  value = ["${data.oci_core_vnic.kvm-guest-vnic.*.mac_address}"]
}

output "KVM_GUEST_VNC_PWD" {
  value = "${var.kvm_guest_vnc_pwd}"
}

output "KVM_GUEST_VNC_PORT" {
  value = "${var.kvm_guest_vnc_port}"
}

output "KVM_HOST_SSH_KEY_PATH" {
  value = "${var.ssh_private_key_path}"
}

output "SSH_TUNNEL_TO_GUEST" {
  value = "ssh -i ${var.ssh_private_key_path} -L ${var.kvm_guest_vnc_port}:localhost:${var.kvm_guest_vnc_port} opc@${oci_core_instance.kvm-host-instance.public_ip}"
}
