# Output the private and public IPs of the instance
  
output "private_ips" {
  value = ["${oci_core_instance.dcos_bootstrap.*.private_ip}"]
}

output "instance_public_ips" {
  value = ["${oci_core_instance.dcos_bootstrap.*.public_ip}"]
}

output "instance_id" {
  value = "${oci_core_instance.dcos_bootstrap.id}"
}
