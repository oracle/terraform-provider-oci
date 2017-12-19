#Variables declaration file

#OCI Environment variables
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "compartment_ocid" {}
variable "region" {}
variable "fingerprint" {}
variable "ssh_api_private_key_path" {}
variable "ssh_user_private_key_path" {}
variable "ssh_user_public_key_path" {}


#Prefix to identify your resources
variable "prefix" {
  default = "samplenestedkvm"
}

#availability_domain number For AD1 uses 1. For AD2, uses 2, For AD3, uses 3
variable "availability_domain" {
  default = "1"
}

#Cidr block for your VCN
variable "vcn_cidr_block" {
  default = "10.0.0.0/16"
}

#Cidr block for your subnet
variable "kvm_host_subnet_cidr_block" {
  default = "10.0.10.0/24"
}

#KVM Host instance shape. Only VM Shapes are supported on this example
variable "instance_vm_shape" {
  default = "VM.Standard1.8"
}

#URL of your image file for downloading (you can place your image in the object storage!)
variable "kvm_image_url" {
  default = "https://<my-qcow2-image-url>"
}

#File name of the qcow2 image file stored in the KVM host
variable "kvm_image_name" {
  default = "my-qcow2-image.qcow2"
}

#Location where the qcow2 file will be placed in the KVM host
variable "kvm_image_path" {
  default = "/kvm-imgs/"
}

########################################################################
#### KVM domain settings, modify them accordingly with your needs. #####
########################################################################


variable "kvm_guest_domain_name" {
  default = "MyKVMDomain"
}
variable "kvm_guest_memory" {
  default = "16384"
}
variable "kvm_guest_vcpu" {
  default = "8"
}
variable "kvm_emulation_mode" {
  default = "virtio"
}
variable "kvm_guest_os_type" {
  default = "linux"
}
variable "kvm_guest_vnc_port" {
  default = "5901"
}
variable "kvm_guest_vnc_pwd" {
  default = "Test123"
}
