/*
 * This example file shows how to list available OCI regions.
 */

data "oci_identity_regions" "regions" {
}

output "regions" {
  value = "${data.oci_identity_regions.regions.regions}"
}