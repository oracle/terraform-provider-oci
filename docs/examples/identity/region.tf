/*
 * This example file shows how to list available OCI regions, as well as region subscriptions of your tenancy
 */

data "oci_identity_regions" "regions" {}

output "regions" {
  value = "${data.oci_identity_regions.regions.regions}"
}

data "oci_identity_region_subscriptions" "test_region_subscriptions" {
  #Required
  tenancy_id = "${var.tenancy_ocid}"

  filter {
    name   = "is_home_region"
    values = [true]
  }
}

output "region_subscriptions" {
  value = "${data.oci_identity_region_subscriptions.test_region_subscriptions.region_subscriptions}"
}
