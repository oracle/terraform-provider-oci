variable "tenancy_ocid" {
  type        = string
  description = "OCID of the OCI tenancy."
}

variable "user_ocid" {
  type        = string
  description = "OCID of the OCI user used to authenticate Terraform."
}

variable "fingerprint" {
  type        = string
  description = "Fingerprint of the OCI API signing key associated with the user."
}

variable "private_key_path" {
  type        = string
  description = "Path to the private API signing key file used for OCI authentication."
}

variable "private_key_password" {
  type        = string
  description = "Optional passphrase for the private API signing key. Set to null if the key is not encrypted."
  default     = null
}

variable "region" {
  type        = string
  description = "OCI region identifier (for example: eu-frankfurt-1, us-ashburn-1)."
}

variable "compartment_id" {
  description = "OCID of the compartment where the private endpoint will be created"
  type        = string
}

variable "display_name" {
  description = "Display name for the Distributed Database Private Endpoint"
  type        = string
}

variable "subnet_id" {
  description = "OCID of the subnet used for the Distributed Database Private Endpoint"
  type        = string
}

variable "description" {
  description = "Optional description for the Distributed Database Private Endpoint"
  type        = string
  default     = null
}

variable "nsg_ids" {
  description = "Optional NSG OCIDs to attach to the Distributed Database Private Endpoint"
  type        = list(string)
  default     = []
}

variable "defined_tags" {
  description = "Defined tags to apply to the Distributed Database Private Endpoint"
  type        = map(string)
  default     = {}
}

variable "freeform_tags" {
  description = "Freeform tags to apply to the Distributed Database Private Endpoint"
  type        = map(string)
  default     = {}
}

variable "reinstate_proxy_instance_trigger" {
  type        = number
  description = "Optional trigger to run Reinstate Proxy Instance action. Increment to invoke action."
  default     = null
}
