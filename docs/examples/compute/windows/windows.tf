variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "subnet_ocid" {}

variable "InstanceShape" {
  default = "VM.Standard1.2"
}

variable "InstanceOS" {
  default = "Windows"
}

variable "InstanceOSVersion" {
  default = "Server 2012 R2 Standard"
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

data "baremetal_core_images" "ImageOCID" {
  compartment_id = "${var.compartment_ocid}"
  operating_system = "${var.InstanceOS}"
  operating_system_version = "${var.InstanceOSVersion}"
}

data "baremetal_core_instance_credentials" "InstanceCredentials" {
  instance_id = "${baremetal_core_instance.TFInstance.id}"
}

resource "baremetal_core_instance" "TFInstance" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "TFWindows"
  image = "${lookup(data.baremetal_core_images.ImageOCID.images[0], "id")}"
  shape = "${var.InstanceShape}"
  subnet_id = "${var.subnet_ocid}"
  hostname_label = "winmachine"
  metadata {}
}

output "Username" {
  value = ["${data.baremetal_core_instance_credentials.InstanceCredentials.username}"]
}

output "Password" {
  value = ["${data.baremetal_core_instance_credentials.InstanceCredentials.password}"]
}

output "InstancePublicIP" {
  value = ["${baremetal_core_instance.TFInstance.public_ip}"]
}

output "InstancePrivateIP" {
  value = ["${baremetal_core_instance.TFInstance.private_ip}"]
}
