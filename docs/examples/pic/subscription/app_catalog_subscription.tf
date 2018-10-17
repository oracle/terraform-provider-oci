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

resource "oci_core_app_catalog_listing_resource_version_agreement" "test_app_catalog_listing_resource_version_agreement" {
  #Required
  listing_id               = "${lookup(data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0], "listing_id")}"
  listing_resource_version = "${lookup(data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0], "listing_resource_version")}"
}

resource "oci_core_app_catalog_subscription" "test_app_catalog_subscription" {
  compartment_id           = "${var.compartment_ocid}"
  eula_link                = "${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.eula_link}"
  listing_id               = "${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.listing_id}"
  listing_resource_version = "${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.listing_resource_version}"
  oracle_terms_of_use_link = "${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.oracle_terms_of_use_link}"
  signature                = "${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.signature}"
  time_retrieved           = "${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.time_retrieved}"

  timeouts {
    create = "20m"
  }
}

data "oci_core_app_catalog_subscriptions" "test_app_catalog_subscriptions" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  listing_id = "${oci_core_app_catalog_subscription.test_app_catalog_subscription.listing_id}"

  filter {
    name   = "listing_resource_version"
    values = ["${oci_core_app_catalog_subscription.test_app_catalog_subscription.listing_resource_version}"]
  }
}

output "subscriptions" {
  value = ["${data.oci_core_app_catalog_subscriptions.test_app_catalog_subscriptions.app_catalog_subscriptions}"]
}
