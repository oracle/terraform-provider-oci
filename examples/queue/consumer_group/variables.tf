variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "queue_display_name" {
  default = "displayName"
}

variable "queue_display_name2" {
  default = "displayName2"
}

variable "queue_display_name3" {
  default = "displayName3"
}

variable "cg_display_name" {
  default = "CGdisplayName"
}

variable "cg_display_name2" {
  default = "CGdisplayName2"
}

variable "is_primary_consumer_group_enabled" {
  default = true
}

variable "primary_consumer_group_display_name" {
  default = "primaryConsumerGroupDisplayName"
}

variable "primary_consumer_group_dead_letter_queue_delivery_count" {
  default = 10
}

variable "primary_consumer_group_filter" {
  default = ""
}