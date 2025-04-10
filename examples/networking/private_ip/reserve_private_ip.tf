// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Create Reserve PrivateIP with vnic
resource "oci_core_private_ip" "private_reserve_ip" {
  vnic_id        = data.oci_core_vnic_attachments.instance_vnics.vnic_attachments[0]["vnic_id"]
  display_name   = "reserve_private_ip"
  route_table_id = oci_core_vcn.example_vcn.default_route_table_id
  lifetime       = "RESERVED"
}

# Create Reserve PrivateIP with only subnet
resource "oci_core_private_ip" "private_reserve_ip_available" {
  subnet_id      = oci_core_subnet.example_subnet.id
  display_name   = "available_reserve_private_ip"
  route_table_id = oci_core_vcn.example_vcn.default_route_table_id
  lifetime       = "RESERVED"
}


# List Private IPs
data "oci_core_private_ips" "reserve_private_ip_datasource" {
  depends_on = [oci_core_private_ip.private_reserve_ip, oci_core_private_ip.private_reserve_ip_available]
  subnet_id      = oci_core_subnet.example_subnet.id
}

output "private_reserve_ips" {
  value = [data.oci_core_private_ips.reserve_private_ip_datasource.private_ips]
}