// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_core_app_catalog_listings" "test_app_catalog_listings" {
}

data "oci_core_app_catalog_listing_resource_versions" "test_app_catalog_listing_resource_versions" {
  #Required
  listing_id = data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings[0]["listing_id"]
}

data "oci_core_app_catalog_listing_resource_version" "test_app_catalog_listing_resource_version" {
  #Required
  listing_id       = data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings[0]["listing_id"]
  resource_version = data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0]["listing_resource_version"]
}

resource "oci_core_app_catalog_listing_resource_version_agreement" "test_app_catalog_listing_resource_version_agreement" {
  #Required
  listing_id               = data.oci_core_app_catalog_listing_resource_version.test_app_catalog_listing_resource_version.listing_id
  listing_resource_version = data.oci_core_app_catalog_listing_resource_version.test_app_catalog_listing_resource_version.listing_resource_version
}

output "agreement_eula_link" {
  value = [oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.eula_link]
}

