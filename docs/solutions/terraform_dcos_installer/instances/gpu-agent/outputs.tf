# Output the private and public IPs of the instance
  
output "private_ips" {
  value = ["${oci_core_instance.dcos_gpu_agent.*.private_ip}"]
}

output "instance_public_ips" {
  value = ["${oci_core_instance.dcos_gpu_agent.*.public_ip}"]
}
