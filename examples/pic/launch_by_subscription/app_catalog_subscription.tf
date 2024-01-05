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

variable "ssh_public_key" {
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

/*
 * This example file shows how to configure the oci provider to target a single region.
 */
provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_vcn" "pic_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "PICVcn"
  dns_label      = "PICVcn"
}

resource "oci_core_subnet" "pic_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "PICSubnet"
  dns_label           = "PICSubnet"
  security_list_ids   = [oci_core_vcn.pic_vcn.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.pic_vcn.id
  route_table_id      = oci_core_route_table.pic_rt.id
  dhcp_options_id     = oci_core_vcn.pic_vcn.default_dhcp_options_id
}

resource "oci_core_internet_gateway" "pic_ig" {
  compartment_id = var.compartment_ocid
  display_name   = "PICIG"
  vcn_id         = oci_core_vcn.pic_vcn.id
}

resource "oci_core_route_table" "pic_rt" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.pic_vcn.id
  display_name   = "PICRT"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.pic_ig.id
  }
}

data "oci_core_app_catalog_listings" "test_app_catalog_listings" {
  /*filter {
    name   = "publisher_name"
    values = ["Oracle CCE Image Management Pipeline"]
  }*/
}

data "oci_core_app_catalog_listing_resource_versions" "test_app_catalog_listing_resource_versions" {
  #Required
  listing_id = data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings[0]["listing_id"]
}

resource "oci_core_app_catalog_listing_resource_version_agreement" "test_app_catalog_listing_resource_version_agreement" {
  #Required
  listing_id               = data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0]["listing_id"]
  listing_resource_version = data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0]["listing_resource_version"]
}

resource "oci_core_app_catalog_subscription" "test_app_catalog_subscription" {
  compartment_id           = var.compartment_ocid
  eula_link                = oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.eula_link
  listing_id               = oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.listing_id
  listing_resource_version = oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.listing_resource_version
  oracle_terms_of_use_link = oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.oracle_terms_of_use_link
  signature                = oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.signature
  time_retrieved           = oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.time_retrieved

  timeouts {
    create = "20m"
  }
}

resource "oci_core_instance" "pic_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "pic_instance"
  shape               = var.instance_shape

  create_vnic_details {
    subnet_id        = oci_core_subnet.pic_subnet.id
    display_name     = "picprimaryvnic"
    assign_public_ip = true
    hostname_label   = "PICInstance"
  }

  source_details {
    source_type = "image"
    source_id   = data.oci_core_app_catalog_subscriptions.test_app_catalog_subscriptions.app_catalog_subscriptions[0]["listing_resource_id"]
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
  }

  timeouts {
    create = "60m"
  }
}

data "oci_core_app_catalog_subscriptions" "test_app_catalog_subscriptions" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  listing_id = oci_core_app_catalog_subscription.test_app_catalog_subscription.listing_id

  filter {
    name   = "listing_resource_version"
    values = [oci_core_app_catalog_subscription.test_app_catalog_subscription.listing_resource_version]
  }
}

output "subscriptions" {
  value = [data.oci_core_app_catalog_subscriptions.test_app_catalog_subscriptions.app_catalog_subscriptions]
}

