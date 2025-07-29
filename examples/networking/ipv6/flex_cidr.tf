// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Create flex cidr with vnic
resource "oci_core_ipv6" "test_flex_cidr" {
  vnic_id        = data.oci_core_vnic_attachments.instance_vnics.vnic_attachments[0]["vnic_id"]
  display_name   = "flex_cidr"
  route_table_id = oci_core_vcn.example_vcn.default_route_table_id
  cidr_prefix_length = 80
}

# List IPv6s
data "oci_core_ipv6s" "flex_cidr_datasource" {
  vnic_id    = oci_core_ipv6.test_flex_cidr.vnic_id
}

output "flex_cidrs" {
  value = [data.oci_core_ipv6s.flex_cidr_datasource.ipv6s]
}