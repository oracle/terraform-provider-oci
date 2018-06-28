resource "oci_core_instance" "fss_client_instance_src" {
  provider            = "oci.src"
  availability_domain = "${var.src_availability_domain}"
  compartment_id      = "${var.compartment_id}"
  shape               = "${var.src_instance_shape}"
  display_name        = "${var.src_instance_hostname}"

  create_vnic_details {
    subnet_id        = "${var.src_subnet_id}"
    display_name     = "${var.src_instance_hostname}"
    assign_public_ip = true
    hostname_label   = "${var.src_instance_hostname}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.src_instance_image_id}"
  }

  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key_path)}"
    user_data           = "${base64encode(data.template_file.bootstrap_src.rendered)}"
  }

  timeouts {
    create = "60m"
  }
}

resource "null_resource" "verify_cloud_init_src" {
  depends_on = ["oci_core_instance.fss_client_instance_src"]

  connection {
    agent       = false
    timeout     = "30m"
    host        = "${oci_core_instance.fss_client_instance_src.public_ip}"
    user        = "opc"
    private_key = "${file(var.ssh_private_key_path)}"
  }

  provisioner "file" {
    source      = "${path.module}/scripts/cloud_init_checker.sh"
    destination = "~/cloud_init_checker.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "sh -x ~/cloud_init_checker.sh",
      "echo 'finished cloud_init'",
      "sudo cat /etc/cron.d/fss_sync_up_snapshot",
      "sudo crontab -l",
      "echo 'finished'",
    ]
  }
}

resource "oci_core_instance" "fss_client_instance_dst" {
  provider            = "oci.dst"
  availability_domain = "${var.dst_availability_domain}"
  compartment_id      = "${var.compartment_id}"
  shape               = "${var.dst_instance_shape}"
  display_name        = "${var.dst_instance_hostname}"

  create_vnic_details {
    subnet_id        = "${var.dst_subnet_id}"
    display_name     = "${var.dst_instance_hostname}"
    assign_public_ip = true
    hostname_label   = "${var.dst_instance_hostname}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.dst_instance_image_id}"
  }

  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key_path)}"
    user_data           = "${base64encode(data.template_file.bootstrap_dst.rendered)}"
  }

  timeouts {
    create = "60m"
  }
}

resource "null_resource" "upload_ssh_key_dst" {
  depends_on = ["oci_core_instance.fss_client_instance_src"]

  connection {
    agent       = false
    timeout     = "30m"
    host        = "${oci_core_instance.fss_client_instance_dst.public_ip}"
    user        = "opc"
    private_key = "${file(var.ssh_private_key_path)}"
  }

  provisioner "file" {
    source      = "${var.ssh_private_key_path}"
    destination = "~/.ssh/id_rsa"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo chmod 0600 ~/.ssh/id_rsa",
    ]
  }
}

resource "null_resource" "verify_cloud_init_dst" {
  depends_on = ["oci_core_instance.fss_client_instance_dst"]

  connection {
    agent       = false
    timeout     = "30m"
    host        = "${oci_core_instance.fss_client_instance_dst.public_ip}"
    user        = "opc"
    private_key = "${file(var.ssh_private_key_path)}"
  }

  provisioner "file" {
    source      = "${path.module}/scripts/cloud_init_checker.sh"
    destination = "~/cloud_init_checker.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "sh -x ~/cloud_init_checker.sh",
      "echo 'finished cloud_init'",
      "sudo cat /etc/cron.d/fss_sync_up_snapshot",
      "sudo crontab -l",
      "echo 'finished'",
    ]
  }
}
