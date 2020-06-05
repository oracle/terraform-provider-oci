resource "oci_dns_rrset" "rrset-a" {
  zone_name_or_id = "${oci_dns_zone.zone3.name}"
  domain          = "${oci_dns_zone.zone3.name}"
  rtype           = "A"

  items {
    domain = "${oci_dns_zone.zone3.name}"
    rtype  = "A"
    rdata  = "192.168.0.3"
    ttl    = 3600
  }

  items {
    domain = "${oci_dns_zone.zone3.name}"
    rtype  = "A"
    rdata  = "192.168.0.4"
    ttl    = 3600
  }
}

resource "oci_dns_rrset" "rrset-aaaa" {
  zone_name_or_id = "${oci_dns_zone.zone3.name}"
  domain          = "${oci_dns_zone.zone3.name}"
  rtype           = "AAAA"

  items {
    domain = "${oci_dns_zone.zone3.name}"
    rtype  = "AAAA"
    rdata  = "0000:0000:8a2e:0000:0000:0370:0000:0000"
    ttl    = 3600
  }
}

resource "oci_dns_rrset" "rrset-cname" {
  zone_name_or_id = "${oci_dns_zone.zone3.name}"
  domain          = "el.${oci_dns_zone.zone3.name}"
  rtype           = "CNAME"

  items {
    domain = "el.${oci_dns_zone.zone3.name}"
    rtype  = "CNAME"
    rdata  = "${oci_dns_zone.zone3.name}"
    ttl    = 86400
  }
}
