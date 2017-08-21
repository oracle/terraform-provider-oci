variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}
variable "ssh_public_key" {}

variable "SecondaryVnicCount" {
    default = 2
}

# Choose an Availability Domain
variable "AD" {
    default = "1"
}

variable "InstanceShape" {
    default = "VM.Standard1.8"
}

variable "InstanceOS" {
    default = "Oracle Linux"
}

variable "InstanceOSVersion" {
    default = "7.3"
}

provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

data "baremetal_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "baremetal_core_virtual_network" "ExampleVCN" {
  cidr_block = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name = "CompleteVCN"
  dns_label = "examplevcn"
}

resource "baremetal_core_subnet" "ExampleSubnet" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  cidr_block = "10.0.1.0/24"
  display_name = "ExampleSubnet"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.ExampleVCN.id}"
  route_table_id = "${baremetal_core_virtual_network.ExampleVCN.default_route_table_id}"
  security_list_ids = ["${baremetal_core_virtual_network.ExampleVCN.default_security_list_id}"]
  dhcp_options_id = "${baremetal_core_virtual_network.ExampleVCN.default_dhcp_options_id}"
  dns_label = "examplesubnet"
}

# Gets the OCID of the OS image to use
data "baremetal_core_images" "OLImageOCID" {
    compartment_id = "${var.compartment_ocid}"
    operating_system = "${var.InstanceOS}"
    operating_system_version = "${var.InstanceOSVersion}"
}

resource "baremetal_core_instance" "ExampleInstance" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "ExampleInstance"
  image = "${lookup(data.baremetal_core_images.OLImageOCID.images[0], "id")}"
  shape = "${var.InstanceShape}"
  subnet_id = "${baremetal_core_subnet.ExampleSubnet.id}"
  create_vnic_details {
    subnet_id = "${baremetal_core_subnet.ExampleSubnet.id}"
    hostname_label = "exampleinstance"
  }
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "60m"
  }
}

resource "baremetal_core_vnic_attachment" "PrivateVnicAttachment" {
  instance_id = "${baremetal_core_instance.ExampleInstance.id}"
  create_vnic_details {
    subnet_id = "${baremetal_core_subnet.ExampleSubnet.id}"
    display_name = "PrivateVnic"
    assign_public_ip = false
  }
}

resource "baremetal_core_vnic_attachment" "SecondaryVnicAttachment" {
  instance_id = "${baremetal_core_instance.ExampleInstance.id}"
  display_name = "SecondaryVnicAttachment_${count.index}"
  create_vnic_details {
    subnet_id = "${baremetal_core_subnet.ExampleSubnet.id}"
    display_name = "SecondaryVnic_${count.index}"
    assign_public_ip = true
  }
  count = "${var.SecondaryVnicCount}"
}

data "baremetal_core_vnic" "SecondaryVnic" {
  count = "${var.SecondaryVnicCount}"
  vnic_id = "${element(baremetal_core_vnic_attachment.SecondaryVnicAttachment.*.vnic_id, count.index)}"
}

output "PrimaryIPAddresses" {
  value = ["${baremetal_core_instance.ExampleInstance.public_ip}",
           "${baremetal_core_instance.ExampleInstance.private_ip}"]
}

output "SecondaryPublicIPAddresses" {
  value = ["${data.baremetal_core_vnic.SecondaryVnic.*.public_ip_address}"]
}

output "SecondaryPrivateIPAddresses" {
  value = ["${data.baremetal_core_vnic.SecondaryVnic.*.private_ip_address}"]
}

data "baremetal_core_vnic" "PrivateVnic" {
  vnic_id = "${baremetal_core_vnic_attachment.PrivateVnicAttachment.vnic_id}"
}

output "PrivateVnicIPAddresses" {
  value = ["${data.baremetal_core_vnic.PrivateVnic.private_ip_address}", "${data.baremetal_core_vnic.PrivateVnic.public_ip_address}"]
}