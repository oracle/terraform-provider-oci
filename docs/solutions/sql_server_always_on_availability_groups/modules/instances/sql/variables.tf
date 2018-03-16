variable "subnets" {
  type = "list"
}

variable "label_prefix" {}
variable "compartment_ocid" {}
variable "vcn_id" {}
variable "tenancy_ocid" {}
variable "dns_label" {}
variable "image_id" {}
variable "shape" {}

variable "db_volumes" {
  type = "list"
}

variable "log_volumes" {
  type = "list"
}

variable "backup_volumes" {
  type = "list"
}

variable "ad_count" {}
