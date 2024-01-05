// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "hpc_island_id" {
}

variable "bm_image_id" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
  dns_label      = "dnslabel"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "TestRouteTable"
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = lower(
    // Since sufficient capacity available only in AD3
    data.oci_identity_availability_domains.test_availability_domains.availability_domains[2].name,
  )
  cidr_block        = "10.0.2.0/24"
  compartment_id    = var.compartment_ocid
  dhcp_options_id   = oci_core_vcn.test_vcn.default_dhcp_options_id
  display_name      = "TestSubnet"
  dns_label         = "dnslabel"
  route_table_id    = oci_core_route_table.test_route_table.id
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  vcn_id            = oci_core_vcn.test_vcn.id
}

variable "InstanceImageOCID" {
  type = map(string)

  default = {
    // See https://docs.us-phoenix-1.oraclecloud.com/images/
    // Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaadjnj3da72bztpxinmqpih62c2woscbp6l3wjn36by2cvmdhjub6a"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaawufnve5jxze4xf7orejupw5iq3pms6cuadzjc7klojix6vmk42va"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaagbrvhganmn7awcr7plaaf5vhabmzhx763z5afiitswjwmzh7upna"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaajwtut4l7fo3cvyraate6erdkyf2wdk5vpk6fp6ycng3dv2y3ymvq"
  }
}

resource "oci_core_network_security_group" "test_network_security_group1" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_instance_configuration" "test_instance_configuration" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInstanceConfiguration"

  instance_details {
    instance_type = "compute"

    launch_details {
      // Since sufficient capacity available only in AD3
      availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[2].name
      compartment_id      = var.compartment_ocid

      create_vnic_details {
        assign_public_ip       = "false"
        display_name           = "backend-servers"
        hostname_label         = "hostnameLabel"
        nsg_ids                = [oci_core_network_security_group.test_network_security_group1.id]
        private_ip             = "privateIp"
        skip_source_dest_check = "false"
        subnet_id              = oci_core_subnet.test_subnet.id
      }

      display_name = "backend-servers"

      extended_metadata = {
        "extendedMetadata" = "extendedMetadata"
      }

      ipxe_script = "ipxeScript"

      metadata = {
        "metadata" = "metadata"
      }

      // Only shape that has sufficient capacity.
      shape = "BM.Optimized3.36"

      source_details {
        boot_volume_size_in_gbs = "55"
        image_id                = var.bm_image_id
        source_type             = "image"
      }
    }
  }

  source = "NONE"
}

resource "oci_core_cluster_network" "test_cluster_network" {
  compartment_id = var.compartment_ocid
  display_name   = "hpc-cluster-network"

  instance_pools {
    display_name              = "hpc-cluster-network"
    instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id
    size                      = "1"
  }

  cluster_configuration{
    hpc_island_id             = var.hpc_island_id
  }

  placement_configuration {
    // Since sufficient capacity available only in AD3
    availability_domain       = data.oci_identity_availability_domains.test_availability_domains.availability_domains[2].name
    primary_subnet_id         = oci_core_subnet.test_subnet.id
    placement_constraint      = "PACKED_DISTRIBUTION_MULTI_BLOCK"

  }
}


