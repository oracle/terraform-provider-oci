variable "tenancy_ocid" {
  description = "The OCID of your OCI tenancy. This value should start with something similar to 'ocid1.tenancy...'."
}

variable "user_ocid" {
  description = "The OCID of the OCI user that will execute this script. This value should start with something similar to 'ocid1.user...'."
}

variable "secret_ocid" {
  description = "The OCID of an OCI Vault secret that contains the password that will be used by the user configured in the database tools connection."
}

variable "fingerprint" {
  description = "Fingerprint for the key pair being used."
}

variable "private_key_path" {
  description = "The path (including filename) of the private key stored on the computer being used to access OCI."
}

variable "compartment_ocid" {
  description = "The OCID of the compartment you want to create resources in."
}

variable "region" {
  description = "The region that you want to create resources in."
}

variable "ca_certificate_pem_secret_ocid" {
  description = "An OCI Vault secret containing the CA Certificate PEM file for the database."
}

variable "client_private_key_pem_secret_ocid" {
  description = "An OCI Vault secret containing the Client Private Key PEM file for the database."
}

variable "client_private_key_pem_password_secret_ocid" {
  description = "An OCI Vault secret containing the password for the Client Private Key PEM file contained in var.client_private_key_pem_secret_ocid"
}

variable "client_certificate_pem_secret_ocid" {
  description = "An OCI Vault secret containing the Client Certificate PEM file for the database."
}