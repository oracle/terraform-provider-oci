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

variable "compartment_id" {
}

variable "compartment_id_for_update" {
}

variable "instance_type" {
  default = "STANDARDX"
}

variable "integration_instance_idcs_access_token" {
  default = ""
}

variable "integration_instance_consumption_model" {
  default = "UCM"
}

variable allow_listed_http_vcn {
  default = ""
}

variable certificate_secret_id {
  default = ""
}

variable domain_id {
  default = ""
}

variable subnet_id {
  default = ""
}

variable nsg_id {
  default = ""
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "random_integer" "seq" {
  min = 1000
  max = 1100
}

resource "oci_integration_integration_instance" "test_integration_instance" {
  #Required
  compartment_id            = var.compartment_id
  integration_instance_type = var.instance_type
  shape                     = "DEVELOPMENT"
  display_name              = "instance-created-via-tf-${random_integer.seq.result}"
  is_byol                   = "false"
  message_packs             = "1"
  domain_id                 = var.domain_id

  lifecycle {
    ignore_changes = [
      system_tags,
    ]
  }
  # idcs_at                 = var.integration_instance_idcs_access_token
  #Optional
  # custom_endpoint {
  #   hostname = "xyz.toronto.oicg3dev.ohaiarch.cloud"
  #   dns_zone_name = "toronto.oicg3dev.ohaiarch.cloud"
  # }
# For stand / enterprise type only
#  consumption_model = "${var.integration_instance_consumption_model}"
#  custom_endpoint {
#    certificate_secret_id = var.certificate_secret_id
#    hostname = "hostname.com"
#  }
#  freeform_tags = {
#    "bar-key" = "value"
#  }

#  is_file_server_enabled = true
#  is_visual_builder_enabled = true
#  state                  = "ACTIVE"
}

# home region
provider "oci" {
  alias = "phx"
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = "us-phoenix-1"
}

resource "oci_identity_dynamic_group" "DG_managed_custom_endpoint" {
  provider = oci.phx

  compartment_id = var.tenancy_ocid
  name = "DG_${oci_integration_integration_instance.test_integration_instance.display_name}"
  description = "DG for Oracle Managed Custom Endpoint"
  matching_rule = "any { resource.id = '${oci_integration_integration_instance.test_integration_instance.idcs_info[0].idcs_app_name}' }"
}

resource "oci_identity_policy" "policy_managed_custom_endpoint" {
  provider = oci.phx

  #Required
  compartment_id = var.tenancy_ocid
  description = "Policy for Oracle Managed Custom Endpoint"
  name = "policy_${oci_integration_integration_instance.test_integration_instance.display_name}"
  statements = [
    "ENDORSE any-user TO MANAGE certificate-authority-family IN any-tenancy",
    "Allow dynamic-group ${oci_identity_dynamic_group.DG_managed_custom_endpoint.name} to manage dns-zones in tenancy",
    "Allow dynamic-group ${oci_identity_dynamic_group.DG_managed_custom_endpoint.name} to manage dns-records in tenancy",
  ]
}

data "oci_integration_integration_instances" "test_integration_instances" {
  #Required
  compartment_id = var.compartment_id

  display_name = "instance-created-via-tf"
  state        = "Active"
}

data "oci_integration_integration_instance" "test_integration_instance" {
  #Required
  integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
}

# resource "oci_integration_integration_instance" "test_integration_instance_idcs" {
#   #Required
#   compartment_id            = var.compartment_id
#   display_name              = "instance4643_idcs"
#   integration_instance_type = "STANDARDX"
#   shape                     = "DEVELOPMENT"
#   # shape                     = "PRODUCTION"
#   is_byol                   = "false"
#   message_packs             = "10"
#   idcs_at                   = var.integration_instance_idcs_access_token
# }

resource "time_sleep" "wait" {
  depends_on = [oci_identity_policy.policy_managed_custom_endpoint]

  create_duration = "180s"
}

resource "oci_integration_oracle_managed_custom_endpoint" "integretion_custom_endpoint" {
  integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
  hostname = "${replace(oci_integration_integration_instance.test_integration_instance.display_name, "-", "")}.toronto.oicg3dev.ohaiarch.cloud"
  dns_zone_name = "toronto.oicg3dev.ohaiarch.cloud"
  depends_on = [time_sleep.wait]
}

resource "oci_integration_private_endpoint_outbound_connection" "integration_private_endpoint" {
  integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
  nsg_ids = [var.nsg_id]
  subnet_id = var.subnet_id

  depends_on = [
    oci_integration_oracle_managed_custom_endpoint.integretion_custom_endpoint
  ]
# resource "oci_integration_integration_instance" "test_integration_instance_idcs" {
#   #Required
#   compartment_id            = var.compartment_id
#   display_name              = "instance4643_idcs"
#   integration_instance_type = "STANDARDX"
#   shape                     = "DEVELOPMENT"
#   # shape                     = "PRODUCTION"
#   is_byol                   = "false"
#   message_packs             = "10"
#   idcs_at                   = var.integration_instance_idcs_access_token
# }

resource "oci_integration_private_endpoint_outbound_connection" "integration_private_endpoint" {
  integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
  nsg_ids = [var.nsg_id]
  subnet_id = var.subnet_id
}

resource "oci_integration_integration_instance" "test_integration_instance_with_dr" {
  #Required
  compartment_id            = var.compartment_id
  integration_instance_type = "STANDARDX"
  shape                     = "DEVELOPMENT"
  display_name              = "DR"
  is_byol                   = "false"
  message_packs             = "1"
  domain_id                 = var.domain_id
  is_disaster_recovery_enabled = "true"
  lifecycle {
    ignore_changes = ["system_tags"]
  }
}
