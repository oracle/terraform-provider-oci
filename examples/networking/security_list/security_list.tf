// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# This example demonstrates some basic security list options.

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "security_list_egress_security_rules_description" {
  default = "description"
}

variable "security_list_ingress_security_rules_description" {
  default = "description"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_vcn" "example_vcn" {
  cidr_block     = "10.0.0.0/16"
  dns_label      = "examplevcn"
  compartment_id = var.compartment_ocid
  display_name   = "exampleVCN"
}

# Protocols are specified as protocol numbers.
# http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml

resource "oci_core_security_list" "example_security_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.example_vcn.id
  display_name   = "exampleSecurityList"

  // allow outbound tcp traffic on all ports
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "6"
  }

  // allow outbound udp traffic on a port range
  egress_security_rules {
    description = var.security_list_egress_security_rules_description
    destination = "0.0.0.0/0"
    protocol    = "17" // udp
    stateless   = true

    udp_options {
      // These values correspond to the destination port range.
      min = 319
      max = 320
    }
  }

  // allow inbound ssh traffic from a specific port
  ingress_security_rules {
    protocol  = "6" // tcp
    source    = "0.0.0.0/0"
    stateless = false

    tcp_options {
      source_port_range {
        min = 100
        max = 100
      }

      // These values correspond to the destination port range.
      min = 22
      max = 22
    }
  }

  // allow inbound icmp traffic of a specific type
  ingress_security_rules {
    description = var.security_list_ingress_security_rules_description
    protocol    = 1
    source      = "0.0.0.0/0"
    stateless   = true

    icmp_options {
      type = 3
      code = 4
    }
  }
}

