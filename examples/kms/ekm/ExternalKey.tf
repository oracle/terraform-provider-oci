
variable "display_name" {
  default = "ext-key-tf-example"
}


variable "key_key_shape_algorithm" {
  default = "AES"
}

variable "key_key_shape_length" {
  default = 32
}

variable "external_key_id" {
  default = "f3cf68ae-659c-4e9e-8be7-ee39fa9ffa3c"
}

resource "oci_kms_key" "test_key" {
  #Required
  compartment_id      = var.compartment_ocid
  display_name        = var.display_name
  management_endpoint = "avsnmg6paahhm-management.kms.r1.oracleiaas.com"
  protection_mode = "EXTERNAL"

  key_shape {
    #Required
    algorithm = var.key_key_shape_algorithm
    length    = var.key_key_shape_length
  }

  external_key_reference {
    external_key_id = var.external_key_id
  }

}