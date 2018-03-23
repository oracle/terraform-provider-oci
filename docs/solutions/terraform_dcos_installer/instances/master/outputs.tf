# Output the private and public IPs of the instance
  
output "private_ips" {
  value = ["${oci_core_instance.dcos_master.*.private_ip}"]
}

output "public_ips" {
  value = ["${oci_core_instance.dcos_master.*.public_ip}"]
}
