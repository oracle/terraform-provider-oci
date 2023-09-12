// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This file demonstrates initial setup of a dnssec enabled zone when the zone's
 * parent zone is in OCI. It does not demonstrate setting up dnssec for the parent
 * zone or handle rotating the dnssec key versions.
 */

resource "random_string" "random_prefix" {
  length  = 4
  numeric = false
  special = false
}

resource "oci_dns_zone" "dnssec_parent_zone" {
  compartment_id = var.compartment_ocid
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example-dnssec-parent.oci-dns"
  zone_type      = "PRIMARY"
  scope          = "GLOBAL"
  dnssec_state   = "ENABLED"
}

resource "oci_dns_zone" "dnssec_child_zone" {
  compartment_id = var.compartment_ocid
  name           = "child.${oci_dns_zone.dnssec_parent_zone.name}"
  zone_type      = "PRIMARY"
  scope          = "GLOBAL"
  dnssec_state   = "ENABLED"
}

resource "oci_dns_rrset" "parent_zone_ns_rrset" {
  zone_name_or_id = oci_dns_zone.dnssec_parent_zone.id
  domain          = oci_dns_zone.dnssec_child_zone.name
  rtype           = "NS"

  items {
    domain = oci_dns_zone.dnssec_child_zone.name
    rtype  = "NS"
    rdata  = oci_dns_zone.dnssec_child_zone.nameservers[0].hostname
    ttl    = 86400
  }
}

locals {
  ksk = oci_dns_zone.dnssec_child_zone.dnssec_config[0].ksk_dnssec_key_versions[0]
}

resource "oci_dns_rrset" "parent_zone_ds_rrset" {
  zone_name_or_id = oci_dns_zone.dnssec_parent_zone.id
  domain          = oci_dns_zone.dnssec_child_zone.name
  rtype           = "DS"

  items {
    domain = oci_dns_zone.dnssec_child_zone.name
    rtype  = "DS"
    rdata  = local.ksk.ds_data[0].rdata
    ttl    = 86400
  }

  lifecycle {
    ignore_changes = [
      items,
    ]
  }
}

resource "oci_dns_zone_promote_dnssec_key_version" "promote_dnssec_key_version" {
  dnssec_key_version_uuid = local.ksk.uuid
  zone_id                 = oci_dns_zone.dnssec_child_zone.id
  scope                   = "GLOBAL"
  depends_on              = [oci_dns_rrset.parent_zone_ds_rrset]
  lifecycle {
    ignore_changes = [
      dnssec_key_version_uuid,
    ]
  }
}

data "oci_identity_tenancy" "tenancy" {
  tenancy_id = var.tenancy_ocid
}

