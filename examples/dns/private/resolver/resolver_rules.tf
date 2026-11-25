// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

# Compose forwarding rules between the two networks in the root module.
# See README.md for setup and a walkthrough.

# Compose rules here using endpoint outputs from both network modules.
# These references order rule configuration after endpoint creation.
resource "oci_dns_resolver" "resolver_a" {
  resolver_id  = module.network_a.resolver_id
  display_name = "${var.name_prefix}-a"
  scope        = "PRIVATE"

  rules {
    action                 = "FORWARD"
    destination_addresses  = [module.network_b.listening_address]
    qname_cover_conditions = ["b.example.internal"]
    source_endpoint_name   = module.network_a.forwarding_endpoint_name
  }
}

# Configure the reverse direction using a different query suffix.
resource "oci_dns_resolver" "resolver_b" {
  resolver_id  = module.network_b.resolver_id
  display_name = "${var.name_prefix}-b"
  scope        = "PRIVATE"

  rules {
    action                 = "FORWARD"
    destination_addresses  = [module.network_a.listening_address]
    qname_cover_conditions = ["a.example.internal"]
    source_endpoint_name   = module.network_b.forwarding_endpoint_name
  }
}
