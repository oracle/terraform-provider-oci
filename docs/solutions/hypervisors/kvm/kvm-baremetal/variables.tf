#Variables declaration file

#OCI Environment variables
variable "tenancy_ocid" {}

variable "user_ocid" {}
variable "fingerprint" {}
variable "compartment_ocid" {}
variable "region" {}
variable "ssh_public_key_path" {}
variable "ssh_private_key_path" {}

#Template used variables
variable "prefix" {}
variable "availability_domain" {}
variable "vcn_cidr_block" {}
variable "kvm_host_subnet_cidr_block" {}
variable "instance_shape" {}
variable "kvm_image_url" {}
variable "kvm_image_name" {}
variable "kvm_image_path" {}
variable "kvm_guest_domain_name" {}
variable "kvm_guest_memory" {}
variable "kvm_guest_vcpu" {}
variable "kvm_emulation_mode" {}
variable "kvm_guest_os_type" {}
variable "kvm_guest_vnc_port" {}
variable "kvm_guest_vnc_pwd" {}
