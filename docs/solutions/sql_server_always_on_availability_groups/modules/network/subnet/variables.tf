variable "compartment_ocid" {}
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
variable "tenancy_ocid" {}

# ad_deployment - index of the AD in which to deploy the subnet
variable "ad_deployment" {
  default = "0"
}

# private - define whether the network is a public or private network
variable "private" {
  default = "false"
}

# additional_security_list_ids - list of additional security lists to include
# in the subnet definition, if needed.
variable "additional_security_lists_ids" {
  type    = "list"
  default = []
}
