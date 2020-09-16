// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

# This is an example of how to resize a volume that’s attached and online without having to reattach it.
# The null provisioner ensures that a rescan command is run when the volume size has changed

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}

variable "ssh_public_key" {
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDM+NcnBp1i4iKNeYf853h0UA2GVLOaa+85lZGykQJMzOWcS3oB3OHVN1nRrS6foQQ7EuNpEMlxfKESCFKh4UQynzb3pWmy/dczPP/EwQY061v0uPJ5tb2SlYgExFt+gzgBpKejEVdN+OH8Jt63+YIadaSbPHdjUYU2xckMXbeYe0mGXJEUdAtcurML81V0BfMEA7TCPVxXf8Fl2JcIHhuMdJv/VfjxAIqA8tj7fI2hqFs4Cf2zMF4nKvHHgkDCXIiqhjOY+R2B44eQHXzXtTlqUr5XgHD3tGbvt+mcfuQr/eV1pnBbIteop/XHpyS0MKJctfeod8+glEukjSfzQwhL0hJ8JVxlidOcoYEIrAIKm0RLDYd4kO5FvChUps1yzm+bVYgQrFkUtzUu9oMpqc053Y0CPJ5YZu50nlx8syzhXtDKItXLifFrf8ekcmsPA2X6G9/v/YXcVUX4Rq1fftwua4Ftl4JQU/3A6Jiw5OSEPz8jf4Z5edGNcBCnPzeCdFkpxbkN/XFNb0azDq+6ke1XhKdj+WrWZedJSXzxP8Iur2ZVlMLIytFhypnZj82hImCRJnXGthrxvklzWT6Wm2znzUvZL0kUbZVSRCIA0oULc1l9pX22zgjfqtifjPiCH+d41FUINBKBHwkRObQ3xQ/RttRqYDYLjSPphqR3phO6fw== <zexin.wan@oracle.com>"
}

variable "ssh_private_key" {
  default = "~/.ssh/id_rsa"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

variable "instance_image_ocid" {
  type = "map"

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "CentOS-6.10-2020.04.21-0"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaajubunwtkoob57bkjgsdinb6xjqgggwkn7hjt3j3lqdtrh72jcnaq"

    uk-london-1  = "ocid1.image.oc1.uk-london-1.aaaaaaaaqxjy3ushb72q45cu2ekzvllnkgaf2pdnqyw7q3sv3krybzlcarxa"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaa27fldayogs4sre3l6arufehn6ist4u77hbylux5oisnarijlne7a"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaa66qzvli64ojblgjgzrqcutcynjqfinoiyib3u7vn4stauloyig3q"
  }
}

variable "size" {
  default = "130" # size in GBs
}

resource "oci_core_instance" "test_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstance1"
  shape               = var.instance_shape

  shape_config {
    ocpus = 1
  }

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "tfexampleinstance1"
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]

    # Apply this to set the size of the boot volume that's created for this instance.
    # Otherwise, the default boot volume size of the image is used.
    # This should only be specified when source_type is set to "image".
    #boot_volume_size_in_gbs = "60"
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
    user_data           = base64encode(file("./userdata/bootstrap"))
  }

  # Apply the following flag only if you wish to preserve the attached boot volume upon destroying this instance
  # Setting this and destroying the instance will result in a boot volume that should be managed outside of this config.
  # When changing this value, make sure to run 'terraform apply' so that it takes effect before the resource is destroyed.
  #preserve_boot_volume = true
  timeouts {
    create = "60m"
  }
}

# Define the volumes that are attached to the compute instances.

resource "oci_core_volume" "test_block_volume" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestBlock0"
  size_in_gbs         = var.size
}

resource "oci_core_volume_attachment" "test_block_attach" {
  attachment_type = "iscsi"
  instance_id     = oci_core_instance.test_instance.*.id[0]
  volume_id       = oci_core_volume.test_block_volume.*.id[0]
  device          = "/dev/oracleoci/oraclevdb"

  # Set this to enable CHAP authentication for an ISCSI volume attachment. The oci_core_volume_attachment resource will
  # contain the CHAP authentication details via the "chap_secret" and "chap_username" attributes.
  use_chap = true

  # Set this to attach the volume as read-only.
  #is_read_only = true
}

resource "null_resource" "remote-exec" {
  depends_on = ["oci_core_instance.test_instance", "oci_core_volume_attachment.test_block_attach"]

  provisioner "remote-exec" {
    connection {
      agent       = false
      timeout     = "30m"
      host        = oci_core_instance.test_instance.*.public_ip[0]
      user        = "opc"
      private_key = file(var.ssh_private_key)
    }

    inline = [
      "touch ~/IMadeAFile.Right.Here",
      "sudo iscsiadm -m node -o new -T ${oci_core_volume_attachment.test_block_attach.*.iqn[0]} -p ${oci_core_volume_attachment.test_block_attach.*.ipv4[0]}:${oci_core_volume_attachment.test_block_attach.*.port[0]}",
      "sudo iscsiadm -m node -o update -T ${oci_core_volume_attachment.test_block_attach.*.iqn[0]} -n node.startup -v automatic",
      "sudo iscsiadm -m node -T ${oci_core_volume_attachment.test_block_attach.*.iqn[0]} -p ${oci_core_volume_attachment.test_block_attach.*.ipv4[0]}:${oci_core_volume_attachment.test_block_attach.*.port[0]} -o update -n node.session.auth.authmethod -v CHAP",
      "sudo iscsiadm -m node -T ${oci_core_volume_attachment.test_block_attach.*.iqn[0]} -p ${oci_core_volume_attachment.test_block_attach.*.ipv4[0]}:${oci_core_volume_attachment.test_block_attach.*.port[0]} -o update -n node.session.auth.username -v ${oci_core_volume_attachment.test_block_attach.*.chap_username[0]}",
      "sudo iscsiadm -m node -T ${oci_core_volume_attachment.test_block_attach.*.iqn[0]} -p ${oci_core_volume_attachment.test_block_attach.*.ipv4[0]}:${oci_core_volume_attachment.test_block_attach.*.port[0]} -o update -n node.session.auth.password -v ${oci_core_volume_attachment.test_block_attach.*.chap_secret[0]}",
      "sudo iscsiadm -m node -T ${oci_core_volume_attachment.test_block_attach.*.iqn[0]} -p ${oci_core_volume_attachment.test_block_attach.*.ipv4[0]}:${oci_core_volume_attachment.test_block_attach.*.port[0]} -l",
      "sudo iscsiadm -m node -T ${oci_core_volume_attachment.test_block_attach.*.iqn[0]} -p ${oci_core_volume_attachment.test_block_attach.*.ipv4[0]}:${oci_core_volume_attachment.test_block_attach.*.port[0]} -l",

      # if you are using the Oracle Linux 7.7, Ubuntu 16.04 or CentOS 7, please use the following commands.
      # If you are using windows, please comment it out.
      "sudo dd iflag=direct if=/dev/sdb of=/dev/null count=1",

      "echo '1' | sudo tee /sys/class/block/sdb/device/rescan",
    ]
  }

  triggers = {
    always_run = oci_core_volume.test_block_volume.size_in_gbs
  }
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
  dns_label      = "testvcn"
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInternetGateway"
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_default_route_table" "default_route_table" {
  manage_default_resource_id = oci_core_vcn.test_vcn.default_route_table_id
  display_name               = "DefaultRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "TestSubnet"
  dns_label           = "testsubnet"
  security_list_ids   = [oci_core_vcn.test_vcn.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  route_table_id      = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id     = oci_core_vcn.test_vcn.default_dhcp_options_id
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}
