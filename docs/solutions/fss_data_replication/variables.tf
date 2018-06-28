variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region_iad" {}
variable "region_phx" {}

variable "compartment_id" {}

variable "ssh_public_key_path" {}
variable "ssh_private_key_path" {}

variable "instance_image_id" {
  type = "map"

  default = {
    // Oracle_provided image "Oracle_Linux_7.4_2017.12.18_0"
    // See https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/OracleProvidedImageOCIDs.pdf
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaasc56hnpnx7swoyd2fw5gyvbn3kcdmqc2guiiuvnztl2erth62xnq"

    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaaxrqeombwty6jyqgk3fraczdd63bv66xgfsqka4ktr7c57awr3p5a"
    eu_frankfurt_1 = "ocid1.image.oc1.eu_frankfurt_1.aaaaaaaayxmzu6n5hsntq4wlffpb4h6qh6z3uskpbm5v3v4egqlqvwicfbyq"
  }
}

variable "my_vcn_cidr_iad" {
  default = "10.0.0.0/16"
}

variable "my_vcn_cidr_phx" {
  default = "10.1.0.0/16"
}

variable "subnet_cidr_iad_ad1" {
  default = "10.0.1.0/24"
}

variable "subnet_cidr_iad_ad2" {
  default = "10.0.2.0/24"
}

variable "subnet_cidr_phx_ad1" {
  default = "10.1.1.0/24"
}

variable "src_file_system" {
  default = "src_file_system_demo"
}

variable "dst_file_system" {
  default = "dst_file_system_demo"
}

variable "src_mount_target" {
  default = "src_mount_target_demo"
}

variable "dst_mount_target" {
  default = "dst_mount_target_demo"
}

variable "src_export_path" {
  default = "/src_fs_demo"
}

variable "dst_export_path" {
  default = "/dst_fs_demo"
}

variable "max_byte" {
  default = 23843202333
}

variable "max_files" {
  default = 223442
}
