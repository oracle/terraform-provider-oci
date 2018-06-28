variable "ssh_public_key_path" {}
variable "ssh_private_key_path" {}

variable "compartment_id" {}
variable "src_availability_domain" {}
variable "dst_availability_domain" {}

variable "src_instance_shape" {
  default = "VM.Standard2.4"
}

variable "dst_instance_shape" {
  default = "VM.Standard2.4"
}

variable "src_instance_hostname" {}
variable "dst_instance_hostname" {}

variable "src_instance_image_id" {}
variable "dst_instance_image_id" {}

variable "src_subnet_id" {}
variable "dst_subnet_id" {}

variable "src_export_path" {}
variable "src_mount_target_private_ip" {}

variable "dst_export_path" {}
variable "dst_mount_target_private_ip" {}

variable "mount_point_path" {
  default = "/mnt"
}

variable "snapshot_frequency" {
  default = "@hourly"
}

variable "data_sync_frequency" {
  default = "*/10 * * * *"
}
