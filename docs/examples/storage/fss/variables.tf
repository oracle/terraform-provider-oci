variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}

# Refer https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingkeypairs.htm on how to setup SSH key pairs for compute instances
variable "ssh_public_key" {}

variable "ssh_private_key" {}

variable "api_public_key" {
  default = "-----BEGIN PUBLIC KEY-----MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4fGHcxbEs3VaWoKaGUiPHGZ5ILiOXCcWN4nOgLr6CSzUjtgjmN3aA6rsT2mYiD+M5EecDbEUMectUhNtLl5LPABN9kpjuR0zxCJXvYYQiCBtdjb1/YxrZI9T/9Jtd+cTabCahJHR/cR8jFmvO4cKJCa/0+Y00zvktrqniHIn3edGAKC4Ttlwj/1NqT0ZVePMXg3rWHPsIW6ONfdn6FNfMet8Qa8K3C9xVvzImlYx8PQBy/44Ilu5T3A+puwb2QMeZnQZGDALOY4MvrBTTA1TdjFpg1NChj2rGYzreysqlnKFu+1qg64wel39kHkppz4Fv2vaLXF9qIeDjeo3G4sHQIDAQAB-----END PUBLIC KEY-----"
}

# Choose an Availability Domain
variable "availability_domain" {
  default = "3"
}

variable "my_vcn-cidr" {
  default = "10.0.0.0/16"
}

variable "my_subnet_cidr" {
  default = "10.0.1.0/24"
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

variable "export_read_write_access_source" {
  default = "10.0.0.0/8"
}

variable "export_read_only_access_source" {
  default = "0.0.0.0/0"
}

variable "instance_image_ocid" {
  type = "map"

  default = {
    // See https://docs.us-phoenix-1.oraclecloud.com/images/
    // Oracle-provided image "Oracle-Linux-7.5-2018.05.09-1"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaazregkysspxnktw35k4r5vzwurxk6myu44umqthjeakbkvxvxdlkq"

    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaa6ybn2lkqp2ejhijhehf5i65spqh3igt53iyvncyjmo7uhm5235ca"
    uk-london-1  = "ocid1.image.oc1.uk-london-1.aaaaaaaayodsld656eh5stds5mo4hrmwuhk2ugin4eyfpgoiiskqfxll6a4a"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaozjbzisykoybkppaiwviyfzusjzokq7jzwxi7nvwdiopk7ligoia"
  }
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

locals {
  mount_target_1_ip_address = "${lookup(data.oci_core_private_ips.ip_mount_target1.private_ips[0], "ip_address")}"
}
