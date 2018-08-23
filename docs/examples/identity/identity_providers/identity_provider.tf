resource "oci_identity_identity_provider" "test_identity_provider" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  description    = "${var.identity_provider_description}"
  metadata       = "${var.identity_provider_metadata != "" ? var.identity_provider_metadata : "${file("${var.identity_provider_metadata_file}")}"}"
  metadata_url   = "${var.identity_provider_metadata_url}"
  name           = "${var.identity_provider_name}"
  product_type   = "${var.identity_provider_product_type}"
  protocol       = "${var.identity_provider_protocol}"

  #Optional
  freeform_tags = "${var.identity_provider_freeform_tags}"
}

data "oci_identity_identity_providers" "test_identity_providers" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  protocol       = "${var.identity_provider_protocol}"
}

output "identity_providers" {
  value = "${data.oci_identity_identity_providers.test_identity_providers.identity_providers}"
}
