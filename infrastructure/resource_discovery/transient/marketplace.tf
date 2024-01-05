// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_marketplace_accepted_agreement" "accepted_agreement_rd" {
  #Required
  agreement_id    = "${oci_marketplace_listing_package_agreement.listing_package_agreement_rd.agreement_id}"
  compartment_id  = "${var.compartment_ocid}"
  listing_id      = "${data.oci_marketplace_listing.listing_rd.id}"
  package_version = "${data.oci_marketplace_listing.listing_rd.default_package_version}"
  signature       = "${oci_marketplace_listing_package_agreement.listing_package_agreement_rd.signature}"
}

resource "oci_marketplace_listing_package_agreement" "listing_package_agreement_rd" {
  #Required
  agreement_id    = "${data.oci_marketplace_listing_package_agreements.listing_package_agreements_rd.agreements.0.id}"
  listing_id      = "${data.oci_marketplace_listing.listing_rd.id}"
  package_version = "${data.oci_marketplace_listing.listing_rd.default_package_version}"
}

data "oci_marketplace_listing_package_agreements" "listing_package_agreements_rd" {
  #Required
  listing_id      = "${data.oci_marketplace_listing.listing_rd.id}"
  package_version = "${data.oci_marketplace_listing.listing_rd.default_package_version}"

  #Optional
  compartment_id = "${var.compartment_ocid}"
}

data "oci_marketplace_listing" "listing_rd" {
  listing_id     = "${data.oci_marketplace_listings.listings_rd.listings.0.id}"
  compartment_id = "${var.compartment_ocid}"
}

data "oci_marketplace_listings" "listings_rd" {
  category       = ["Analytics"]
  compartment_id = "${var.compartment_ocid}"
}
