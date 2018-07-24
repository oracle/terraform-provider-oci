# Output the private and public IPs of the instance

output "4 - Cloudera Manager Login" {
  value = <<END

	http://${data.oci_core_vnic.sandbox_vnic.public_ip_address}:7180/cmf/

END
}

output "3 - HUE Login" {
  value = <<END

	http://${data.oci_core_vnic.sandbox_vnic.public_ip_address}:8888/

END
}

output "2 - Cloudera Guided Demo" {
  value = <<END
	
	http://${data.oci_core_vnic.sandbox_vnic.public_ip_address}/

END
}

output "1 - Sandbox SSH" { 
  value = <<END

	ssh -i ~/.ssh/id_rsa opc@${data.oci_core_vnic.sandbox_vnic.public_ip_address}

END
}

