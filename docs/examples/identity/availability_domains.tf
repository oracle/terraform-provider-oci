/*
 * This example file demonstrates how to read AD values from a region and employ filters
 * to isolate specific ADs.
 */

data "oci_identity_availability_domains" "ads" {
  compartment_id = "${var.tenancy_ocid}"

  filter {
    name   = "name"
    values = ["\\w*-AD-1"]
    regex  = true
  }
}

output "ads" {
  value = "${data.oci_identity_availability_domains.ads.availability_domains}"
}
