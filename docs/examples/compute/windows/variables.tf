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
    # Image OCIDs for Windows-Server-2012-R2-Standard-Edition-VM-2018.07.19-0 - https://docs.cloud.oracle.com/iaas/images/image/256a6d7c-4fc0-47c7-a61c-b6bbf25c8aba/
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa3nh5j3l3ip62tb2z4gkcrfn23yhwucui5do5abrk4ttvwclxu7ja"

    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaauybkm3gymmgenl3e7eqjcjuh324hbocftnxhyn5o2ghpy4st6xza"
    uk-london-1  = "ocid1.image.oc1.uk-london-1.aaaaaaaad3q2sx2ngclnggc4nvpop6szposwxnljvuswhimiszwcsltsvi2q"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaqcty27fjcgov3a2hl6bimama5l3isv2ejs7utulnpfw5btyyb7gq"
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
