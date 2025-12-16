## Required connection credentials ##
variable "user_ocid" {
  description = "user OCID for the user account with which to connect to the API."
}

variable "private_key_path" {
  description = "full file path of the private key to use for API access with the user account. Does not support environment variables or ~ to abbreviate a user's home directory."
}

variable "fingerprint" {
  description = "PEM fingerprint for private key."
}

variable "tenancy_ocid" {
  description = "tenancy OCID in which to launch instances."
}

variable "compartment_ocid" {
  description = "compartment OCID in which to launch instances."
}

## Optional parameters ##
variable "instance_count" {
  description = "number of instances to launch."
  default = "1"
}

variable "instance_name_prefix" {
  default = "capacity_reservation_"
}

variable "region" {
  description = "region in which to launch instances."
  default = "us-phoenix-1"
}

# Dynamic image lookup filters
variable "image_operating_system" {
  description = "Operating system to filter images."
  default     = "Oracle Linux"
}

variable "image_operating_system_version" {
  description = "Operating system version to filter images (e.g., 8 or 8.10)."
  default     = "8"
}

variable "instance_shape" {
  default = "VM.Standard.E5.Flex"
}
