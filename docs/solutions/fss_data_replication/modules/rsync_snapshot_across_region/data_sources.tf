data "template_file" "bootstrap_src" {
  template = "${file("${path.module}/bootstrap_src.tpl")}"

  vars {
    mount_point_path = "${var.mount_point_path}"

    src_export_path             = "${var.src_export_path}"
    src_mount_target_private_ip = "${var.src_mount_target_private_ip}"
    src_mount_path              = "${var.mount_point_path}${var.src_export_path}/"

    dst_export_path             = "${var.dst_export_path}"
    dst_mount_target_private_ip = "${var.dst_mount_target_private_ip}"
    dst_mount_path              = "${var.mount_point_path}${var.dst_export_path}/"

    snapshot_frequency = "${var.snapshot_frequency}"
  }
}

data "template_file" "bootstrap_dst" {
  template = "${file("${path.module}/bootstrap_dst.tpl")}"

  vars {
    src_host         = "${local.fss_client_instance_src_private_ip}"
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
