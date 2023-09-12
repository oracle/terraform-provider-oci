// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This file demonstrates dns zone management
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
}

resource "oci_dns_tsig_key" "test_tsig_key" {
  algorithm      = "hmac-sha1"
  compartment_id = var.compartment_ocid
  name           = "${random_string.random_prefix.result}-test-tsig-key-name"
  secret         = "c2VjcmV0"
}

resource "oci_dns_zone" "zone2" {
  compartment_id = var.compartment_ocid
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example2-primary.oci-dns2"
  zone_type      = "PRIMARY"

  external_downstreams {
    address     = "77.64.12.1"
    tsig_key_id = oci_dns_tsig_key.test_tsig_key.id
  }

  external_downstreams {
    address     = "77.64.12.2"
    tsig_key_id = oci_dns_tsig_key.test_tsig_key.id
  }
}

resource "oci_dns_zone" "zone3" {
  compartment_id = var.compartment_ocid
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example3-primary.oci-dns1"
  zone_type      = "PRIMARY"
}

resource "oci_dns_zone" "zone4" {
  compartment_id = var.compartment_ocid
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example-primary.oci-dns4"
  zone_type      = "PRIMARY"
  scope          = "GLOBAL"
  dnssec_state   = "ENABLED"
}

resource "oci_dns_zone_stage_dnssec_key_version" "stage_dnssec_key_version" {
  predecessor_dnssec_key_version_uuid = oci_dns_zone.zone4.dnssec_config[0].zsk_dnssec_key_versions[0].uuid
  zone_id                             = oci_dns_zone.zone4.id
  scope                               = "GLOBAL"
}

data "oci_dns_zones" "zs" {
  compartment_id = var.compartment_ocid
  name_contains  = "example"
  state          = "ACTIVE"
  zone_type      = "PRIMARY"
  sort_by        = "name" # name|zoneType|timeCreated
  sort_order     = "DESC" # ASC|DESC
}

data "oci_identity_tenancy" "tenancy" {
  tenancy_id = var.tenancy_ocid
}

output "zones" {
  value = data.oci_dns_zones.zs.zones
}

