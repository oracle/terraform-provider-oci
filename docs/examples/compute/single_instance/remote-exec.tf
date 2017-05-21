resource "null_resource" "remote-exec" {
    depends_on = ["baremetal_core_instance.TFInstance","baremetal_core_volume_attachment.TFBlock0Attach"]
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.baremetal_core_vnic.InstanceVnic.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    }
      inline = [
        "touch ~/IMadeAFile.Right.Here",
        "sudo iscsiadm -m node -o new -T ${baremetal_core_volume_attachment.TFBlock0Attach.iqn} -p ${baremetal_core_volume_attachment.TFBlock0Attach.ipv4}:${baremetal_core_volume_attachment.TFBlock0Attach.port}",
        "sudo iscsiadm -m node -o update -T ${baremetal_core_volume_attachment.TFBlock0Attach.iqn} -n node.startup -v automatic",
        "echo sudo iscsiadm -m node -T ${baremetal_core_volume_attachment.TFBlock0Attach.iqn} -p ${baremetal_core_volume_attachment.TFBlock0Attach.ipv4}:${baremetal_core_volume_attachment.TFBlock0Attach.port} -l >> ~/.bashrc"
      ]
    }
}

