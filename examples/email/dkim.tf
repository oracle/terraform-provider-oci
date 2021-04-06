variable "dkim_name" {
  default = "test-selector-1"
}

variable "dkim_description" {
  default = "Test DKIM"
}

resource "oci_email_dkim" "testdkim" {
	#Required
	email_domain_id = oci_email_email_domain.test_email_domain.id
	#Optional
	description = var.dkim_description
	name = var.dkim_name
  depends_on = [oci_email_email_domain.test_email_domain]
}
