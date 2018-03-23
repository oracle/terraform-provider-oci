resource "oci_core_instance" "dcos_master" {
    availability_domain = "${var.availability_domain}"
    compartment_id = "${var.compartment_ocid}"
    display_name  = "${var.dcos_cluster_name}-master-${var.display_name_prefix}-${count.index}"
    image = "${var.image}"
    shape = "${var.shape}"
    create_vnic_details {
      subnet_id = "${var.subnet_id}"
    }
    metadata {
      ssh_authorized_keys = "${var.ssh_public_key}"
    }
    timeouts {
      create = "10m"
    }

  count              = "${var.count}"

  connection {
      host = "${oci_core_instance.dcos_master.public_ip}"
      type = "ssh"
      user = "opc"
      private_key = "${file(var.ssh_private_key)}"
      timeout = "3m"
  }

  provisioner "local-exec" {
    command = "rm -rf ./do-install.sh"
  }

  provisioner "local-exec" {
    command = "echo ${format("MASTER_%02d", count.index)}=\"${oci_core_instance.dcos_master.private_ip}\" >> ips.txt"
  }

  provisioner "local-exec" {
    command = "while [ ! -f ./do-install.sh ]; do sleep 1; done"
  }

  provisioner "file" {
    source      = "./do-install.sh"
    destination = "/tmp/do-install.sh"
  }

  provisioner "file" {
    source      = "./setup.sh"
    destination = "$HOME/os_setup.sh"
  }
  provisioner "remote-exec" {
    inline = [
      "sudo systemctl stop firewalld",
      "sudo bash $HOME/os_setup.sh",
      "sudo bash /tmp/do-install.sh master",
#      "sudo firewall-cmd --zone=public --permanent --add-port=80/tcp",
#      "sudo firewall-cmd --zone=public --permanent --add-port=2181/tcp",
#      "sudo firewall-cmd --reload",
    ] 
  }
}
