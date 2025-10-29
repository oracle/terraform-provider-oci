variable "tenancy_ocid" {
  description = "The OCID of your OCI Tenancy. This value should start with something similar to 'ocid1.tenancy...'."
}

variable "user_ocid" {
  description = "The OCID of the OCI User that will execute this script. This value should start with something similar to 'ocid1.user...'."
}

variable "secret_ocid" {
  description = "The OCID of an OCI Vault Secret that contains the password that will be used by the user configured in the database tools connection."
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

variable "connection_ocid" {
  description = "The OCID of the Database Tools connection that this identity is for."
}

variable "connection_string" {
  type = string
  description = "A string specifying how to connect to the database."
}

variable "user_password_secret_ocid" {
  type = string
  description = "The OCID of a Vault secret containing the password of the database user specified in the connection."
}

variable "database_wallet_secret_ocid" {
  type = string
  description = "The OCID of a Vault secret containing the SSO wallet of the database to connect to."
}