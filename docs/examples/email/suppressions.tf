variable "suppression_email_address" {
  default = "JohnSmith@example.com"
}

variable "suppression_time_created_greater_than_or_equal_to" {
  default = "2018-01-01T00:00:00.000Z"
}

variable "suppression_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}

resource "oci_email_suppression" "test_suppression" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  email_address  = "${var.suppression_email_address}"
}

data "oci_email_suppressions" "test_suppressions" {
  #Required
  compartment_id = "${var.tenancy_ocid}"

  #Optional
  email_address                         = "${var.suppression_email_address}"
  time_created_greater_than_or_equal_to = "${var.suppression_time_created_greater_than_or_equal_to}"
  time_created_less_than                = "${var.suppression_time_created_less_than}"

  filter {
    name   = "id"
    values = ["${oci_email_suppression.test_suppression.id}"]
  }
}

output "emailSuppressions" {
  value = "${data.oci_email_suppressions.test_suppressions.suppressions}"
}
