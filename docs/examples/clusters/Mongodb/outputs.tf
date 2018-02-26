# Output the private and public IPs of the instance

output "BastionPublicIP" {
value = ["${data.oci_core_vnic.BastionVnic.public_ip_address}"]
}

output "MongoDBAD1PrivateIP" {
value = ["${data.oci_core_vnic.MongoDBAD1Vnic.private_ip_address}"]
}

output "MongoDBAD2PrivateIP" {
value = ["${data.oci_core_vnic.MongoDBAD2Vnic.private_ip_address}"]
}
