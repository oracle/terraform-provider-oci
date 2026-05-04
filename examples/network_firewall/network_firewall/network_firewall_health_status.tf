// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_network_firewall_network_firewall_health_status" "test_network_firewall_health_status" {
  #Required
  network_firewall_id = oci_network_firewall_network_firewall.test_network_firewall.id
}
