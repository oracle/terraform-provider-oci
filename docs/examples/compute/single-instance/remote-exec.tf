resource "null_resource" "remote-exec" {
    depends_on = ["baremetal_core_instance.TFInstance"]
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "10m"
        host = "${data.baremetal_core_vnic.InstanceVnic.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
      inline = [
        "touch ~/IMadeAFile.Right.Here",
      ]
    }
}

