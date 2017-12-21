variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}

variable "AD" {
  default = "1"
}

variable "InstanceShape" {
  default = "VM.Standard1.2"
}

variable "InstanceOS" {
  default = "Windows"
}

variable "InstanceOSVersion" {
  default = "Server 2012 R2 Standard"
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
  cidr_block = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name = "TFExampleVCN"
  dns_label = "tfexamplevcn"
}

resource "oci_core_subnet" "ExampleSubnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  cidr_block = "10.1.20.0/24"
  display_name = "TFExampleSubnet"
  dns_label = "tfexamplesubnet"
  security_list_ids = ["${oci_core_virtual_network.ExampleVCN.default_security_list_id}"]
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.ExampleVCN.id}"
  route_table_id = "${oci_core_virtual_network.ExampleVCN.default_route_table_id}"
  dhcp_options_id = "${oci_core_virtual_network.ExampleVCN.default_dhcp_options_id}"
}

data "oci_core_images" "ImageOCID" {
  compartment_id = "${var.compartment_ocid}"
  operating_system = "${var.InstanceOS}"
  operating_system_version = "${var.InstanceOSVersion}"
}

data "oci_core_instance_credentials" "InstanceCredentials" {
  instance_id = "${oci_core_instance.TFInstance.id}"
}

resource "oci_core_instance" "TFInstance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "TFWindows"
  image = "${lookup(data.oci_core_images.ImageOCID.images[0], "id")}"
  shape = "${var.InstanceShape}"
  subnet_id = "${oci_core_subnet.ExampleSubnet.id}"
  hostname_label = "winmachine"
  metadata {}
}

output "Username" {
  value = ["${data.oci_core_instance_credentials.InstanceCredentials.username}"]
}

output "Password" {
  value = ["${data.oci_core_instance_credentials.InstanceCredentials.password}"]
}

output "InstancePublicIP" {
  value = ["${oci_core_instance.TFInstance.public_ip}"]
}

output "InstancePrivateIP" {
  value = ["${oci_core_instance.TFInstance.private_ip}"]
}
