// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example file demonstrates how to read AD values from a region and employ filters
 * to isolate specific ADs.
 */

data "oci_identity_availability_domains" "ads" {
  compartment_id = var.tenancy_ocid
}

output "ads" {
  value = data.oci_identity_availability_domains.ads.availability_domains
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

output "ad" {
  value = <<EOF

    data.oci_identity_availability_domain.ad.id
    data.oci_identity_availability_domain.ad.compartment_id
    data.oci_identity_availability_domain.ad.name
    data.oci_identity_availability_domain.ad.ad_number
EOF

}

