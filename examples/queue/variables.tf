variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "queue_custom_encryption_key_id" {}

variable "tag_namespace_description" {
  default = "Just a test"
}

variable "queue_defined_tags_value" {
  default = "value"
}

variable "tag_namespace_name" {
  default = "testexamples-tag-namespace"
}

variable "queue_dead_letter_queue_delivery_count" {
  default = 10
}

variable "queue_display_name" {
  default = "displayName"
}

variable "queue_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "queue_retention_in_seconds" {
  default = 10
}

variable "queue_state" {
  default = "ACTIVE"
}

variable "queue_timeout_in_seconds" {
  default = 10
}

variable "queue_visibility_in_seconds" {
  default = 10
}

variable "queue_channel_consumption_limit" {
  default = 10
}

variable "purge_type" {
  default = "NORMAL"
}

variable "purge_trigger" {
  default = 1
}