// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This file demonstrates DNS resolver creation with an endpoint and resolver rules
 */

/*
 * Step 1: Create core VCN, subnet and DNS resolver.
 *         Run Terraform plan/apply and wait until created VCN's state shows as Available in OCI console.
 * Step 2: Uncomment the DNS resources below and comment out the resolver rules.
 *         Run Terraform refresh/plan/apply to update the DNS resolver and endpoint.
 * Step 3: Uncomment DNS resolver rules.  Run Terraform plan/apply to add rules to the resolver.
 */

/*
resource "oci_dns_resolver" "test_resolver" {
  attached_views {
    view_id = "${oci_dns_view.test_view.id}"
  }
  display_name = "test_resolver"
  resolver_id = "${data.oci_core_vcn_dns_resolver_association.test_vcn_dns_resolver_association.dns_resolver_id}"
  scope = "PRIVATE"

  // In Step 2: comment out rules here while creating resolver and endpoint
  rules {
    action = "FORWARD"
    client_address_conditions = ["192.0.20.0/24"]
    destination_addresses = ["10.0.0.11"]
    qname_cover_conditions = []
    source_endpoint_name = "test_endpoint"
  }
  rules {
    action = "FORWARD"
    client_address_conditions = []
    destination_addresses = ["10.0.0.11"]
    qname_cover_conditions = ["internal.example.com"]
    source_endpoint_name = "test_endpoint"
  }
}

resource "oci_dns_view" "test_view" {
  compartment_id = "${var.compartment_ocid}"
  scope = "PRIVATE"
  display_name = "test_view"
}

resource "oci_dns_resolver_endpoint" "test_resolver_endpoint" {
  endpoint_type = "VNIC"
  forwarding_address = "10.0.0.5"
  is_forwarding = "true"
  is_listening = "false"
  name = "test_endpoint"
  resolver_id = "${oci_dns_resolver.test_resolver.id}"
  scope = "PRIVATE"
  subnet_id = "${oci_core_subnet.test_subnet.id}"
}
*/