variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "api_public_key" {
  default = "-----BEGIN PUBLIC KEY-----MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4fGHcxbEs3VaWoKaGUiPHGZ5ILiOXCcWN4nOgLr6CSzUjtgjmN3aA6rsT2mYiD+M5EecDbEUMectUhNtLl5LPABN9kpjuR0zxCJXvYYQiCBtdjb1/YxrZI9T/9Jtd+cTabCahJHR/cR8jFmvO4cKJCa/0+Y00zvktrqniHIn3edGAKC4Ttlwj/1NqT0ZVePMXg3rWHPsIW6ONfdn6FNfMet8Qa8K3C9xVvzImlYx8PQBy/44Ilu5T3A+puwb2QMeZnQZGDALOY4MvrBTTA1TdjFpg1NChj2rGYzreysqlnKFu+1qg64wel39kHkppz4Fv2vaLXF9qIeDjeo3G4sHQIDAQAB-----END PUBLIC KEY-----"
}

variable "my_vcn-cidr" {
  default = "10.0.0.0/16"
}

variable "my_subnet_cidr" {
  default = "10.0.1.0/24"
}

variable "availability_domain" {
  default = "kIdk:PHX-AD-1"
}

variable "file_system_1_display_name" {
  default = "my_fs_1"
}

variable "file_system_2_display_name" {
  default = "my_fs_2"
}

variable "mount_target_1_display_name" {
  default = "my_mount_target_1"
}

variable "mount_target_2_display_name" {
  default = "my_mount_target_2"
}

variable "export_path_fs1_mt1" {
  default = "/myfsspaths/fs1/path1"
}

variable "export_path_fs1_mt2" {
  default = "/myfsspaths/fs1/path2"
}

variable "export_path_fs2_mt1" {
  default = "/myfsspaths/fs2/path1"
}

variable "snapshot_name" {
  default = "20180320_daily"
}

variable "export_set_name_1" {
  default = "export set for mount target 1"
}

variable "export_set_name_2" {
  default = "export set for mount target 2"
}

variable "max_byte" {
  default = 23843202333
}

variable "max_files" {
  default = 223442
}

variable "instance_image_ocid" {
  type = "map"
  default = {
    // Oracle-provided image "Oracle-Linux-7.4-2017.12.18-0"
    // See https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/OracleProvidedImageOCIDs.pdf
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaasc56hnpnx7swoyd2fw5gyvbn3kcdmqc2guiiuvnztl2erth62xnq"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaxrqeombwty6jyqgk3fraczdd63bv66xgfsqka4ktr7c57awr3p5a"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaayxmzu6n5hsntq4wlffpb4h6qh6z3uskpbm5v3v4egqlqvwicfbyq"
  }
}

variable "instance_shape" {
  default = "VM.Standard1.2"
}

locals {
  mount_target_1_ip_address = "${lookup(data.oci_core_private_ips.ip_mount_target1.private_ips[0], "ip_address")}"
}