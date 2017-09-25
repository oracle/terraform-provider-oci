resource "null_resource" "remote-exec" {
    depends_on = ["oci_core_instance.TFInstance","oci_core_volume_attachment.TFBlock0Attach"]
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "30m"
        host = "${data.oci_core_vnic.InstanceVnic.public_ip_address}"
        user = "opc"
        private_key = "${file(var.ssh_private_key)}"
    }
      inline = [
        "touch ~/IMadeAFile.Right.Here",
        "sudo iscsiadm -m node -o new -T ${oci_core_volume_attachment.TFBlock0Attach.iqn} -p ${oci_core_volume_attachment.TFBlock0Attach.ipv4}:${oci_core_volume_attachment.TFBlock0Attach.port}",
        "sudo iscsiadm -m node -o update -T ${oci_core_volume_attachment.TFBlock0Attach.iqn} -n node.startup -v automatic",
        "echo sudo iscsiadm -m node -T ${oci_core_volume_attachment.TFBlock0Attach.iqn} -p ${oci_core_volume_attachment.TFBlock0Attach.ipv4}:${oci_core_volume_attachment.TFBlock0Attach.port} -l >> ~/.bashrc"
      ]
    }
}

