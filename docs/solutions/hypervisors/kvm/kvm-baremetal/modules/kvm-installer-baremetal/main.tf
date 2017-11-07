data "template_file" "generate_virsh_attach" {
  template = "${file("${path.module}/generate_virsh_attach.sh.tpl")}"

  vars {
    vnic_mac_address = "${var.kvm_guest_vnic_mac_address}"
    vnic_id          = "${var.kvm_guest_vnic_id}"
    emulation_model  = "${var.kvm_guest_emulation_mode}"
  }
}

resource "null_resource" "install-kvm-dependencies" {
  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.private_key}"
    host        = "${var.host}"
    timeout     = "30m"
  }

  provisioner "file" {
    source      = "${path.module}/scripts/install-kvm-dependencies.sh"
    destination = "~/install-kvm-dependencies.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x ~/install-kvm-dependencies.sh",
      "sudo sh -x ~/install-kvm-dependencies.sh",
    ]
  }
}

resource "null_resource" "restart-instance" {
  depends_on = ["null_resource.install-kvm-dependencies"]

  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.private_key}"
    host        = "${var.host}"
    timeout     = "30m"
  }

  provisioner "remote-exec" {
    inline = [
      "echo waiting...",
      "echo restarting instance",
      "sudo shutdown -r now",
    ]
  }
}

resource "null_resource" "waiting_for_reboot" {
  depends_on = ["null_resource.restart-instance"]

  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.private_key}"
    host        = "${var.host}"
    timeout     = "30m"
  }

  provisioner "remote-exec" {
    inline = [
      "ls -la ~/",
    ]
  }
}

resource "null_resource" "configure-secondary-vnics" {
  depends_on = ["null_resource.waiting_for_reboot"]

  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.private_key}"
    host        = "${var.host}"
    timeout     = "30m"
  }

  provisioner "file" {
    source      = "${path.module}/scripts/configure-secondary-vnics.sh"
    destination = "~/configure-secondary-vnics.sh"
  }

  provisioner "file" {
    source      = "${path.module}/scripts/secondary_vnics.service"
    destination = "~/secondary_vnics.service"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo mv ~/configure-secondary-vnics.sh /usr/bin/configure-secondary-vnics.sh",
      "sudo mv ~/secondary_vnics.service /etc/systemd/system/secondary_vnics.service",
      "chmod 744 /usr/bin/configure-secondary-vnics.sh",
      "sudo systemctl daemon-reload",
      "sudo systemctl enable secondary_vnics.service",
      "sudo systemctl start secondary_vnics.service",
      "sudo ifconfig",
    ]
  }
}

resource "null_resource" "upload-kvm-qcow2-image" {
  depends_on = ["null_resource.configure-secondary-vnics"]

  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.private_key}"
    host        = "${var.host}"
    timeout     = "30m"
  }

  provisioner "remote-exec" {
    inline = [
      "wget ${var.qcow2_image_url}",
      "sudo mkdir ${var.qcow2_image_target_path}",
      "sudo mv ~/${var.qcow2_image_filename} ${var.qcow2_image_target_path}${var.qcow2_image_filename}",
    ]
  }
}

resource "null_resource" "create-kvm-domain" {
  depends_on = ["null_resource.upload-kvm-qcow2-image"]

  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.private_key}"
    host        = "${var.host}"
    timeout     = "30m"
  }

  provisioner "remote-exec" {
    inline = [
      "${data.template_file.generate_virsh_attach.rendered}",
      "ls -la ~/",
      "chmod 755 generate_virsh_attach.sh",
      "sudo sh generate_virsh_attach.sh",
      "ls -la ~/",
      "cat attach.xml",
      "sudo systemctl enable libvirtd",
      "sudo virt-install --arch=x86_64 --name=${var.kvm_guest_domain_name} --ram=${var.kvm_guest_memory} --cpu Haswell-noTSX --vcpus ${var.kvm_guest_vcpu} --hvm --video qxl --nonetwork --os-type ${var.kvm_guest_os_type} --noautoconsole --disk ${var.qcow2_image_target_path}${var.qcow2_image_filename},format=raw,bus=virtio --graphics vnc,port=${var.kvm_guest_vnc_port},listen=0.0.0.0,password=${var.kvm_guest_vnc_pwd} --import",
      "sudo virsh attach-device ${var.kvm_guest_domain_name} ~/attach.xml --config",
      "sudo virsh autostart ${var.kvm_guest_domain_name}",
    ]
  }
}

resource "null_resource" "start-kvm-domain" {
  depends_on = ["null_resource.create-kvm-domain"]

  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.private_key}"
    host        = "${var.host}"
    timeout     = "30m"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo virsh destroy ${var.kvm_guest_domain_name}",
      "sudo virsh start ${var.kvm_guest_domain_name}",
    ]
  }
}
