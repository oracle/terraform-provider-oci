# Output the private IP of the instance

output "DBNodePrivateIP" {
value = ["${data.baremetal_core_vnic.DBNodeVnic.private_ip_address}"]
}
