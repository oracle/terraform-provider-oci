// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This file demonstrates DNS rrset management
 */

resource "oci_dns_rrset" "rrset-a" {
  zone_name_or_id = oci_dns_zone.zone3.name
  domain          = oci_dns_zone.zone3.name
  rtype           = "A"
  scope           = "PRIVATE"
  view_id         = oci_dns_view.test_view.id

  items {
    domain = oci_dns_zone.zone3.name
    rtype  = "A"
    rdata  = "192.168.0.3"
    ttl    = 3600
  }

  items {
    domain = oci_dns_zone.zone3.name
    rtype  = "A"
    rdata  = "192.168.0.4"
    ttl    = 3600
  }
}

resource "oci_dns_rrset" "rrset-aaaa" {
  zone_name_or_id = oci_dns_zone.zone3.name
  domain          = oci_dns_zone.zone3.name
  rtype           = "AAAA"
  scope           = "PRIVATE"
  view_id         = oci_dns_view.test_view.id

  items {
    domain = oci_dns_zone.zone3.name
    rtype  = "AAAA"
    rdata  = "0000:0000:8a2e:0000:0000:0370:0000:0000"
    ttl    = 3600
  }
}

resource "oci_dns_rrset" "rrset-cname" {
  zone_name_or_id = oci_dns_zone.zone3.name
  domain          = "el.${oci_dns_zone.zone3.name}"
  rtype           = "CNAME"
  scope           = "PRIVATE"
  view_id         = oci_dns_view.test_view.id

  items {
    domain = "el.${oci_dns_zone.zone3.name}"
    rtype  = "CNAME"
    rdata  = oci_dns_zone.zone3.name
    ttl    = 86400
  }
}

data "oci_dns_rrset" "test_rrset" {
  domain          = "el.${oci_dns_zone.zone3.name}"
  rtype           = "CNAME"
  zone_name_or_id = oci_dns_zone.zone3.id
  scope           = "PRIVATE"
  view_id         = oci_dns_view.test_view.id
}

data "oci_dns_rrsets" "test_rrsets" {
  domain          = "el.${oci_dns_zone.zone3.name}"
  rtype           = "CNAME"
  zone_name_or_id = oci_dns_zone.zone3.id
  scope           = "PRIVATE"
  view_id         = oci_dns_view.test_view.id
}
