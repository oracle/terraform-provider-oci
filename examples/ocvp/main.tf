// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

# VCN comes with default route table, security list and DHCP options

variable "tenancy_ocid" {
}

variable "compartment_ocid" {
}

variable "region" {
}

#variable "billing_donor_host_id" {
#}

variable "config_file_profile" {}

provider "oci" {
  region              = var.region
  tenancy_ocid        = var.tenancy_ocid
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
}

data "oci_core_services" "test_services" {
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = var.compartment_ocid
}

data "oci_ocvp_supported_vmware_software_versions" "test_supported_vmware_software_versions" {
  compartment_id = "${var.compartment_ocid}"
}

data "oci_ocvp_supported_skus" "test_supported_skus" {
  compartment_id = "${var.compartment_ocid}"
}

data "oci_ocvp_supported_host_shapes" "test_supported_host_shapes" {
  // Required
  compartment_id = "${var.compartment_ocid}"
  // Optional
  name = "BM.Standard2.52"
  sddc_type = "PRODUCTION"
}

resource "oci_core_vcn" "test_vcn_ocvp" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "VmWareOCVP"
  dns_label      = "vmwareocvp"
}

resource oci_core_nat_gateway test_nat_gateway_ocvp {
  block_traffic  = "false"
  compartment_id = var.compartment_ocid

  display_name = "NAT Gateway OCVP"
  freeform_tags = {
    "VCN" = "VCN-2020-09-11T00:43:42"
  }
  vcn_id = oci_core_vcn.test_vcn_ocvp.id
}

resource oci_core_route_table test_route_table_for_vsphere_vlan {
  compartment_id = var.compartment_ocid

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
  compartment_id = var.compartment_ocid
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
  compartment_id = var.compartment_ocid
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
  compartment_id = var.compartment_ocid
  display_name   = "private-rt"

  route_rules {
    destination       = data.oci_core_services.test_services.services[0]["cidr_block"]
    destination_type  = "SERVICE_CIDR_BLOCK"
    network_entity_id = oci_core_service_gateway.export_sgw.id
  }

  vcn_id = oci_core_vcn.test_vcn_ocvp.id
}

resource "oci_core_security_list" "private_sl" {
  compartment_id = var.compartment_ocid
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
  compartment_id             = var.compartment_ocid
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
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  cidr_block          = "10.0.103.0/25"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_nsx_edge_uplink1_vlan" {
  display_name        = "NSX-Edge-UP1"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  cidr_block          = "10.0.100.0/25"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_nsx_vtep_vlan" {
  display_name        = "NSX-vTep"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  cidr_block          = "10.0.101.0/25"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_nsx_edge_vtep_vlan" {
  display_name        = "NSX Edge-vTep"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  cidr_block          = "10.0.102.0/25"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_vsan_net_vlan" {
  display_name        = "vSAN-Net"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  cidr_block          = "10.0.101.128/25"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_vmotion_net_vlan" {
  display_name        = "vMotion-Net"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  cidr_block          = "10.0.102.128/25"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_vsphere_net_vlan" {
  display_name        = "vSphere-Net"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  cidr_block          = "10.0.100.128/26"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_route_table.test_route_table_for_vsphere_vlan.id
}

resource "oci_core_vlan" "test_hcx_vlan" {
  display_name        = "hcx"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  cidr_block          = "10.0.100.192/26"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_provisioning_vlan" {
  display_name        = "provisioning-vlan"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  cidr_block          = "10.0.104.128/25"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_replication_vlan" {
  display_name        = "replication-vlan"
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  cidr_block          = "10.0.104.0/25"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_volume" "test_block_volume" {
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  compartment_id      = var.compartment_ocid
  display_name        = "TestBlockVolume"
  size_in_gbs         = "4096"
}

resource "oci_core_compute_capacity_reservation" "test_compute_capacity_reservation" {
  #Required
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  compartment_id = var.compartment_ocid

  instance_reservation_configs {
    #Required
    instance_shape = "BM.Standard2.52"
    reserved_count = 2
    fault_domain = "FAULT-DOMAIN-1"
  }
  instance_reservation_configs {
    #Required
    instance_shape = "BM.Standard2.52"
    reserved_count = 1
    fault_domain = "FAULT-DOMAIN-2"
  }
  instance_reservation_configs {
    #Required
    instance_shape = "BM.Standard2.52"
    reserved_count = 1
    fault_domain = "FAULT-DOMAIN-3"
  }
}

// SDDC resource with deprecated fields
resource "oci_ocvp_sddc" "test_sddc_deprecated" {
  // Required
  compartment_id              = var.compartment_ocid
  compute_availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  esxi_hosts_count            = "3"
  nsx_edge_uplink1vlan_id     = oci_core_vlan.test_nsx_edge_uplink1_vlan.id
  nsx_edge_uplink2vlan_id     = oci_core_vlan.test_nsx_edge_uplink2_vlan.id
  nsx_edge_vtep_vlan_id       = oci_core_vlan.test_nsx_edge_vtep_vlan.id
  nsx_vtep_vlan_id            = oci_core_vlan.test_nsx_vtep_vlan.id
  provisioning_subnet_id      = oci_core_subnet.test_provisioning_subnet.id
  ssh_authorized_keys         = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
  vmotion_vlan_id             = oci_core_vlan.test_vmotion_net_vlan.id
  vmware_software_version     = "7.0 update 3"
  vsan_vlan_id                = oci_core_vlan.test_vsan_net_vlan.id
  vsphere_vlan_id             = oci_core_vlan.test_vsphere_net_vlan.id
  // Optional
  provisioning_vlan_id        = oci_core_vlan.test_provisioning_vlan.id
  replication_vlan_id         = oci_core_vlan.test_replication_vlan.id
  hcx_vlan_id                 = oci_core_vlan.test_hcx_vlan.id
  is_hcx_enabled              = true
  initial_sku                 = "HOUR"
  initial_host_ocpu_count     = "12.0"
  initial_host_shape_name     = "BM.Standard2.52"
  capacity_reservation_id     = oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id
  datastores {
    #Required
    block_volume_ids = ["${oci_core_volume.test_block_volume.id}"]
    datastore_type   = "MANAGEMENT"
  }
  is_shielded_instance_enabled = false
  hcx_action                   = "upgrade"
  refresh_hcx_license_status   = true
  #reserving_hcx_on_premise_license_keys = var.reserving_hcx_on_premise_license_keys
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.sddc_defined_tags_value}"}
  #display_name  = var.sddc_display_name
  #freeform_tags = var.sddc_freeform_tags
  #instance_display_name_prefix = "prefix"
  #workload_network_cidr = "172.20.0.0/24"
}

resource "oci_ocvp_sddc" "test_sddc" {
  // Required
  compartment_id          = var.compartment_ocid
  ssh_authorized_keys     = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
  vmware_software_version = "7.0 update 3"
  initial_configuration {
    initial_cluster_configurations {
      // required
      compute_availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
      esxi_hosts_count            = 3
      vsphere_type                = "MANAGEMENT"

      network_configuration {
        nsx_edge_uplink1vlan_id = oci_core_vlan.test_nsx_edge_uplink1_vlan.id
        nsx_edge_uplink2vlan_id = oci_core_vlan.test_nsx_edge_uplink2_vlan.id
        nsx_edge_vtep_vlan_id   = oci_core_vlan.test_nsx_edge_vtep_vlan.id
        nsx_vtep_vlan_id        = oci_core_vlan.test_nsx_vtep_vlan.id
        provisioning_subnet_id  = oci_core_subnet.test_provisioning_subnet.id
        vmotion_vlan_id         = oci_core_vlan.test_vmotion_net_vlan.id
        vsan_vlan_id            = oci_core_vlan.test_vsan_net_vlan.id
        vsphere_vlan_id         = oci_core_vlan.test_vsphere_net_vlan.id
        provisioning_vlan_id    = oci_core_vlan.test_provisioning_vlan.id
        replication_vlan_id     = oci_core_vlan.test_replication_vlan.id
        hcx_vlan_id             = oci_core_vlan.test_hcx_vlan.id
      }

      // optional
      initial_host_ocpu_count      = "12.0"
      initial_host_shape_name      = "BM.Standard2.52"
      instance_display_name_prefix = "prefix"
      is_shielded_instance_enabled = true
      datastores {
        #Required
        block_volume_ids = ["${oci_core_volume.test_block_volume.id}"]
        datastore_type   = "MANAGEMENT"
      }
      workload_network_cidr = "172.20.0.0/24"
    }
  }

  //optional
  is_hcx_enabled               = true
  hcx_action                   = "upgrade"
  refresh_hcx_license_status   = true
  #reserving_hcx_on_premise_license_keys = var.reserving_hcx_on_premise_license_keys
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.sddc_defined_tags_value}"}
  #display_name  = var.sddc_display_name
  #freeform_tags = var.sddc_freeform_tags
}

resource "oci_ocvp_cluster" "test_cluster" {
  // Required
  sddc_id                     = oci_ocvp_sddc.test_sddc.id
  compute_availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  esxi_hosts_count            = "3"
  vmware_software_version     = "7.0 update 3"
  network_configuration {
    nsx_edge_uplink1vlan_id = oci_core_vlan.test_nsx_edge_uplink1_vlan.id
    nsx_edge_uplink2vlan_id = oci_core_vlan.test_nsx_edge_uplink2_vlan.id
    nsx_edge_vtep_vlan_id   = oci_core_vlan.test_nsx_edge_vtep_vlan.id
    nsx_vtep_vlan_id        = oci_core_vlan.test_nsx_vtep_vlan.id
    provisioning_subnet_id  = oci_core_subnet.test_provisioning_subnet.id
    vmotion_vlan_id         = oci_core_vlan.test_vmotion_net_vlan.id
    vsan_vlan_id            = oci_core_vlan.test_vsan_net_vlan.id
    vsphere_vlan_id         = oci_core_vlan.test_vsphere_net_vlan.id
    provisioning_vlan_id    = oci_core_vlan.test_provisioning_vlan.id
    replication_vlan_id     = oci_core_vlan.test_replication_vlan.id
    hcx_vlan_id             = oci_core_vlan.test_hcx_vlan.id
  }

  // optional
  initial_commitment          = "HOUR"
  initial_host_ocpu_count     = "12.0"
  initial_host_shape_name     = "BM.Standard2.52"
  capacity_reservation_id     = oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id
  is_shielded_instance_enabled = true
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.sddc_defined_tags_value}"}
  #display_name  = var.sddc_display_name
  #freeform_tags = var.sddc_freeform_tags
  #instance_display_name_prefix = "prefix"
  #workload_network_cidr = "172.20.0.0/24"
}

// ESXi host resource with deprecated fields
resource "oci_ocvp_esxi_host" "test_esxi_host_deprecated" {
  #Required
  sddc_id = oci_ocvp_sddc.test_sddc.id
  #Optional
  compute_availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  current_sku = "HOUR"
  host_ocpu_count             = "12.0"
  host_shape_name             = "BM.Standard2.52"
  next_sku    = "HOUR"
  #non_upgraded_esxi_host_id = data.oci_ocvp_esxi_hosts.non_upgraded_esxi_hosts.esxi_host_collection[0].id
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.esxihost_defined_tags_value}"}
  #display_name  = var.esxihost_display_name
  #freeform_tags = var.esxihost_freeform_tags
  #failed_esxi_host_id = var.failed_esxi_host_ocid
  #billing_donor_host_id = var.billing_donor_host_id
}

resource "oci_ocvp_esxi_host" "test_esxi_host" {
  #Required
  cluster_id = oci_ocvp_cluster.test_cluster.id
  #Optional
  compute_availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  host_ocpu_count             = "12.0"
  host_shape_name             = "BM.Standard2.52"
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.esxihost_defined_tags_value}"}
  #display_name  = var.esxihost_display_name
  #freeform_tags = var.esxihost_freeform_tags
  #failed_esxi_host_id = var.failed_esxi_host_ocid
  #billing_donor_host_id = var.billing_donor_host_id
}

data "oci_ocvp_sddcs" "test_sddcs" {
  compartment_id              = var.compartment_ocid
  compute_availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
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

data "oci_ocvp_clusters" "test_clusters" {
  sddc_id        = oci_ocvp_sddc.test_sddc.id
  compartment_id = var.compartment_ocid
}

data "oci_ocvp_cluster" "test_cluster" {
  cluster_id        = oci_ocvp_cluster.test_cluster.id
}

data "oci_ocvp_esxi_hosts" "test_esxi_hosts" {
  compute_instance_id    = oci_ocvp_esxi_host.test_esxi_host.compute_instance_id
  display_name           = "displayName"
  is_swap_billing_only   = false
  is_billing_donors_only = false

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