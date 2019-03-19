// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.
/*resource "null_resource" "remote-exec" {
    provisioner "file" {
        connection {
            agent = false
            timeout = "10m"
            host = "${data.oci_core_vnic.db_node_vnic.public_ip_address}"
            user = "${var.host_user_name}"
            private_key = "${var.ssh_private_key}"
        }
        source = "./scripts/bootstrap.sh"
        destination = "~/bootstrap.sh"
    }    
    
    provisioner "remote-exec" {
        connection {
            agent = false
            timeout = "10m"
            host = "${data.oci_core_vnic.db_node_vnic.public_ip_address}"
            user = "${var.host_user_name}"
            private_key = "${var.ssh_private_key}"
        }
        inline = [
            "chmod +x ~/bootstrap.sh",
            "sudo ~/bootstrap.sh",
        ]
    }
}*/

