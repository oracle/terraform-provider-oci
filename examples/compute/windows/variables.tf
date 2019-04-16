// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

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
    # Images released in and after July 2018 have cloudbase-init and winrm enabled by default, refer to the release notes - https://docs.cloud.oracle.com/iaas/images/
    # Image OCIDs for Windows-Server-2012-R2-Standard-Edition-VM-Gen2-2018.10.12-0 - https://docs.cloud.oracle.com/iaas/images/image/80b70ffd-5efc-479e-872c-d1bf6bcbefbd/
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaat5km25plmetj6gtnhrr5xprmv7boe25q2vrzwhbgno5yh2owybja"

    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaxgzzrdoge7zxrjtmjqjhicaxsujljvaju3mbwryo5x5k5axlmsza"
    uk-london-1  = "ocid1.image.oc1.uk-london-1.aaaaaaaaedntd3p6jed5d2p7gsohfu6x3k67s364amtzb5vwfzrvfzt2rrlq"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaskz7sq3mlmiwazehuqzoxdq4xz7sinrwn5m6kedxz3td2c7it2vq"
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
