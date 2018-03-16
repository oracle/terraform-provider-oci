variable "subnets" {
  type = "list"
}

variable "label_prefix" {}
variable "compartment_ocid" {}
variable "vcn_id" {}
variable "tenancy_ocid" {}
variable "dns_label" {}
variable "ad_count" {}

variable "vnic_ids" {
  type = "list"
}
