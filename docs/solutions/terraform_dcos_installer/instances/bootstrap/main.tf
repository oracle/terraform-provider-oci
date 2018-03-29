resource "oci_core_instance" "dcos_bootstrap" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.BootstrapAD - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "${format("${var.dcos_cluster_name}-bootstrap-%02d", count.index)}"
    image = "${var.image}"
    shape = "${var.shape}"
    create_vnic_details {
    subnet_id   = "${var.subnets[var.BootstrapAD - 1]}"
    }
    metadata {
      ssh_authorized_keys = "${var.ssh_public_key}"
    }
    timeouts {
      create = "10m"
    }

  connection {
      host = "${oci_core_instance.dcos_bootstrap.public_ip}"
      type = "ssh"
      user = "opc"
      private_key = "${file(var.ssh_private_key)}"
      timeout = "3m"
  }

  provisioner "local-exec" {
    command = "rm -rf ./do-install.sh"
  }

  provisioner "local-exec" {
    command = "echo BOOTSTRAP=\"${oci_core_instance.dcos_bootstrap.private_ip}\" >> ips.txt"
  }

  provisioner "local-exec" {
    command = "echo CLUSTER_NAME=\"${var.dcos_cluster_name}\" >> ips.txt"
  }

  provisioner "remote-exec" {
    inline = [
      "wget -q -O dcos_generate_config.sh -P $HOME ${var.dcos_installer_url}",
      "mkdir $HOME/genconf",
    ]
  }

  provisioner "local-exec" {
    command = "./make-files.sh"
  }

  provisioner "local-exec" {
    command = "sed -i -e '/^- *$/d' ./config.yaml"
  }

  provisioner "file" {
    source      = "./ip-detect"
    destination = "$HOME/genconf/ip-detect"
  }

  provisioner "file" {
    source      = "./config.yaml"
    destination = "$HOME/genconf/config.yaml"
  }

  provisioner "file" {
    source      = "./setup.sh"
    destination = "$HOME/os_setup.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo bash $HOME/os_setup.sh",
      "sudo bash $HOME/dcos_generate_config.sh",
      "sudo docker run -d -p 4040:80 -v $HOME/genconf/serve:/usr/share/nginx/html:ro nginx 2>/dev/null",
      "sudo docker run -d -p 2181:2181 -p 2888:2888 -p 3888:3888 --name=dcos_int_zk jplock/zookeeper 2>/dev/null",
    ]
  }
}

