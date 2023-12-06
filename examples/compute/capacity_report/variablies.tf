variable "region" {
  description = "region in which to call API."
}

variable "compartment_ocid" {
  description = "compartment_id in which to call API."
}

variable "tenancy_ocid" {
  description = "tenancy OCID in which to launch instances."
}

variable "user_ocid" {
  description = "user OCID for the user account with which to connect to the API."
}

variable "private_key_path" {
  description = "full file path of the private key to use for API access with the user account. Does not support environment variables or ~ to abbreviate a user's home directory."
}

variable "fingerprint" {
  description = "PEM fingerprint for private key."
}

variable "compute_capacity_report_availability_domain" {
  description = "AD in which to call API."
  default = "UgLr:PHX-AD-1"
}

variable "compute_capacity_report_shape_availabilities_fault_domain" {
  description = "FD in which to call API."
}

variable "compute_capacity_report_shape_availabilities_instance_shape_flex" {
  default = "VM.Standard.E4.Flex"
}
variable "compute_capacity_report_shape_availabilities_instance_shape_fix" {
  default = "VM.Standard1.8"
}

variable "compute_capacity_report_shape_availabilities_instance_shape_config_memory_in_gbs" {
  default = 16
}

variable "compute_capacity_report_shape_availabilities_instance_shape_config_nvmes" {
  default = 2
}

variable "compute_capacity_report_shape_availabilities_instance_shape_config_ocpus" {
  default = 1
}