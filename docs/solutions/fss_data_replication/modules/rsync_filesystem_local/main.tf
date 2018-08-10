data "template_file" "bootstrap" {
  template = "${file("${path.module}/bootstrap.tpl")}"

  vars {
    mount_point_path = "${var.mount_point_path}"

    src_export_path             = "${var.src_export_path}"
    src_mount_target_private_ip = "${var.src_mount_target_private_ip}"
    src_mount_path              = "${var.mount_point_path}${var.src_export_path}/"

    dst_export_path             = "${var.dst_export_path}"
    dst_mount_target_private_ip = "${var.dst_mount_target_private_ip}"
    dst_mount_path              = "${var.mount_point_path}${var.dst_export_path}/"

    data_sync_frequency = "${var.data_sync_frequency}"
  }
}

resource "oci_core_instance" "rsync_fss" {
  availability_domain = "${var.availability_domain}"
  compartment_id      = "${var.compartment_id}"
  shape               = "${var.instance_shape}"
  display_name        = "${var.instance_hostname}"

  create_vnic_details {
    subnet_id        = "${var.subnet_id}"
    display_name     = "${var.instance_hostname}"
    assign_public_ip = true
    hostname_label   = "${var.instance_hostname}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_id}"
  }

  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key_path)}"
    user_data           = "${base64encode(data.template_file.bootstrap.rendered)}"
  }

  timeouts {
    create = "60m"
  }
}

resource "null_resource" "verify_cloud_init" {
  depends_on = ["oci_core_instance.rsync_fss"]

  connection {
    agent       = false
    timeout     = "30m"
    host        = "${oci_core_instance.rsync_fss.public_ip}"
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
      "sudo cat /etc/cron.d/fss_sync_up_file_system",
      "sudo crontab -l",
      "echo 'finished'",
    ]
  }
}
