// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// Configure a DNS that has the TXT record to setup the SPF for the email
resource "oci_dns_zone" "zone1" {
  compartment_id = var.tenancy_ocid
  name           = "${data.oci_identity_compartment.compartment.name}-tf-example-primary.oci-email-dns"
  zone_type      = "PRIMARY"
}

resource "oci_dns_record" "record-txt" {
  zone_name_or_id = oci_dns_zone.zone1.name
  rtype           = "TXT"
  rdata           = "v=spf1 include:spf.oracleemaildelivery.com -all"
  domain          = oci_dns_zone.zone1.name
  ttl             = 86400
}

data "oci_dns_records" "rs" {
  zone_name_or_id = oci_dns_zone.zone1.name

  # optional
  domain     = oci_dns_zone.zone1.name
  sort_by    = "rtype" # domain|rtype|ttl
  sort_order = "DESC"  # ASC|DESC
}

data "oci_identity_compartment" "compartment" {
  id = var.tenancy_ocid
}

output "dns_records" {
  value = data.oci_dns_records.rs.records
}

