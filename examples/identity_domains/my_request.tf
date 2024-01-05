// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_request_my_request_count" {
  default = 10
}

variable "my_request_my_request_filter" {
  default = ""
}

variable "my_request_authorization" {
  default = "authorization"
}

variable "my_request_justification" {
  default = "justification"
}

variable "my_request_requesting_type" {
  default = "Group"
}

variable "my_request_requesting_value" {
  default = "value"
}

variable "my_request_requestor_value" {
  default = "value"
}

variable "my_request_start_index" {
  default = 1
}

variable "my_request_status" {
  default = "status"
}

variable "my_request_tags_key" {
  default = "key"
}

variable "my_request_tags_value" {
  default = "value"
}


resource "oci_identity_domains_group" "group_to_request" {
  display_name = "groupToRequest"
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:core:2.0:Group"]
  attribute_sets = ["all"]
  urnietfparamsscimschemasoracleidcsextensionrequestable_group {
    #Optional
    requestable = true
  }
  lifecycle {
    ignore_changes = [schemas]
  }
}

resource "oci_identity_domains_my_request" "test_my_request" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  justification = var.my_request_justification
  requesting {
    #Required
    type  = var.my_request_requesting_type
    value = oci_identity_domains_group.group_to_request.id
  }
  schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:Request"]

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.my_request_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.my_request_resource_type_schema_version
  tags {
    #Required
    key   = var.my_request_tags_key
    value = var.my_request_tags_value
  }

  lifecycle {
    ignore_changes = [schemas, tags]
  }
}

data "oci_identity_domains_my_requests" "test_my_requests" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  my_request_count             = var.my_request_my_request_count
  my_request_filter            = var.my_request_my_request_filter
  attribute_sets               = []
  attributes                   = ""
  authorization                = var.my_request_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.my_request_resource_type_schema_version
  start_index                  = var.my_request_start_index
}

