// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "approval_workflow_assignment_approval_workflow_assignment_count" {
  default = 10
}

variable "approval_workflow_assignment_approval_workflow_assignment_filter" {
  default = ""
}

variable "approval_workflow_assignment_approval_workflow_ocid" {
  default = ""
}

variable "approval_workflow_assignment_approval_workflow_type" {
  default = "ApprovalWorkflow"
}

variable "approval_workflow_assignment_assigned_to_ocid" {
  default = ""
}

variable "approval_workflow_assignment_assigned_to_type" {
  default = "Group"
}

variable "approval_workflow_assignment_assigned_to_value" {
  default = "value"
}

variable "approval_workflow_assignment_assignment_type" {
  default = "MEMBERSHIP"
}

variable "approval_workflow_assignment_attribute_sets" {
  default = []
}

variable "approval_workflow_assignment_authorization" {
  default = "authorization"
}

variable "approval_workflow_assignment_ocid" {
  default = ""
}

variable "approval_workflow_assignment_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "approval_workflow_assignment_start_index" {
  default = 1
}

variable "approval_workflow_assignment_tags_key" {
  default = "key"
}

variable "approval_workflow_assignment_tags_value" {
  default = "value"
}

resource "oci_identity_domains_approval_workflow" "test_approval_workflow_assignment_approval_workflow" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  max_duration {
    #Required
    unit  = "MONTH"
    value = "1"
  }
  name    = "test"
  schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:ApprovalWorkflow"]
}

resource "oci_identity_domains_approval_workflow_assignment" "test_approval_workflow_assignment" {
  #Required
  approval_workflow {
    #Required
    type  = var.approval_workflow_assignment_approval_workflow_type
    value = oci_identity_domains_approval_workflow.test_approval_workflow_assignment_approval_workflow.id
  }
  assigned_to {
    #Required
    type  = var.approval_workflow_assignment_assigned_to_type
    value = var.approval_workflow_assignment_assigned_to_value
  }
  assignment_type = var.approval_workflow_assignment_assignment_type
  idcs_endpoint   = data.oci_identity_domain.test_domain.url
  schemas         = ["urn:ietf:params:scim:schemas:oracle:idcs:ApprovalWorkflowAssignment"]

  #Optional
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.approval_workflow_assignment_authorization
#  resource_type_schema_version = var.approval_workflow_assignment_resource_type_schema_version
  tags {
    #Required
    key   = var.approval_workflow_assignment_tags_key
    value = var.approval_workflow_assignment_tags_value
  }
}

data "oci_identity_domains_approval_workflow_assignments" "test_approval_workflow_assignments" {
  #Required
  idcs_endpoint   = data.oci_identity_domain.test_domain.url

  #Optional
  approval_workflow_assignment_count  = var.approval_workflow_assignment_approval_workflow_assignment_count
  approval_workflow_assignment_filter = var.approval_workflow_assignment_approval_workflow_assignment_filter
  attribute_sets                      = ["all"]
  attributes                          = ""
  authorization                       = var.approval_workflow_assignment_authorization
#  resource_type_schema_version        = var.approval_workflow_assignment_resource_type_schema_version
  start_index                         = var.approval_workflow_assignment_start_index
}

