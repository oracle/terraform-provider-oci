variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}
variable "ssh_public_key" {}

variable "SecondaryVnicCount" {
    default = 1
}

# Choose an Availability Domain
variable "AD" {
    default = "1"
}

variable "InstanceShape" {
    default = "VM.Standard1.1"
}

variable "instance_image_ocid" {
  type = "map"
  default = {
    // See https://docs.us-phoenix-1.oraclecloud.com/images/
    // Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
    uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
  }
}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_virtual_network" "ExampleVCN" {
  cidr_block = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name = "CompleteVCN"
  dns_label = "examplevcn"
}

resource "oci_core_subnet" "ExampleSubnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  cidr_block = "10.0.1.0/24"
  display_name = "TFExampleSubnet"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.ExampleVCN.id}"
  route_table_id = "${oci_core_virtual_network.ExampleVCN.default_route_table_id}"
  security_list_ids = ["${oci_core_virtual_network.ExampleVCN.default_security_list_id}"]
  dhcp_options_id = "${oci_core_virtual_network.ExampleVCN.default_dhcp_options_id}"
  dns_label = "examplesubnet"
}

resource "oci_core_instance" "ExampleInstance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "TFExampleInstance"
  image = "${var.instance_image_ocid[var.region]}"
  shape = "${var.InstanceShape}"
  subnet_id = "${oci_core_subnet.ExampleSubnet.id}"
  create_vnic_details {
    subnet_id = "${oci_core_subnet.ExampleSubnet.id}"
    hostname_label = "exampleinstance"
  }
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "60m"
  }
}

resource "oci_core_vnic_attachment" "SecondaryVnicAttachment" {
  instance_id = "${oci_core_instance.ExampleInstance.id}"
  display_name = "SecondaryVnicAttachment_${count.index}"
  create_vnic_details {
    subnet_id = "${oci_core_subnet.ExampleSubnet.id}"
    display_name = "SecondaryVnic_${count.index}"
    assign_public_ip = true
    skip_source_dest_check = true
  }
  count = "${var.SecondaryVnicCount}"
}

data "oci_core_vnic" "SecondaryVnic" {
  count = "${var.SecondaryVnicCount}"
  vnic_id = "${element(oci_core_vnic_attachment.SecondaryVnicAttachment.*.vnic_id, count.index)}"
}

output "PrimaryIPAddresses" {
  value = ["${oci_core_instance.ExampleInstance.public_ip}",
           "${oci_core_instance.ExampleInstance.private_ip}"]
}

output "SecondaryPublicIPAddresses" {
  value = ["${data.oci_core_vnic.SecondaryVnic.*.public_ip_address}"]
}

output "SecondaryPrivateIPAddresses" {
  value = ["${data.oci_core_vnic.SecondaryVnic.*.private_ip_address}"]
}
