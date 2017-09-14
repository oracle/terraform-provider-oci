module "oci_resources" {
  source = "../modules/datasources"
  tenancy_ocid = "${var.tenancy_ocid}"
}

resource "oci_core_instance" "omc_managed_instance" {
  availability_domain = "${lookup(module.oci_resources.ads[var.ad - 1],"name")}"
  compartment_id = "${lookup(module.oci_resources.compartments, var.compartment_name)}"
  display_name = "${var.server_display_name}"
  hostname_label = "${var.hostname}"
  image = "${var.omc_custom_image_id}"
  shape = "${var.shape_name}"
  subnet_id = "${var.subnet_id}"
  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key)}"
  }
}

data "template_file" "omc_install_script" {
  template =  "${file("${path.module}/omc_config/install_omc.tpl")}"
  vars {
    registration_key = "${var.omc_reg_key}"
  }
}

resource "null_resource" "omc_instance_configure"{

  provisioner "file" {
    content = "${file("${path.module}/omc_config/omc_entity.json")}"
      destination = "/omc/stage/omc_entity.json"
    connection {
      host = "${oci_core_instance.omc_managed_instance.public_ip}"
      type = "ssh"
      user = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

  provisioner "file" {
    content = "${data.template_file.omc_install_script.rendered}"
    destination = "/omc/stage/omc_agent_install.sh"
    connection {
      host = "${oci_core_instance.omc_managed_instance.public_ip}"
      type = "ssh"
      user = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /omc/stage/omc_agent_install.sh",
      "/omc/stage/omc_agent_install.sh"
    ]
    connection {
      host = "${oci_core_instance.omc_managed_instance.public_ip}"
      type = "ssh"
      user = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

  provisioner "remote-exec" {
    when = "destroy"
    inline = [
      "/omc/app/cloud_agent/agent_inst/bin/omcli delete_entity agent /omc/stage/omc_entity_update.json"
    ]
    connection {
      host = "${oci_core_instance.omc_managed_instance.public_ip}"
      type = "ssh"
      user = "oracle"
      private_key = "${file(var.ssh_private_key)}"
    }
  }
}

