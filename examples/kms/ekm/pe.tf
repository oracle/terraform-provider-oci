variable "ekms_subnet_id" {
}

variable "displayName" {
  default = "TF-PE"
}

variable "externalKeyManagerIp" {
  default = "10.0.0.31"
}

variable "ekms_ca_bundle" {
}

resource "oci_kms_ekms_private_endpoint" "tf_example_pe" {
  subnet_id = var.ekms_subnet_id
  compartment_id = var.compartment_ocid
  display_name = var.displayName
  external_key_manager_ip = var.externalKeyManagerIp
  ca_bundle = var.ekms_ca_bundle
}