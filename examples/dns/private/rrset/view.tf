// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Create the private view used by the zones and RRSets.
# See README.md for setup and the RRSet walkthrough.

resource "oci_dns_view" "test_view" {
  compartment_id = var.compartment_ocid
  scope          = "PRIVATE"
}

