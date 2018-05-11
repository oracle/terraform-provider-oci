# Output the private and public IPs of the instance

output "Cloudera Manager Login" {
  value = ["http://${data.oci_core_vnic.sandbox_vnic.public_ip_address}:7180/cmf/"]
}

output "HUE Login" {
  value = ["http://${data.oci_core_vnic.sandbox_vnic.public_ip_address}:8888/"]
}

output "Cloudera Guided Demo" {
  value = ["http://${data.oci_core_vnic.sandbox_vnic.public_ip_address}/"]
}

output "Sandbox SSH" { 
  value = ["ssh -i ~/.ssh/id_rsa opc@${data.oci_core_vnic.sandbox_vnic.public_ip_address}"]
}

