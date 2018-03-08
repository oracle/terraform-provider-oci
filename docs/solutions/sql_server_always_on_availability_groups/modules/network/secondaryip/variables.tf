variable "subnets" {
  type = "list"
}

variable "label_prefix" {}
variable "compartment_id" {}
variable "vcn_id" {}
variable "tenancy_id" {}
variable "dns_label" {}
variable "ad_count" {}

variable "vnic_ids" {
  type = "list"
}
