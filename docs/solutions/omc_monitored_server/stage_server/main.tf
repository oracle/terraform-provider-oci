module "oci_resources" {
  source = "../modules/datasources"
  tenancy_ocid = "${var.tenancy_ocid}"
}

data "template_file" "omc_cloudinit_script" {
  template =  "${file("${path.module}/userdata/omc-init.tpl")}"
  vars {
    ssh_public_key = "${file(var.ssh_public_key)}"
  }
}

resource "oci_core_instance" "omc_managed_instance" {
  availability_domain = "${lookup(module.oci_resources.ads[var.ad - 1],"name")}"
  compartment_id = "${lookup(module.oci_resources.compartments, var.compartment_name)}"
  display_name = "${var.server_display_name}"
  image = "${var.InstanceImageOCID[var.region]}"
  shape = "${var.shape_name}"
  subnet_id = "${var.subnet_id}"
  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key)}"
    user_data = "${base64encode(data.template_file.omc_cloudinit_script.rendered)}"
  }

  #Wait for cloud-init to complete before continuing
  provisioner "remote-exec" {
    inline = [
      "while [ ! -f /tmp/signal ]; do sleep 2; done",
    ]
    connection {
      host = "${self.public_ip}"
      type = "ssh"
      user = "opc"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

  # Copies the agentInstall.zip file to the /u01/omc directory
  provisioner "file" {
    source = "${var.omc_agent_path}"
    destination = "/omc/install/agentInstall.zip"
    connection {
      host = "${self.public_ip}"
      type = "ssh"
      user = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

  provisioner "remote-exec" {
    inline = [
      "unzip /omc/install/agentInstall.zip -d /omc/install",
      "chmod +x /omc/install/AgentInstall.sh",
      "/omc/install/AgentInstall.sh AGENT_TYPE=cloud_agent STAGE_LOCATION=/omc/stage -download_only AGENT_REGISTRATION_KEY=${var.omc_reg_key}"
    ]
    connection {
      host = "${self.public_ip}"
      type = "ssh"
      user = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }
}

