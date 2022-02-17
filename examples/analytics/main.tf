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

variable "kms_key_id" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_analytics_analytics_instance" "test_oce_instance_public" {
  compartment_id     = var.compartment_ocid
  description        = "OAC instance"
  email_notification = var.email_notification
  feature_set        = "ENTERPRISE_ANALYTICS"
  license_type       = "LICENSE_INCLUDED"

  capacity {
    capacity_type  = "OLPU_COUNT"
    capacity_value = 2
  }

  name = "testoacinstance1"
  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }
  state             = "ACTIVE"
  idcs_access_token = var.idcs_access_token

  # Optional
  network_endpoint_details {
    #Required
    network_endpoint_type = "PUBLIC"

    #Optional
    whitelisted_ips = [oci_core_vcn.test_vcn.cidr_block]

    whitelisted_vcns {
      #Optional
      id              = oci_core_vcn.test_vcn.id
      whitelisted_ips = [oci_core_vcn.test_vcn.cidr_block]
    }
  }

  # Optionally specify a master encryption key for analytics metadata.
  kms_key_id = var.kms_key_id
}

# Create a private access channel for the instance
resource "oci_analytics_analytics_instance_private_access_channel" "test_private_access_channel" {
#Required
  analytics_instance_id = oci_analytics_analytics_instance.test_oce_instance_public.id
  display_name = "ExamplePAC"
  subnet_id = oci_core_subnet.test_subnet.id
  vcn_id = oci_core_vcn.test_vcn.id
  private_source_dns_zones {
    dns_zone = "examplecorp.com"
    description = "Example dns zone"
  }
}

# Create a vanity url for the instance
resource "oci_analytics_analytics_instance_vanity_url" "test_analytics_instances_vanity_url" {
  #Required
  analytics_instance_id = oci_analytics_analytics_instance.test_oce_instance_public.id
  ca_certificate        = file("/path/to/the/file/RootCA.crt")
  hosts                 = ["analyticsdev.mycompany.com"]
  private_key           = file("/path/to/the/file/analyticsdevCA.key")
  public_certificate    = file("/path/to/the/file/analyticsdevCA.crt")
  
  #Optional
  description = "description"
  # passphrase is required if the private key was created with a passphrase
  passphrase  = "passphrase"
}

