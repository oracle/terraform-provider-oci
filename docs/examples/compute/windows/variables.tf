############
# Variables
############
variable "tenancy_ocid" {}

variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}

variable "cloudinit_ps1" {
  default = "cloudinit.ps1"
}

variable "cloudinit_config" {
  default = "cloudinit.yml"
}

variable "setup_ps1" {
  default = "setup.ps1"
}

variable "userdata" {
  default = "userdata"
}

variable "size_in_gbs" {
  default = "256"
}

variable "instance_name" {
  default = "TFWindows"
}

variable "instance_user" {
  default = "opc"
}

variable "availability_domain" {
  default = "2"
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

variable instance_image_ocid {
  type = "map"

  default = {
    # Images released after June 2018 have cloudbase-init and winrm enabled by default, refer to the release notes - https://docs.cloud.oracle.com/iaas/images/
    # The below Image OCIDs are for Windows-Server-2012-R2-Standard-Edition-VM-Gen2-2018.12.12-0
    # See https://docs.cloud.oracle.com/iaas/images/image/5e34cde5-6cef-4cc3-b8f1-c8fc3a088302/
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaarlo3ace3wq34aompwj3u2z2xteonboapg663woz6d2iovarowhja"

    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaabzwak2haqxh3r7h6dajgu4enp7q7hcrreql45awryd5frjsd5l6a"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaourcjktoe3gprvwfksxc36r4rxgbcjs5qvtrja6w6euivci635vq"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaadb4mg7ii73wkrntmiunr7x7qrh7ompczvy3xbggm27pkhotpgj2q"
  }
}

############
# Provider
############
provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}
