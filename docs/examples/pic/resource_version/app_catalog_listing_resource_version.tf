data "oci_core_app_catalog_listings" "test_app_catalog_listings" {
  filter {
    name   = "publisher_name"
    values = ["Oracle CCE Image Management Pipeline"]
  }
}

data "oci_core_app_catalog_listing_resource_versions" "test_app_catalog_listing_resource_versions" {
  #Required
  listing_id = "${lookup(data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings[0],"listing_id")}"
}

data "oci_core_app_catalog_listing_resource_version" "test_app_catalog_listing_resource_version" {
  #Required
  listing_id               = "${lookup(data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings[0],"listing_id")}"
  listing_resource_version = "${lookup(data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0],"listing_resource_version")}"
}

output "resource_versions" {
  value = ["${data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions}"]
}

output "single_resource_version" {
  value = ["${data.oci_core_app_catalog_listing_resource_version.test_app_catalog_listing_resource_version.listing_resource_version}"]
}

output "compatible_shapes" {
  value = ["${data.oci_core_app_catalog_listing_resource_version.test_app_catalog_listing_resource_version.compatible_shapes}"]
}
