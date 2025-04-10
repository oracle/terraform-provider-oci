// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "tf-vcn"
  dns_label      = "dnslabel"
}

resource "oci_core_internet_gateway" "test_network_entity" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "-tf-internet-gateway"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_ocid

  route_rules {
    cidr_block        = "0.0.0.0/0"
    network_entity_id = oci_core_internet_gateway.test_network_entity.id
  }

  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_network_security_group" "test_network_security_group" {
    compartment_id = var.compartment_ocid
    vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block                 = "10.0.0.0/16"
  compartment_id             = var.compartment_ocid
  dhcp_options_id            = oci_core_vcn.test_vcn.default_dhcp_options_id
  display_name               = "tf-subnet"
  dns_label                  = "dnslabel"
  prohibit_public_ip_on_vnic = "false"
  route_table_id             = oci_core_route_table.test_route_table.id
  security_list_ids          = [oci_core_vcn.test_vcn.default_security_list_id]
  vcn_id                     = oci_core_vcn.test_vcn.id
}

# Terraform will take 5 minutes after destroying an application due to a known service issue.
# please refer: https://docs.cloud.oracle.com/iaas/Content/Functions/Tasks/functionsdeleting.htm
resource "oci_functions_application" "test_application" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "example-application-test"
  subnet_ids     = [oci_core_subnet.test_subnet.id]

  #Optional
  config                     = var.config
  syslog_url                 = var.syslog_url
  network_security_group_ids = [oci_core_network_security_group.test_network_security_group.id]
  image_policy_config {
    #Required
    is_policy_enabled = var.application_image_policy_config_is_policy_enabled

    #Optional
    key_details {
      #Required
      kms_key_id = var.kms_key_ocid
    }
  }
  trace_config {
    domain_id  = var.application_trace_config.domain_id
    is_enabled = var.application_trace_config.is_enabled
  }
  shape = var.application_shape
}

data "oci_functions_applications" "test_applications" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = "example-application-test"
  id           = oci_functions_application.test_application.id
  state        = var.application_state
}

resource "oci_functions_function" "test_function" {
  #Required
  application_id = oci_functions_application.test_application.id
  display_name   = "example-function-test"
  image          = var.function_image
  memory_in_mbs  = "128"

  #Optional
  config             = var.config
  image_digest       = var.function_image_digest
  timeout_in_seconds = "30"
  trace_config {
    is_enabled = var.function_trace_config.is_enabled
  }
  is_dry_run         = var.dry_run

  provisioned_concurrency_config {
    strategy = "CONSTANT"
    count = 40
  }
}

data "oci_functions_pbf_listings" "test_listings" {
  #Optional
  name = var.pbf_listing_name
}

data "oci_functions_pbf_listing" "test_listing" {
  #Required
  pbf_listing_id = var.pbf_listing_id
}

data "oci_functions_pbf_listing_versions" "test_versions" {
  #Required
  pbf_listing_id = var.pbf_listing_id

  #Optional
  is_current_version = true
}

data "oci_functions_pbf_listing_version" "test_version" {
  #Required
  pbf_listing_version_id = var.pbf_listing_version_id
}

data "oci_functions_pbf_listing_triggers" "test_triggers" {
  #Optional
  name = var.pbf_trigger_name
}

resource "oci_functions_function" "test_pre_built_function" {
  application_id = oci_functions_application.test_application.id
  display_name = "example-pre-built-function"
  memory_in_mbs = "128"
  source_details {
    pbf_listing_id = var.pbf_listing_id
    source_type = "PRE_BUILT_FUNCTIONS"
  }
}

data "oci_functions_functions" "test_pre_built_functions" {
  #Required
  application_id = oci_functions_application.test_application.id

  #Optional
  display_name = "example-pre-built-function"
  id           = oci_functions_function.test_pre_built_function.id
  state        = "ACTIVE"
}

data "oci_functions_functions" "test_functions" {
  #Required
  application_id = oci_functions_application.test_application.id

  #Optional
  display_name = "example-function-test"
  id           = oci_functions_function.test_function.id
  state        = "ACTIVE"
}

resource "time_sleep" "wait_function_provisioning" {
  depends_on      = [oci_functions_function.test_function]

  create_duration = "5s"
}
