# Output the private and public IPs of the instance
  
output "private_ips" {
  value = ["${oci_core_instance.dcos_public_agent.*.private_ip}"]
}

output "instance_public_ips" {
  value = ["${oci_core_instance.dcos_public_agent.*.public_ip}"]
}
