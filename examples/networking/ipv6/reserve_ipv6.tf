// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Create Reserved IPv6 with vnic
resource "oci_core_ipv6" "test_reserve_ipv6" {
  vnic_id        = data.oci_core_vnic_attachments.instance_vnics.vnic_attachments[0]["vnic_id"]
  subnet_id      = oci_core_subnet.example_subnet.id
  display_name   = "reserved_ipv6"
  route_table_id = oci_core_vcn.example_vcn.default_route_table_id
  lifetime       = "RESERVED"
}

# Create Reserved IPv6 with only subnet
resource "oci_core_ipv6" "test_reserve_ipv6_available" {
  subnet_id      = oci_core_subnet.example_subnet.id
  display_name   = "available_reserved_ipv6"
  route_table_id = oci_core_vcn.example_vcn.default_route_table_id
  lifetime       = "RESERVED"
}

# List IPv6s
data "oci_core_ipv6s" "reserve_ipv6_datasource" {
  depends_on = [oci_core_ipv6.test_reserve_ipv6,oci_core_ipv6.test_reserve_ipv6_available]
  subnet_id  = oci_core_subnet.example_subnet.id
}

output "reserve_ipv6s" {
  value = [data.oci_core_ipv6s.reserve_ipv6_datasource.ipv6s]
}