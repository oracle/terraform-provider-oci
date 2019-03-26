// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

/*
 * This file demonstrates dns zone management
 */

resource "random_string" "random_prefix" {
  length  = 4
  number  = false
  special = false
}

resource "oci_dns_zone" "zone1" {
  compartment_id = "${var.compartment_ocid}"
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example-primary.oci-dns1"
  zone_type      = "PRIMARY"
}

resource "oci_dns_zone" "zone2" {
  compartment_id = "${var.compartment_ocid}"
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example-secondary.oci-dns2"
  zone_type      = "SECONDARY"

  external_masters {
    address = "77.64.12.1"

    tsig {
      algorithm = "hmac-sha1"
      name      = "key-name"
      secret    = "c2VjcmV0"
    }
  }

  external_masters {
    address = "77.64.12.2"

    tsig {
      algorithm = "hmac-sha1"
      name      = "key-name"
      secret    = "c2VjcmV0"
    }
  }
}

data "oci_dns_zones" "zs" {
  compartment_id = "${var.compartment_ocid}"
  name_contains  = "example"
  state          = "ACTIVE"
  zone_type      = "PRIMARY"
  sort_by        = "name"                    # name|zoneType|timeCreated
  sort_order     = "DESC"                    # ASC|DESC
}

data "oci_identity_tenancy" "tenancy" {
  tenancy_id = "${var.tenancy_ocid}"
}

output "zones" {
  value = "${data.oci_dns_zones.zs.zones}"
}
