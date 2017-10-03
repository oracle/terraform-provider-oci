output "KVM_HOST_PUBLIC_IP" {
  value = ["${oci_core_instance.KVM-HOST.*.public_ip}"]
}

output "KVM_HOST_PRIVATE_IP" {
  value = ["${oci_core_instance.KVM-HOST.*.private_ip}"]
}

output "KVM_GUEST_PUBLIC_IP" {
  value = ["${data.oci_core_vnic.KVM-mgmt-vnic.*.public_ip_address}"]
}

output "KVM_GUEST_PRIVATE_IP" {
  value = ["${data.oci_core_vnic.KVM-mgmt-vnic.*.private_ip_address}"]
}

output "KVM_GUEST_MAC_ADDRESS" {
  value = ["${data.oci_core_vnic.KVM-mgmt-vnic.*.mac_address}"]
}

output "KVM_GUEST_VNC_ADDRESS" {
  value = ["${data.oci_core_vnic.KVM-mgmt-vnic.*.public_ip_address}:${var.kvm_guest_vnc_port}"]
}

output "KVM_GUEST_VNC_PWD" {
  value = "${var.kvm_guest_vnc_pwd}"
}

output "KVM_FRONTEND_PUBLIC_IP" {
  value = ["${data.oci_core_vnic.KVM-frontend-vnic.*.public_ip_address}"]
}

output "KVM_FRONTEND_PRIVATE_IP" {
  value = ["${data.oci_core_vnic.KVM-frontend-vnic.*.private_ip_address}"]
}

output "KVM_BACKEND_PRIVATE_IP" {
  value = ["${data.oci_core_vnic.KVM-backend-vnic.*.private_ip_address}"]
}
