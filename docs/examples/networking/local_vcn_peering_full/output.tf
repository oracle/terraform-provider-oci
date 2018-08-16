# You can test the peering connection by ssh-ing into an instance (using the public_ip) and doing a ping command to the private IP address of the other instance"]

output "requestorInstancePublicIP" {
  value = ["${oci_core_instance.requestor_instance.public_ip}"]
}

output "requestorInstancePrivateIP" {
  value = ["${oci_core_instance.requestor_instance.private_ip}"]
}

output "acceptorInstancePublicIP" {
  value = ["${oci_core_instance.acceptor_instance.public_ip}"]
}

output "acceptorInstancePrivateIP" {
  value = ["${oci_core_instance.acceptor_instance.private_ip}"]
}
