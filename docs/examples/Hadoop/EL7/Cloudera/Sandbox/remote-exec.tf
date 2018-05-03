resource "null_resource" "cdh-setup" {
    depends_on = ["oci_core_instance.Sandbox"]
    provisioner "file" {
      source = "scripts/sandbox.sh"
      destination = "/home/opc/sandbox.sh"
      connection {
        agent = false
        timeout = "10m"
        host = "${data.oci_core_vnic.sandbox_vnic.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
    } 
    }
    provisioner "remote-exec" {
      connection {
        agent = false
        timeout = "10m"
        host = "${data.oci_core_vnic.sandbox_vnic.public_ip_address}"
        user = "opc"
        private_key = "${var.ssh_private_key}"
      }
      inline = [
	"chmod +x /home/opc/sandbox.sh",
	"sudo /home/opc/sandbox.sh",
	]
    }
}

