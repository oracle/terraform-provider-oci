// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "approval_workflow_step_approval_workflow_step_count" {
  default = 10
}

variable "approval_workflow_step_approval_workflow_step_filter" {
  default = ""
}

variable "approval_workflow_step_approvers_ref" {
  default = "ref"
}

variable "approval_workflow_step_approvers_type" {
  default = "User"
}

variable "approval_workflow_step_approvers_value" {
  default = "value"
}

variable "approval_workflow_step_approvers_expressions" {
  default = []
}

variable "approval_workflow_step_authorization" {
  default = "authorization"
}

variable "approval_workflow_step_compartment_ocid" {
  default = "compartmentOcid"
}

variable "approval_workflow_step_delete_in_progress" {
  default = false
}

variable "approval_workflow_step_domain_ocid" {
  default = "domainOcid"
}

variable "approval_workflow_step_minimum_approvals" {
  default = 10
}

variable "approval_workflow_step_order" {
  default = 10
}

variable "approval_workflow_step_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "approval_workflow_step_schemas" {
  default = []
}

variable "approval_workflow_step_start_index" {
  default = 1
}

variable "approval_workflow_step_tags_key" {
  default = "key"
}

variable "approval_workflow_step_tags_value" {
  default = "value"
}

variable "approval_workflow_step_tenancy_ocid" {
  default = "tenancyOcid"
}

variable "approval_workflow_step_type" {
  default = "escalation"
}

resource "oci_identity_domains_user" "test_approval_workflow_step_user" {
  # Required
  emails {
    value = "value@email.com"
    type = "work"
    primary = "true"
  }
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  name {
    family_name = "testApprovalWorkflowStepFamilyName"
  }
  schemas = ["urn:ietf:params:scim:schemas:core:2.0:User"]
  user_name = "testApprovalWorkflowStepUserName"
  lifecycle {
    ignore_changes = ["urnietfparamsscimschemasoracleidcsextension_oci_tags[0].defined_tags", "emails", "schemas"]
  }
}

resource "oci_identity_domains_approval_workflow_step" "test_approval_workflow_step" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  order         = var.approval_workflow_step_order
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:ApprovalWorkflowStep"]
  type          = var.approval_workflow_step_type

  #Optional
  approvers {
    #Required
    type  = var.approval_workflow_step_approvers_type
    value = oci_identity_domains_user.test_approval_workflow_step_user.id
  }
  approvers_expressions = var.approval_workflow_step_approvers_expressions
  attribute_sets        = ["all"]
  attributes            = ""
  authorization         = var.approval_workflow_step_authorization
  minimum_approvals     = var.approval_workflow_step_minimum_approvals
  #resource_type_schema_version = var.approval_workflow_step_resource_type_schema_version
  tags {
    #Required
    key   = var.approval_workflow_step_tags_key
    value = var.approval_workflow_step_tags_value
  }
}

data "oci_identity_domains_approval_workflow_steps" "test_approval_workflow_steps" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  approval_workflow_step_count  = var.approval_workflow_step_approval_workflow_step_count
  approval_workflow_step_filter = var.approval_workflow_step_approval_workflow_step_filter
  attribute_sets                = ["all"]
  attributes                    = ""
  authorization                 = var.approval_workflow_step_authorization
  #resource_type_schema_version  = var.approval_workflow_step_resource_type_schema_version
  start_index = var.approval_workflow_step_start_index
}

