// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This file demonstrates dns private zone management
 */

resource "random_string" "random_prefix" {
  length  = 4
  number  = false
  special = false
}

resource "oci_dns_zone" "zone1" {
  compartment_id = var.compartment_ocid
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example-primary.oci-dns1"
  zone_type      = "PRIMARY"
  scope          = "PRIVATE"
  view_id        = oci_dns_view.test_view.id
}

resource "oci_dns_zone" "zone3" {
  compartment_id = var.compartment_ocid
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example3-primary.oci-dns1"
  zone_type      = "PRIMARY"
  scope          = "PRIVATE"
  view_id        = oci_dns_view.test_view.id
}

resource "oci_dns_tsig_key" "test_tsig_key" {
  algorithm      = "hmac-sha1"
  compartment_id = var.compartment_ocid
  name           = "${random_string.random_prefix.result}-test_tsig_key-name"
  secret         = "c2VjcmV0"
}

data "oci_dns_zones" "zs" {
  compartment_id = var.compartment_ocid
  name_contains  = "example"
  state          = "ACTIVE"
  zone_type      = "PRIMARY"
  sort_by        = "name" # name|zoneType|timeCreated
  sort_order     = "DESC" # ASC|DESC
  scope          = "PRIVATE"
  view_id        = oci_dns_view.test_view.id
}

data "oci_identity_tenancy" "tenancy" {
  tenancy_id = var.tenancy_ocid
}

output "zones" {
  value = data.oci_dns_zones.zs.zones
}

