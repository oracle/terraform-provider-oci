// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

# VCN comes with default route table, security list and DHCP options

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_id" {
}

variable "region" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_core_services" "test_services" {
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = var.compartment_id
}

data "oci_ocvp_supported_vmware_software_versions" "test_supported_vmware_software_versions" {
  compartment_id = "${var.compartment_id}"
}

resource "oci_core_vcn" "test_vcn_ocvp" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name   = "VmWareOCVP"
  dns_label      = "vmwareocvp"
}

resource oci_core_nat_gateway test_nat_gateway_ocvp {
  block_traffic  = "false"
  compartment_id = var.compartment_id

  display_name = "NAT Gateway OCVP"
  freeform_tags = {
    "VCN" = "VCN-2020-09-11T00:43:42"
  }
  vcn_id = oci_core_vcn.test_vcn_ocvp.id
}

resource oci_core_route_table test_route_table_for_vsphere_vlan {
  compartment_id = var.compartment_id

  display_name = "Route Table for VLAN-grk-vSphere"
  freeform_tags = {
    "VMware" = "VMware-2020-09-11T00:47:02"
  }
  route_rules {
    #description = <<Optional value not found in discovery>>
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_nat_gateway.test_nat_gateway_ocvp.id
  }
  vcn_id = oci_core_vcn.test_vcn_ocvp.id
}

resource "oci_core_network_security_group" "test_nsg_allow_all" {
  compartment_id = var.compartment_id
  display_name   = "nsg-allow-all"
  vcn_id         = oci_core_vcn.test_vcn_ocvp.id
}

resource "oci_core_network_security_group_security_rule" "test_nsg_security_rule_1" {
  destination_type          = ""
  direction                 = "INGRESS"
  network_security_group_id = oci_core_network_security_group.test_nsg_allow_all.id
  protocol                  = "all"
  source                    = "10.0.0.0/16"
  source_type               = "CIDR_BLOCK"
}

resource "oci_core_network_security_group_security_rule" "test_nsg_security_rule_2" {
  destination               = "0.0.0.0/0"
  destination_type          = "CIDR_BLOCK"
  direction                 = "EGRESS"
  network_security_group_id = oci_core_network_security_group.test_nsg_allow_all.id
  protocol                  = "all"
  source_type               = ""
}

resource "oci_core_service_gateway" "export_sgw" {
  compartment_id = var.compartment_id
  display_name   = "sgw"

  services {
    service_id = data.oci_core_services.test_services.services[0]["id"]
  }

  vcn_id = oci_core_vcn.test_vcn_ocvp.id
}

resource "oci_core_default_dhcp_options" "default_dhcp_options_ocvp" {
  display_name               = "Default DHCP Options for OCVP"
  manage_default_resource_id = oci_core_vcn.test_vcn_ocvp.default_dhcp_options_id

  options {
    custom_dns_servers = []
    server_type        = "VcnLocalPlusInternet"
    type               = "DomainNameServer"
  }

  options {
    search_domain_names = ["vmwareocvp.oraclevcn.com"]
    type                = "SearchDomain"
  }
}

resource "oci_core_route_table" "private_rt" {
  compartment_id = var.compartment_id
  display_name   = "private-rt"

  route_rules {
    destination       = data.oci_core_services.test_services.services[0]["cidr_block"]
    destination_type  = "SERVICE_CIDR_BLOCK"
    network_entity_id = oci_core_service_gateway.export_sgw.id
  }

  vcn_id = oci_core_vcn.test_vcn_ocvp.id
}

resource "oci_core_security_list" "private_sl" {
  compartment_id = var.compartment_id
  display_name   = "private-sl"

  egress_security_rules {
    destination      = "0.0.0.0/0"
    destination_type = "CIDR_BLOCK"
    protocol         = "all"
    stateless        = "false"
  }

  ingress_security_rules {
    description = "TCP traffic for ports: 22 SSH Remote Login Protocol"
    protocol    = "6"
    source      = "10.0.0.0/16"
    source_type = "CIDR_BLOCK"
    stateless   = "false"

    tcp_options {
      max = "22"
      min = "22"
    }
  }

  ingress_security_rules {
    description = "ICMP traffic for: 3 Destination Unreachable"

    icmp_options {
      code = "3"
      type = "3"
    }

    protocol    = "1"
    source      = "10.0.0.0/16"
    source_type = "CIDR_BLOCK"
    stateless   = "false"
  }

  ingress_security_rules {
    protocol    = "all"
    source      = "0.0.0.0/0"
    source_type = "CIDR_BLOCK"
    stateless   = "false"
  }

  vcn_id = oci_core_vcn.test_vcn_ocvp.id
}

resource "oci_core_default_security_list" "default_security_list_ocvp" {
  display_name = "Default Security List for OCVP"

  egress_security_rules {
    destination      = "0.0.0.0/0"
    destination_type = "CIDR_BLOCK"
    protocol         = "all"
    stateless        = "false"
  }

  ingress_security_rules {
    protocol    = "6"
    source      = "0.0.0.0/0"
    source_type = "CIDR_BLOCK"
    stateless   = "false"

    tcp_options {
      max = "22"
      min = "22"
    }
  }

  ingress_security_rules {
    icmp_options {
      code = "4"
      type = "3"
    }

    protocol    = "1"
    source      = "0.0.0.0/0"
    source_type = "CIDR_BLOCK"
    stateless   = "false"
  }

  ingress_security_rules {
    icmp_options {
      code = "-1"
      type = "3"
    }

    protocol    = "1"
    source      = "10.0.0.0/16"
    source_type = "CIDR_BLOCK"
    stateless   = "false"
  }

  manage_default_resource_id = oci_core_vcn.test_vcn_ocvp.default_security_list_id
}

resource "oci_core_subnet" "test_provisioning_subnet" {
  cidr_block                 = "10.0.103.128/25"
  compartment_id             = var.compartment_id
  dhcp_options_id            = oci_core_vcn.test_vcn_ocvp.default_dhcp_options_id
  display_name               = "provisioning-subnet"
  dns_label                  = "provisioningsub"
  prohibit_public_ip_on_vnic = "true"
  route_table_id             = oci_core_route_table.private_rt.id
  security_list_ids          = [oci_core_security_list.private_sl.id]
  vcn_id                     = oci_core_vcn.test_vcn_ocvp.id
}

resource "oci_core_vlan" "test_nsx_edge_uplink2_vlan" {
  display_name        = "NSX-Edge-UP2"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[0]["name"]
  cidr_block          = "10.0.103.0/25"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_nsx_edge_uplink1_vlan" {
  display_name        = "NSX-Edge-UP1"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[0]["name"]
  cidr_block          = "10.0.100.0/25"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_nsx_vtep_vlan" {
  display_name        = "NSX-vTep"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[0]["name"]
  cidr_block          = "10.0.101.0/25"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_nsx_edge_vtep_vlan" {
  display_name        = "NSX Edge-vTep"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[0]["name"]
  cidr_block          = "10.0.102.0/25"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_vsan_net_vlan" {
  display_name        = "vSAN-Net"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[0]["name"]
  cidr_block          = "10.0.101.128/25"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_vmotion_net_vlan" {
  display_name        = "vMotion-Net"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[0]["name"]
  cidr_block          = "10.0.102.128/25"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_vsphere_net_vlan" {
  display_name        = "vSphere-Net"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[0]["name"]
  cidr_block          = "10.0.100.128/26"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_route_table.test_route_table_for_vsphere_vlan.id
}

resource "oci_core_vlan" "test_hcx_vlan" {
  display_name        = "hcx"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[0]["name"]
  cidr_block          = "10.0.100.192/26"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_ocvp_sddc" "test_sddc" {
  // Required
  compartment_id              = var.compartment_id
  compute_availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[0]["name"]
  esxi_hosts_count            = "3"
  hcx_vlan_id                 = oci_core_vlan.test_hcx_vlan.id
  is_hcx_enabled              = true
  nsx_edge_uplink1vlan_id     = oci_core_vlan.test_nsx_edge_uplink1_vlan.id
  nsx_edge_uplink2vlan_id     = oci_core_vlan.test_nsx_edge_uplink2_vlan.id
  nsx_edge_vtep_vlan_id       = oci_core_vlan.test_nsx_edge_vtep_vlan.id
  nsx_vtep_vlan_id            = oci_core_vlan.test_nsx_vtep_vlan.id
  provisioning_subnet_id      = oci_core_subnet.test_provisioning_subnet.id
  ssh_authorized_keys         = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
  vmotion_vlan_id             = oci_core_vlan.test_vmotion_net_vlan.id
  vmware_software_version     = "${lookup(data.oci_ocvp_supported_vmware_software_versions.test_supported_vmware_software_versions.items[1], "version")}"
  vsan_vlan_id                = oci_core_vlan.test_vsan_net_vlan.id
  vsphere_vlan_id             = oci_core_vlan.test_vsphere_net_vlan.id
  // Optional
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.sddc_defined_tags_value}"}
  #display_name  = var.sddc_display_name
  #freeform_tags = var.sddc_freeform_tags
  #instance_display_name_prefix = "prefix"
  #workload_network_cidr = "172.20.0.0/24"
}

resource "oci_ocvp_esxi_host" "test_esxi_host" {
  #Required
  sddc_id = oci_ocvp_sddc.test_sddc.id
  #Optional
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.esxihost_defined_tags_value}"}
  #display_name  = var.esxihost_display_name
  #freeform_tags = var.esxihost_freeform_tags
}

data "oci_ocvp_sddcs" "test_sddcs" {
  compartment_id              = var.compartment_id
  compute_availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[0]["name"]
  display_name                = "displayName"

  filter {
    name   = "id"
    values = [oci_ocvp_sddc.test_sddc.id]
  }

  state = "ACTIVE"
}

data "oci_ocvp_sddc" "test_sddc" {
  sddc_id = oci_ocvp_sddc.test_sddc.id
}

data "oci_ocvp_esxi_hosts" "test_esxi_hosts" {
  compute_instance_id = oci_ocvp_esxi_host.test_esxi_host.compute_instance_id
  display_name        = "displayName"

  filter {
    name   = "id"
    values = [oci_ocvp_esxi_host.test_esxi_host.id]
  }

  sddc_id = oci_ocvp_sddc.test_sddc.id
  state   = "ACTIVE"
}

data "oci_ocvp_esxi_host" "test_esxi_host" {
  esxi_host_id = oci_ocvp_esxi_host.test_esxi_host.id
}

