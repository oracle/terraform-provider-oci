module "oci_resources" {
  source       = "../modules/datasources"
  tenancy_ocid = "${var.tenancy_ocid}"
}

data "template_file" "omc_cloudinit_script" {
  template = "${file("${path.module}/userdata/omc-init.tpl")}"
  vars {
    ssh_public_key = "${file(var.ssh_public_key)}"
  }
}

resource "oci_core_instance" "omc_managed_instance" {
  availability_domain = "${lookup(module.oci_resources.ads[var.ad - 1],"name")}"
  compartment_id      = "${lookup(module.oci_resources.compartments, var.compartment_name)}"
  display_name        = "${var.server_display_name}"
  image               = "${var.InstanceImageOCID[var.region]}"
  shape               = "${var.shape_name}"
  subnet_id           = "${var.subnet_id}"

  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key)}"
    user_data           = "${base64encode(data.template_file.omc_cloudinit_script.rendered)}"
  }

  #Wait for cloud-init to complete before continuing
  provisioner "remote-exec" {
    inline = [
      "while [ ! -f /tmp/signal ]; do sleep 5; done; echo 'cloud-init completed.' ",
    ]

    connection {
      host        = "${self.public_ip}"
      type        = "ssh"
      user        = "opc"
      timeout     = "10m"
      private_key = "${file(var.ssh_private_key)}"
    }
  }
}


resource "null_resource" "omc_instance_install" {
  provisioner "file" {
    source      = "${path.module}/omc_config/setup.sh"
    destination = "/opt/omc/installer/setup.sh"

    connection {
      host        = "${oci_core_instance.omc_managed_instance.public_ip}"
      type        = "ssh"
      user        = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

  provisioner "remote-exec" {
    inline = [
      "export TENANT_NAME=${var.omc_tennant_name}",
      "export OMC_URL=${var.omc_url}",
      "export AGENT_REPO_URL=${var.omc_agent_repo_url}",
      "export REGISTRATION_KEY=${var.omc_registration_key}",
      "chmod +x /opt/omc/installer/setup.sh",
      "/opt/omc/installer/setup.sh",
      "/opt/omc/omc-agent/agent_inst/bin/omcli status agent",
    ]
    connection {
      host        = "${oci_core_instance.omc_managed_instance.public_ip}"
      type        = "ssh"
      user        = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

  provisioner "file" {
    source     = "${path.module}/omc_config/omc_entity.json"
    destination = "/opt/omc/omc-agent/omc_entity.json"
    connection {
      host        = "${oci_core_instance.omc_managed_instance.public_ip}"
      type        = "ssh"
      user        = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

  provisioner "remote-exec" {
    inline = [
      "cat /opt/omc/omc-agent/omc_entity.json | jq '.entities[0].name=\"'$(hostname -f)'\"' | cat > /opt/omc/omc-agent/omc_entity_update.json",
      "/opt/omc/omc-agent/agent_inst/bin/omcli update_entity agent /opt/omc/omc-agent/omc_entity_update.json",
    ]
    connection {
      host        = "${oci_core_instance.omc_managed_instance.public_ip}"
      type        = "ssh"
      user        = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

   provisioner "remote-exec" {
    when = "destroy"
    inline = [
      "/opt/omc/omc-agent/agent_inst/bin/omcli delete_entity agent /opt/omc/omc-agent/omc_entity_update.json",
      "cmd_uninstall=`/opt/omc/omc-agent/agent_inst/bin/omcli status agent|grep \"Binaries\"|awk -F':' '{print $2}'`\"/sysman/install/AgentInstall.sh -deinstall\"",
      "echo $cmd_uninstall",
      "`$cmd_unistall`",
      "cat /tmp/AgentDeinstall*.log",
    ]
    connection {
      host        = "${oci_core_instance.omc_managed_instance.public_ip}"
      type        = "ssh"
      user        = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }
}


