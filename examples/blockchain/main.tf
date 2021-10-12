// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file
variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "idcs_access_token" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_blockchain_blockchain_platform" "test_blockchain_platform" {
  #Required
  compartment_id = var.compartment_ocid
  compute_shape  = "ENTERPRISE_MEDIUM"
  display_name   = "displayname3"
  platform_version   = "Hyperledger Fabric v2.2.4"
  platform_role  = "FOUNDER"
  idcs_access_token = var.idcs_access_token
}

data "oci_blockchain_blockchain_platform" "test_blockchain_platform" {
  #Required
  blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id
}

data "oci_blockchain_blockchain_platforms" "test_blockchain_platforms" {
  #Required
  compartment_id = var.compartment_ocid
}

resource "oci_blockchain_osn" "test_osn" {
  #Required
  ad                     = "AD1"
  blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id

  ocpu_allocation_param {
    #Required
    ocpu_allocation_number = "0.0"
  }
}

data "oci_blockchain_osn" "test_osn" {
  #Required
  osn_id                 = oci_blockchain_osn.test_osn.id
  blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id
}

resource "oci_blockchain_peer" "test_peer" {
  #Required
  ad                     = "AD1"
  blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id

  ocpu_allocation_param {
    #Required
    ocpu_allocation_number = "0.5"
  }

  role = "MEMBER"
  
  # This depends on is added to make the peer and osn creation sequential to avoid conflicts
  depends_on = [oci_blockchain_osn.test_osn]
}

data "oci_blockchain_peer" "test_peer" {
  #Required
  peer_id                = oci_blockchain_peer.test_peer.id
  blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id
}

