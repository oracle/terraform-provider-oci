## DC/OS Public Agent
#
resource "oci_core_instance" "dcos_public_agent" {
    count = "${var.count}"
    availability_domain = "${var.availability_domain}"
    compartment_id = "${var.compartment_ocid}"
    display_name   = "${var.dcos_cluster_name}-public-agent-${var.display_name_prefix}-${count.index}"
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

  connection {
      host = "${oci_core_instance.dcos_public_agent.public_ip}"
      type = "ssh"
      user = "opc"
      private_key = "${file(var.ssh_private_key)}"
      timeout = "3m"
  }

  provisioner "local-exec" {
    ## Trigger dependency on "dcos_bootstrap" instance
    #
    command = "echo \"${var.dcos_bootstrap_instance_id}\""
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
      "sudo bash /tmp/do-install.sh slave_public",
    ] 
  }
}


