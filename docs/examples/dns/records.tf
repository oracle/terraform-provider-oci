/*
 * This file demonstrates dns record usage
 */

resource "oci_dns_record" "record-a" {
  zone_name_or_id = "${oci_dns_zone.zone1.name}"
  domain          = "${oci_dns_zone.zone1.name}"
  rtype           = "A"
  rdata           = "192.168.0.1"
  ttl             = 3600
}

resource "oci_dns_record" "record-aaaa" {
  zone_name_or_id = "${oci_dns_zone.zone1.name}"
  domain          = "${oci_dns_zone.zone1.name}"
  rtype           = "AAAA"
  rdata           = "0000:0000:8a2e:0000:0000:0370:0000:0000"
  ttl             = 3600
}

resource "oci_dns_record" "record-cname" {
  zone_name_or_id = "${oci_dns_zone.zone1.name}"
  domain          = "el.${oci_dns_zone.zone1.name}"
  rtype           = "CNAME"
  rdata           = "${oci_dns_zone.zone1.name}"
  ttl             = 86400
}

resource "oci_dns_record" "record-alias" {
  zone_name_or_id = "${oci_dns_zone.zone1.name}"
  domain          = "${oci_dns_zone.zone1.name}"
  rtype           = "ALIAS"
  rdata           = "red.zone"
  ttl             = 86400
}

resource "oci_dns_record" "record-ns" {
  zone_name_or_id = "${oci_dns_zone.zone1.name}"
  rtype           = "NS"
  rdata           = "ns5.p68.dns.oraclecloud.net"
  domain          = "${oci_dns_zone.zone1.name}"
  ttl             = 86400
}

resource "oci_dns_record" "record-mx" {
  zone_name_or_id = "${oci_dns_zone.zone1.name}"
  rtype           = "MX"
  rdata           = "10 mx.dns.oraclecloud.net"
  domain          = "${oci_dns_zone.zone1.name}"
  ttl             = 86400
}

resource "oci_dns_record" "record-ptr" {
  zone_name_or_id = "${oci_dns_zone.zone1.name}"
  rtype           = "PTR"
  rdata           = "some.other.domain.net"
  domain          = "${oci_dns_zone.zone1.name}"
  ttl             = 86400
}

resource "oci_dns_record" "record-txt" {
  zone_name_or_id = "${oci_dns_zone.zone1.name}"
  rtype           = "TXT"
  rdata           = "arbitrary text"
  domain          = "${oci_dns_zone.zone1.name}"
  ttl             = 86400
}

data "oci_dns_records" "rs" {
  zone_name_or_id = "${oci_dns_zone.zone1.name}"

  # optional
  compartment_id = "${var.compartment_ocid}"
  domain         = "${oci_dns_zone.zone1.name}"
  sort_by        = "rtype"                      # domain|rtype|ttl
  sort_order     = "DESC"                       # ASC|DESC
}

output "records" {
  value = "${data.oci_dns_records.rs.records}"
}
