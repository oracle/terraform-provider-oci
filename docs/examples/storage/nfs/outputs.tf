# Output the private and public IPs of the instance

output "InstancePrivateIP" {
  value = ["${data.oci_core_vnic.InstanceVnic.private_ip_address}"]
}

output "InstancePublicIP" {
  value = ["${data.oci_core_vnic.InstanceVnic.public_ip_address}"]
}
