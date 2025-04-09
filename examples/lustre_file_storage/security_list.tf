// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_security_list" "my_security_list" {
  compartment_id = var.compartment_ocid
  display_name   = "my_security_list"
  vcn_id         = oci_core_vcn.my_vcn.id

  // Allow all outbound requests
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "all"
  }

  ingress_security_rules {
    // Allowing inbound SSH traffic to instances in the subnet from any source
    protocol = "6"
    source   = "0.0.0.0/0"

    tcp_options {
      min = 22
      max = 22
    }
  }

  ingress_security_rules {
    // Allowing inbound ICMP traffic of a specific type and code from any source
    protocol = 1
    source   = "0.0.0.0/0"

    icmp_options {
      type = 3
      code = 4
    }
  }

  ingress_security_rules {
    // Allowing inbound ICMP traffic of a specific type and code from any source
    protocol = 1
    source   = "10.0.0.0/16"
  }

   ingress_security_rules {
     protocol = "6"
     source   = "10.0.0.0/16"
   }
}

