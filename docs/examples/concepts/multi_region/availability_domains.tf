/*
 * This example file demonstrates how to read AD values from multiple regions and employ filters
 * to isolate specific ADs.
 */

data "oci_identity_availability_domains" "ad-phx" {
  compartment_id = "${var.tenancy_ocid}"

  filter {
    name   = "name"
    values = ["\\w*-AD-1"]
    regex  = true
  }
}

data "oci_identity_availability_domains" "ad-iad" {
  provider       = "oci.iad"
  compartment_id = "${var.tenancy_ocid}"

  filter {
    name   = "name"
    values = ["\\w*-AD-1"]
    regex  = true
  }
}

output "ad-phx" {
  value = "${data.oci_identity_availability_domains.ad-phx.availability_domains}"
}

output "ad-iad" {
  value = "${data.oci_identity_availability_domains.ad-iad.availability_domains}"
}
