resource "oci_core_instance" "my_instance" {
  availability_domain = "${var.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "my instance with FSS access"
  hostname_label = "myinstance"
  image = "${var.instance_image_ocid[var.region]}"
  shape = "${var.instance_shape}"
  subnet_id = "${oci_core_subnet.my_subnet.id}"
  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key)}"
  }
  timeouts {
    create = "60m"
  }
}

resource "null_resource" "mount_fss_on_instance" {
  depends_on = ["oci_core_instance.my_instance",
                "oci_file_storage_export.my_export_fs1_mt1"]
  provisioner "remote-exec" {
    connection {
      agent = false
      timeout = "15m"
      host = "${oci_core_instance.my_instance.public_ip}"
      user = "opc"
      private_key = "${file(var.ssh_private_key)}"
    }
    inline = [
      "sudo yum -y install nfs-utils > nfs-utils-install.log",
      "sudo mkdir -p /mnt/myfsspaths/fs1/path1",
      "sudo mount ${local.mount_target_1_ip_address}:${var.export_path_fs1_mt1} /mnt${var.export_path_fs1_mt1}",
    ]
  }
}