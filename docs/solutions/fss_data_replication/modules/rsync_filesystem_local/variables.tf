# ------------------------------------------------------------------------------------------------------------------------
# MODULE VARIABLES
# ------------------------------------------------------------------------------------------------------------------------

variable "ssh_public_key_path" {
  description = "The SSH public key path which Terraform will added to SSH authorized_keys on Oracle Cloud Infrastructure Compute Instance."
}

variable "ssh_private_key_path" {
  description = "The SSH private key path which Terraform will added to SSH authorized_keys on Oracle Cloud Infrastructure Compute Instance."
}

variable "compartment_id" {
  description = "The OCID of the Oracle Cloud Infrastructure Compartment where all resources will be created."
}

variable "availability_domain" {
  description = "The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` "
}

variable "subnet_id" {
  description = "The OCID of the Oracle Cloud Infrastructure Subnet where rsync instance will be created."
}

variable "instance_shape" {
  description = "The Oracle Cloud Infrastructure Compute Shape for running rsync. Network Bandwidth available depends on the compute shape. More info https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Concepts/computeoverview.htm"
  default     = "VM.Standard2.4"
}

variable "instance_hostname" {
  description = "Rsync instance hostname."
}

variable "instance_image_id" {
  description = "The OCID of the Oracle Cloud Infrastructure Image. Image OCIDs is specific to region. https://docs.us-phoenix-1.oraclecloud.com/images/"
}

variable "src_export_path" {
  description = "Source Export Path specified in the file system. Export path is appended to the mount target IP address and is used to mount to the file system. Example: `/src_fs_demo-iad-ad1` "
}

variable "src_mount_target_private_ip" {
  description = "IP address of the FSS Mount Target where the data will be copied from. Example: 10.0.0.5 , Mount Target IP and Export path are used to mount the file system. Example: 10.0.0.5:/mnt/src_fs_demo-iad-ad1"
}

variable "dst_export_path" {
  description = "Destination Export Path specified in the file system. Export path is appended to the mount target IP address and is used to mount to the file system. Example: `/dst_fs_demo-iad-ad2` "
}

variable "dst_mount_target_private_ip" {
  description = "IP address of the FSS Mount Target where the data will be copied to. Example: 10.0.1.5 , Mount Target IP and Export path are used to mount the file system. Example: 10.0.1.5:/mnt/dst_fs_demo-iad-ad2"
}

variable "mount_point_path" {
  description = "Path within the rsync instance to a locally accessible directory to which the remote file system is monted. Example: /mtn"
  default     = "/mnt"
}

variable "data_sync_frequency" {
  description = "Data sync job frequency based on Cron. Source: https://en.wikipedia.org/wiki/Cron"
  default     = "*/30 * * * *"
}
