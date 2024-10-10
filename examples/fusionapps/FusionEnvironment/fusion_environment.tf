// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "fusion_environment_additional_language_packs" {
  default = []
}

variable "fusion_environment_create_fusion_environment_admin_user_details_email_address" {
  default = "JohnSmith@example.com"
}

variable "fusion_environment_create_fusion_environment_admin_user_details_first_name" {
  default = "firstName"
}

variable "fusion_environment_create_fusion_environment_admin_user_details_last_name" {
  default = "lastName"
}

variable "fusion_environment_create_fusion_environment_admin_user_details_password" {
  default = "BEstrO0ng_#11"
}

variable "fusion_environment_create_fusion_environment_admin_user_details_username" {
  default = "username_test"
}

variable "fusion_environment_defined_tags_value" {
  default = "value"
}

variable "fusion_environment_display_name" {
  default = "displayName"
}

variable "fusion_environment_dns_prefix" {
  default = "dnsPrefix"
}

variable "fusion_environment_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "fusion_environment_fusion_environment_type" {
  default = "TEST"
}

variable "fusion_environment_maintenance_policy_environment_maintenance_override" {
  default = "PROD"
}

variable "fusion_environment_maintenance_policy_monthly_patching_override" {
  default = "ENABLED"
}

variable "fusion_environment_rules_action" {
  default = "ALLOW"
}

variable "fusion_environment_rules_conditions_attribute_name" {
  default = "SOURCE_IP_ADDRESS"
}

variable "fusion_environment_rules_conditions_attribute_value" {
  default = "208.128.0.0/10"
}

variable "fusion_environment_rules_description" {
  default = "description"
}

variable "fusion_environment_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_fusion_apps_fusion_environment" "test_fusion_environment" {
  #Required
  compartment_id = var.compartment_id
  create_fusion_environment_admin_user_details {
    #Required
    email_address = var.fusion_environment_create_fusion_environment_admin_user_details_email_address
    first_name    = var.fusion_environment_create_fusion_environment_admin_user_details_first_name
    last_name     = var.fusion_environment_create_fusion_environment_admin_user_details_last_name
    username      = var.fusion_environment_create_fusion_environment_admin_user_details_username
    #Optional
    password      = var.fusion_environment_create_fusion_environment_admin_user_details_password
  }
  display_name                 = var.fusion_environment_display_name
  fusion_environment_family_id = oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id
  fusion_environment_type      = var.fusion_environment_fusion_environment_type

  #Optional
  additional_language_packs = var.fusion_environment_additional_language_packs
  defined_tags              = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.fusion_environment_defined_tags_value)
  dns_prefix                = var.fusion_environment_dns_prefix
  freeform_tags             = var.fusion_environment_freeform_tags
  # This field is related to the subscription you have
  #kms_key_id                = oci_kms_key.test_key.id
  maintenance_policy {

    #Optional
    environment_maintenance_override = var.fusion_environment_maintenance_policy_environment_maintenance_override
    monthly_patching_override        = var.fusion_environment_maintenance_policy_monthly_patching_override
  }
  rules {
    #Required
    action = var.fusion_environment_rules_action
    conditions {
      #Required
      attribute_name  = var.fusion_environment_rules_conditions_attribute_name
      attribute_value = var.fusion_environment_rules_conditions_attribute_value
    }

    #Optional
    description = var.fusion_environment_rules_description
  }
}

data "oci_fusion_apps_fusion_environments" "test_fusion_environments" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name                 = var.fusion_environment_display_name
  fusion_environment_family_id = oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id
  state                        = var.fusion_environment_state
}
