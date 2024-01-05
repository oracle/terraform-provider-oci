// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This file demonstrates creating a zone from a zone file using terraform.
 */

/*
 * Step 1: Create the zone from the zone file.
 *         Run Terraform plan/apply and wait until the zone has been created. Note the OCID of the zone.
 * Step 2: Uncomment the oci_dns_zone below, and comment out the oci_dns_action_create_zone_from_zone_file
           Run `terraform import oci_dns_zone.zone ZONE_OCID`, replacing ZONE_OCID with the OCID of the zone from step 1.
 * Step 3: Run Terraform plan/apply.
           Terraform will show that it is deleting the oci_dns_action_create_zone_from_zone_file resource, but the zone will not be deleted. The zone is now fully managed by the oci_dns_zone resource.
 */

resource "random_string" "random_prefix" {
  length  = 4
  number  = false
  special = false
}

data "oci_identity_tenancy" "tenancy" {
  tenancy_id = var.tenancy_ocid
}

locals {
  zone_name = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example-primary.oci-dns1"
}

resource "oci_dns_action_create_zone_from_zone_file" "zonefile" {
  compartment_id = var.compartment_ocid
  create_zone_from_zone_file_details = "$ORIGIN ${local.zone_name}.\n$TTL 3600\n${local.zone_name}.	IN	SOA	ns1.${local.zone_name}. admin.${local.zone_name}. ( 1 7200 3600 14400 3600)\n${local.zone_name}.	IN	NS	ns1.${local.zone_name}."
}

#resource "oci_dns_zone" "zone" {
#  name = local.zone_name
#  compartment_id = var.compartment_ocid
#  zone_type = "PRIMARY"
#}
