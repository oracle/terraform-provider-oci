resource "random_id" "random_email_domain" {
  byte_length = 8
}

variable "email_domain_description" {
  default = "Test Email Domain"
}

resource "oci_email_email_domain" "test_email_domain" {

	#Required
	compartment_id = var.compartment_ocid
	name = "${random_id.random_email_domain.id}.email.${var.region}.dummydomain.com"

	#Optional
	description = var.email_domain_description
	freeform_tags = {"Department"= "Finance"}

  depends_on = [random_id.random_email_domain]
}
