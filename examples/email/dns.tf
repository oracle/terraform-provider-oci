// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// Configure a DNS that has the TXT record to setup the SPF for the email
resource "oci_dns_zone" "zone1" {
  compartment_id = var.tenancy_ocid
  name           = "${data.oci_identity_compartment.compartment.name}2-tf-example-primary.oci-email-dns"
  zone_type      = "PRIMARY"
}

resource "oci_dns_rrset" "record-txt" {
    #Required
    domain = oci_dns_zone.zone1.name
    rtype = "TXT"
    zone_name_or_id = oci_dns_zone.zone1.name
    items {
        #Required
        domain = oci_dns_zone.zone1.name
        rdata = "v=spf1 include:spf.oracleemaildelivery.com -all"
        rtype = "TXT"
        ttl = 86400
    }
}

data "oci_dns_rrset" "rs" {
    #Required
    domain = oci_dns_zone.zone1.name
    rtype = "TXT"
    zone_name_or_id = oci_dns_zone.zone1.name
}

data "oci_identity_compartment" "compartment" {
  id = var.tenancy_ocid
}

output "dns_records" {
  value = data.oci_dns_rrset.rs
}
