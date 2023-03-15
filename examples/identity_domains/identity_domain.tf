#provide the id of the identity domain to work on
variable "identity_domain_id" {
  default = ""
}

data "oci_identity_domain" "test_domain" {
  domain_id = var.identity_domain_id
}
