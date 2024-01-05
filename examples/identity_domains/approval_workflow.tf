// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "approval_workflow_approval_workflow_filter" {
  default = ""
}

variable "approval_workflow_approval_workflow_steps_type" {
  default = "type"
}

variable "approval_workflow_approval_workflow_steps_value" {
  default = "value"
}

variable "approval_workflow_authorization" {
  default = "authorization"
}

variable "approval_workflow_compartment_ocid" {
  default = "compartmentOcid"
}

variable "approval_workflow_description" {
  default = "description"
}

variable "approval_workflow_domain_ocid" {
  default = "domainOcid"
}

variable "approval_workflow_max_duration_unit" {
  default = "MONTH"
}

variable "approval_workflow_max_duration_value" {
  default = 1
}

variable "approval_workflow_name" {
  default = "name"
}

variable "approval_workflow_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "approval_workflow_start_index" {
  default = 1
}

variable "approval_workflow_tags_key" {
  default = "key"
}

variable "approval_workflow_tags_value" {
  default = "value"
}

variable "approval_workflow_tenancy_ocid" {
  default = "tenancyOcid"
}

resource "oci_identity_domains_user" "test_approval_workflow_user" {
  # Required
  emails {
    value = "value@email.com"
    type = "work"
    primary = "true"
  }
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  name {
    family_name = "testApprovalWorkflowFamilyName"
  }
  schemas = ["urn:ietf:params:scim:schemas:core:2.0:User"]
  user_name = "testApprovalWorkflowUserName"
  lifecycle {
    ignore_changes = ["urnietfparamsscimschemasoracleidcsextension_oci_tags[0].defined_tags", "emails", "schemas"]
  }
}

resource "oci_identity_domains_approval_workflow_step" "test_approval_workflow_approval_workflow_step" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  order         = "1"
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:ApprovalWorkflowStep"]
  type          = "regular"
  approvers {
    type  = "User"
    value = oci_identity_domains_user.test_approval_workflow_user.id
  }
}

resource "oci_identity_domains_approval_workflow" "test_approval_workflow" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  max_duration {
    #Required
    unit  = var.approval_workflow_max_duration_unit
    value = var.approval_workflow_max_duration_value
  }
  name    = var.approval_workflow_name
  schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:ApprovalWorkflow"]

  #Optional
  approval_workflow_steps {
    #Required
    type  = oci_identity_domains_approval_workflow_step.test_approval_workflow_approval_workflow_step.type
    value = oci_identity_domains_approval_workflow_step.test_approval_workflow_approval_workflow_step.id
  }
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.approval_workflow_authorization
  description    = var.approval_workflow_description
  #resource_type_schema_version = var.approval_workflow_resource_type_schema_version
  tags {
    #Required
    key   = var.approval_workflow_tags_key
    value = var.approval_workflow_tags_value
  }
}

data "oci_identity_domains_approval_workflows" "test_approval_workflows" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  approval_workflow_filter = var.approval_workflow_approval_workflow_filter
  attribute_sets           = ["all"]
  attributes               = ""
  authorization            = var.approval_workflow_authorization
  #resource_type_schema_version = var.approval_workflow_resource_type_schema_version
  start_index = var.approval_workflow_start_index
}

