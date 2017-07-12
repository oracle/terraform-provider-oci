# Output the private IP of the instance

output "DBNodePublicIP" {
value = ["${data.baremetal_core_vnic.DBNodeVnic.public_ip_address}"]
}
