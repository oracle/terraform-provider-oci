# Creating the Bridge Instance
resource "oci_core_instance" "BridgeInstance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "BridgeInstance"
  image               = "${var.InstanceImageOCID[var.region]}"
  shape               = "${var.InstanceShape}"

  create_vnic_details {
    subnet_id              = "${oci_core_subnet.MgmtSubnet.id}"
    skip_source_dest_check = true
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data           = "${base64encode(file("user_data.tpl"))}"
  }

  timeouts {
    create = "10m"
  }
}

# Gets a list of VNIC attachments on the instance
data "oci_core_vnic_attachments" "BridgeInstanceVnics" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id         = "${oci_core_instance.BridgeInstance.id}"
}

# Create PrivateIP
resource "oci_core_private_ip" "BridgeInstancePrivateIP" {
  vnic_id      = "${lookup(data.oci_core_vnic_attachments.BridgeInstanceVnics.vnic_attachments[0],"vnic_id")}"
  display_name = "BridgeInstancePrivateIP"
}

# Get the OCID of the first (default) VNIC
data "oci_core_vnic" "BridgeInstanceVnic1" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.BridgeInstanceVnics.vnic_attachments[0],"vnic_id")}"
}

# Creating secondary VNIC on BridgeInstance and attaching it to Second VCN Mgmt subnet
resource "oci_core_vnic_attachment" "SecondaryVnicAttachment" {
  create_vnic_details {
    subnet_id              = "${oci_core_subnet.MgmtSubnet2.id}"
    display_name           = "SecondaryVNIC"
    skip_source_dest_check = true
  }

  instance_id = "${oci_core_instance.BridgeInstance.id}"
}

# Gets a list of VNIC attachments on the instance
data "oci_core_vnic_attachments" "BridgeInstanceVnics2" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id         = "${oci_core_instance.BridgeInstance.id}"
}

# Gets the OCID of the second VNIC
data "oci_core_vnic" "BridgeInstanceVnic2" {
  vnic_id = "${oci_core_vnic_attachment.SecondaryVnicAttachment.vnic_id}"
}

# Gets a list of private IPs on the second VNIC
data "oci_core_private_ips" "BridgeInstancePrivateIP2" {
  vnic_id = "${data.oci_core_vnic.BridgeInstanceVnic2.id}"
}

# Configurations for setting up the secondary VNIC
resource "null_resource" "configure-secondary-vnic" {
  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.ssh_private_key}"
    host        = "${data.oci_core_vnic.BridgeInstanceVnic1.public_ip_address}"
    timeout     = "30m"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo wget https://docs.cloud.oracle.com/iaas/Content/Resources/Assets/secondary_vnic_all_configure.sh",
      "sudo chmod 777 secondary_vnic_all_configure.sh",
      "sudo ./secondary_vnic_all_configure.sh -c ${lookup(data.oci_core_private_ips.BridgeInstancePrivateIP2.private_ips[0],"id")}",
      "sudo ip route add ${var.vcn_cidr2} dev ens4 via ${oci_core_subnet.MgmtSubnet2.virtual_router_ip}",
    ]
  }
}
