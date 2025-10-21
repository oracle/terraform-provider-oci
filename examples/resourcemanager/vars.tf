variable "compartment_ocid" {}
variable "tenancy_ocid" {}
// Private subnet the compute instance will reside in
variable "subnet_ocid" {}
// Private endpoint that has already been created.
variable "orm_private_endpoint_ocid" {
  type = string
}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "auth" {}
variable "config_file_profile" {}