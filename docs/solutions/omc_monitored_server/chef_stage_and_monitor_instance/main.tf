module "oci_resources" {
  source = "../modules/datasources"
  tenancy_ocid = "${var.tenancy_ocid}"
}

resource "oci_core_instance" "omc_managed_instance" {
  availability_domain = "${lookup(module.oci_resources.ads[var.ad - 1],"name")}"
  compartment_id = "${lookup(module.oci_resources.compartments, var.compartment_name)}"
  display_name = "${var.server_display_name}"
  image = "${var.omc_custom_image_id == "notset" ? var.InstanceImageOCID[var.region] : var.omc_custom_image_id}"
  shape = "${var.shape_name}"
  subnet_id = "${var.subnet_id}"
  hostname_label = "${var.hostname}"
  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key)}"
  }

  provisioner "chef" {
    attributes_json = "${file(var.json_attributes)}"
    run_list = "${var.chef_recipes}"
    node_name = "${var.chef_node_name}"
    server_url = "${var.chef_server}"
    user_name = "${var.chef_user}"
    user_key = "${file(var.chef_key)}"
    connection {
      host = "${self.public_ip}"
      type = "ssh"
      user = "opc"
      private_key = "${file(var.ssh_private_key)}"
      timeout = "3m"
    }
  }

  provisioner "remote-exec" {
    when = "destroy"
    on_failure = "continue"
    inline = [
      "/omc/app/cloud_agent/agent_inst/bin/omcli delete_entity agent /omc/stage/omc_entity.json"
    ]
    connection {
      host = "${self.public_ip}"
      type = "ssh"
      user = "omc"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

  #You will need knife.rb in your current path in order for this command to complete successfully.
  provisioner "local-exec" {
    when = "destroy"
    on_failure = "continue"
    command = "knife node delete ${var.chef_node_name} -y",
  }
}



