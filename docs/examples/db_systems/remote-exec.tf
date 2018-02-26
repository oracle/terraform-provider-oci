resource "null_resource" "remote-exec" {
	provisioner "file" {
		connection {
			agent = false
			timeout = "10m"
			host = "${data.oci_core_vnic.DBNodeVnic.public_ip_address}"
			user = "${var.HostUserName}"
			private_key = "${var.ssh_private_key}"
		}
        source = "./scripts/bootstrap.sh"
        destination = "~/bootstrap.sh"
    }    
    
    provisioner "remote-exec" {
        connection {
			agent = false
			timeout = "10m"
			host = "${data.oci_core_vnic.DBNodeVnic.public_ip_address}"
			user = "${var.HostUserName}"
			private_key = "${var.ssh_private_key}"
        }
        inline = [
			"chmod +x ~/bootstrap.sh",
			"sudo ~/bootstrap.sh",
        ]
    }
}
