// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

# Manage single-record and multi-record RRSets, plus other record types.
# See README.md for setup, TTLs and record lifecycle guidance.

variable "rrset_ttl" {
  description = "TTL in seconds shared by the single and multi A record examples."
  type        = number
  default     = 300
}

# The zone, domain and type identify the RRSet; view_id selects the private view.
resource "oci_dns_rrset" "rrset-a-single" {
  zone_name_or_id = oci_dns_zone.zone3.name
  domain          = "single.${oci_dns_zone.zone3.name}"
  rtype           = "A"
  view_id         = oci_dns_view.test_view.id

  # Each item repeats the RRSet domain and type and supplies one record value.
  items {
    domain = "single.${oci_dns_zone.zone3.name}"
    rtype  = "A"
    rdata  = "192.0.2.10"
    ttl    = var.rrset_ttl
  }
}

# Manage both A record values together, with the same TTL for every item.
resource "oci_dns_rrset" "rrset-a" {
  zone_name_or_id = oci_dns_zone.zone3.name
  domain          = "multi.${oci_dns_zone.zone3.name}"
  rtype           = "A"
  view_id         = oci_dns_view.test_view.id

  # Remove an items block to remove that record from the managed RRSet.
  items {
    domain = "multi.${oci_dns_zone.zone3.name}"
    rtype  = "A"
    rdata  = "192.0.2.20"
    ttl    = var.rrset_ttl
  }

  items {
    domain = "multi.${oci_dns_zone.zone3.name}"
    rtype  = "A"
    rdata  = "192.0.2.21"
    ttl    = var.rrset_ttl
  }
}

# Map the zone apex to an IPv6 address.
resource "oci_dns_rrset" "rrset-aaaa" {
  zone_name_or_id = oci_dns_zone.zone3.name
  domain          = oci_dns_zone.zone3.name
  rtype           = "AAAA"
  view_id         = oci_dns_view.test_view.id

  items {
    domain = oci_dns_zone.zone3.name
    rtype  = "AAAA"
    rdata  = "0000:0000:8a2e:0000:0000:0370:0000:0000"
    ttl    = 3600
  }
}

# Point sub.<zone> to the zone apex using a CNAME record.
resource "oci_dns_rrset" "rrset-cname" {
  zone_name_or_id = oci_dns_zone.zone3.name
  domain          = "sub.${oci_dns_zone.zone3.name}"
  rtype           = "CNAME"
  view_id         = oci_dns_view.test_view.id

  items {
    domain = "sub.${oci_dns_zone.zone3.name}"
    rtype  = "CNAME"
    rdata  = oci_dns_zone.zone3.name
    ttl    = 86400
  }
}

# Read sub.<zone> after the CNAME resource creates it.
data "oci_dns_rrset" "test_rrset" {
  domain          = oci_dns_rrset.rrset-cname.domain
  rtype           = "CNAME"
  zone_name_or_id = oci_dns_zone.zone3.id
  view_id         = oci_dns_view.test_view.id
}

# Read the same sub.<zone> CNAME through the plural data source.
data "oci_dns_rrsets" "test_rrsets" {
  domain          = oci_dns_rrset.rrset-cname.domain
  rtype           = "CNAME"
  zone_name_or_id = oci_dns_zone.zone3.id
  view_id         = oci_dns_view.test_view.id
}
