# Output the private and public IPs of the instance
output "InstancePrivateIPs" {
value = ["${oci_core_instance.TFInstance.*.private_ip}"]
}

output "InstancePublicIPs" {
value = ["${oci_core_instance.TFInstance.*.public_ip}"]
}
