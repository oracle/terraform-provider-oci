# Outputing required info for users
output "Bridge Instance Public IP" {
  value = "${data.oci_core_vnic.BridgeInstanceVnic1.public_ip_address}"
}

output "PrivateInstance1 Private IP" {
  value = "${oci_core_instance.PrivateInstance.private_ip}"
}

output "PrivateInstance2 Private IP" {
  value = "${oci_core_instance.PrivateInstance2.private_ip}"
}

output "SSH login to the Bridge Instance" {
  value = "ssh -A opc@${data.oci_core_vnic.BridgeInstanceVnic1.public_ip_address}"
}

output "SSH login to the Private Instance-1 after logging into Bridge Instance as shown above" {
  value = "ssh -A opc@${oci_core_instance.PrivateInstance.private_ip}"
}

output "SSH login to the Private Instance-2 after logging into Bridge Instance as shown above" {
  value = "ssh -A opc@${oci_core_instance.PrivateInstance2.private_ip}"
}

output "Ping from PrivateInstance-1 to PrivateInstance-2" {
  value = "ping ${oci_core_instance.PrivateInstance2.private_ip} "
}
