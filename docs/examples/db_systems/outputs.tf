# Output the private IP of the instance

output "DBNodePublicIP" {
value = ["${data.oci_core_vnic.DBNodeVnic.public_ip_address}"]
}
