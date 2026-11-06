// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

# Create two independent networks with resolver endpoints.
# See README.md for setup and a walkthrough.

variable "name_prefix" {
  description = "Prefix for disposable example resources."
  type        = string
  default     = "dns-resolver-example"
}

# Both modules create endpoints independently; resolver_rules.tf connects them.
module "network_a" {
  source           = "./modules/network"
  compartment_ocid = var.compartment_ocid
  name_prefix      = "${var.name_prefix}-a"
  vcn_cidr         = "10.50.0.0/16"
}

module "network_b" {
  source           = "./modules/network"
  compartment_ocid = var.compartment_ocid
  name_prefix      = "${var.name_prefix}-b"
  vcn_cidr         = "10.60.0.0/16"
}
