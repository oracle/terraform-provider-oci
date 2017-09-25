resource "oci_core_instance" "TFInstance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "TFInstance"
  hostname_label = "instance1"
  image = "${lookup(data.oci_core_images.OLImageOCID.images[0], "id")}"
  shape = "${var.InstanceShape}"
  subnet_id = "${var.SubnetOCID}"
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file(var.BootStrapFile))}"
  }

  timeouts {
    create = "60m"
  }

  provisioner "chef" {
    server_url = "${var.chef_server}"
    node_name = "${var.chef_node_name}"
    run_list = "${var.chef_recipes}"
    user_name = "${var.chef_user}"
    user_key = "${file(var.chef_key)}"
    recreate_client = true
    fetch_chef_certificates = true
    connection {
      host = "${self.public_ip}"
      type = "ssh"
      user = "opc"
      private_key = "${var.ssh_private_key}"
      timeout = "3m"
    }
  }

  #You will need knife.rb in your current path in order for this command to complete successfully.
  provisioner "local-exec" {
    when = "destroy"
    on_failure = "continue"
    command = "knife node delete ${var.chef_node_name} -y",
  }

}
