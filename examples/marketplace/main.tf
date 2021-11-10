// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file
variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_marketplace_accepted_agreement" "test_accepted_agreement" {
  #Required
  agreement_id    = oci_marketplace_listing_package_agreement.test_listing_package_agreement.agreement_id
  compartment_id  = var.compartment_ocid
  listing_id      = data.oci_marketplace_listing.test_listing.id
  package_version = data.oci_marketplace_listing.test_listing.default_package_version
  signature       = oci_marketplace_listing_package_agreement.test_listing_package_agreement.signature
}

resource "oci_marketplace_listing_package_agreement" "test_listing_package_agreement" {
  #Required
  agreement_id    = data.oci_marketplace_listing_package_agreements.test_listing_package_agreements.agreements[0].id
  listing_id      = data.oci_marketplace_listing.test_listing.id
  package_version = data.oci_marketplace_listing.test_listing.default_package_version

  #Optional
  compartment_id = var.compartment_ocid
}

data "oci_marketplace_listing_package_agreements" "test_listing_package_agreements" {
  #Required
  listing_id      = data.oci_marketplace_listing.test_listing.id
  package_version = data.oci_marketplace_listing.test_listing.default_package_version

  #Optional
  compartment_id = var.compartment_ocid
}

data "oci_marketplace_listing_package" "test_listing_package" {
  #Required
  listing_id      = data.oci_marketplace_listing.test_listing.id
  package_version = data.oci_marketplace_listing.test_listing.default_package_version

  #Optional
  compartment_id = var.compartment_ocid
}

data "oci_marketplace_listing_packages" "test_listing_packages" {
  #Required
  listing_id = data.oci_marketplace_listing.test_listing.id

  #Optional
  compartment_id = var.compartment_ocid
}

data "oci_marketplace_listing" "test_listing" {
  listing_id     = data.oci_marketplace_listings.test_listings.listings[0].id
  compartment_id = var.compartment_ocid
}

data "oci_marketplace_listings" "test_listings" {
  category       = ["Analytics"]
  compartment_id = var.compartment_ocid
}

data "oci_marketplace_publishers" "test_publishers" {
  compartment_id = var.compartment_ocid
}

data "oci_marketplace_categories" "test_categories" {
  filter {
    name   = "name"
    values = ["Analytics"]
  }
}

data "oci_marketplace_listing_taxes" "test_listing_taxes" {
  #Required
  listing_id = data.oci_marketplace_listings.test_listings.listings.0.id

  #Optional
  compartment_id = var.compartment_ocid
}
