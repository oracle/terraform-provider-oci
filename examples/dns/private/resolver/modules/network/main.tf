// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

# Create one network and its listening and forwarding endpoints.
# See ../../README.md for setup and a walkthrough.

terraform {
  required_providers {
    oci = {
      source = "oracle/oci"
    }
  }
}

variable "compartment_ocid" {
  description = "Compartment for this network and its DNS endpoints."
  type        = string
}

variable "name_prefix" {
  description = "Name prefix that distinguishes this network from its peer."
  type        = string
}

variable "vcn_cidr" {
  description = "IPv4 CIDR for the VCN; the subnet uses eight additional prefix bits."
  type        = string
}

resource "oci_core_vcn" "example" {
  compartment_id = var.compartment_ocid
  cidr_block     = var.vcn_cidr
  display_name   = var.name_prefix
}

resource "oci_core_subnet" "example" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.example.id
  cidr_block     = cidrsubnet(var.vcn_cidr, 8, 0)
  display_name   = var.name_prefix
}

# The VCN owns the resolver. Read its ID without depending on resolver rules.
data "oci_core_vcn_dns_resolver_association" "example" {
  vcn_id = oci_core_vcn.example.id
}

# Forwarding rules use this endpoint as their source.
resource "oci_dns_resolver_endpoint" "forwarding" {
  resolver_id   = data.oci_core_vcn_dns_resolver_association.example.dns_resolver_id
  subnet_id     = oci_core_subnet.example.id
  name          = "forwarding"
  is_forwarding = true
  is_listening  = false
  scope         = "PRIVATE"
}

# The peer resolver sends matching queries to this listening address.
resource "oci_dns_resolver_endpoint" "listening" {
  resolver_id   = data.oci_core_vcn_dns_resolver_association.example.dns_resolver_id
  subnet_id     = oci_core_subnet.example.id
  name          = "listening"
  is_forwarding = false
  is_listening  = true
  scope         = "PRIVATE"
}

output "resolver_id" {
  description = "VCN-owned resolver ID for configuration in the calling module."
  value       = data.oci_core_vcn_dns_resolver_association.example.dns_resolver_id
}

output "forwarding_endpoint_name" {
  description = "Endpoint name referenced by this network's forwarding rules."
  value       = oci_dns_resolver_endpoint.forwarding.name
}

output "listening_address" {
  description = "Listening address used as the peer resolver's forwarding destination."
  value       = oci_dns_resolver_endpoint.listening.listening_address
}
