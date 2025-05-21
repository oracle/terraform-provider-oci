variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "ssh_public_key" {
}

variable "avm_certificate_id"{

}

variable "cloud_exadata_infrastructure_compute_count" {
  default = "2"
}

variable "cloud_exadata_infrastructure_storage_count" {
  default = "3"
}

variable "cloud_exadata_infrastructure_shape" {
  default = "Exadata.X8M"
}

# schedule time should be at least 7 days after the current date
variable "time_for_schedule_mr" {
  default = "2025-04-30T15:15:15.000Z"
}

variable "maintenance_run_history_id" {
}