variable "compartment_id" {}
variable "dns_label" {}
variable "vcn_id" {}
variable "route_table_id" {}
variable "dhcp_options_id" {}

variable "security_list_id" {
  type = "list"
}

variable "label_prefix" {}

variable "cidr_block" {
  type = "list"
}

variable "ad_count" {}
variable "tenancy_id" {}

variable "ad_deployment" {
  default = "0"
}

variable "private" {
  default = "false"
}

variable "additional_security_lists_ids" {
  type    = "list"
  default = []
}
