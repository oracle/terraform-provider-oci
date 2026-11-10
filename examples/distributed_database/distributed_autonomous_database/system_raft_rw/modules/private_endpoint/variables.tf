variable "compartment_id" {
  type = string
}

variable "display_name" {
  type = string
}

variable "subnet_id" {
  type = string
}

variable "vcn_id" {
  type = string
}

variable "description" {
  type    = string
  default = null
}

variable "nsg_ids" {
  type    = list(string)
  default = []
}

variable "freeform_tags" {
  type    = map(string)
  default = {}
}
