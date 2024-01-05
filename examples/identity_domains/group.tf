// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "group_group_count" {
  default = 10
}

variable "group_group_filter" {
  default = ""
}

variable "group_authorization" {
  default = "authorization"
}

variable "group_display_name" {
  default = "displayName"
}

variable "group_members_date_added" {
  default = "dateAdded"
}

variable "group_members_display" {
  default = "display"
}

variable "group_members_membership_ocid" {
  default = "membershipOcid"
}

variable "group_members_name" {
  default = "name"
}

variable "group_members_ocid" {
  default = "ocid"
}

variable "group_members_ref" {
  default = "ref"
}

variable "group_members_type" {
  default = "User"
}

variable "group_members_value" {
  default = "value"
}

variable "group_non_unique_display_name" {
  default = "nonUniqueDisplayName"
}

variable "group_start_index" {
  default = 1
}

variable "group_tags_key" {
  default = "key"
}

variable "group_tags_value" {
  default = "value"
}

variable "group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key" {
  default = "key"
}

variable "group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace" {
  default = "namespace"
}

variable "group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value" {
  default = "value"
}

variable "group_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key" {
  default = "freeformKey"
}

variable "group_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value" {
  default = "freeformValue"
}

variable "group_urnietfparamsscimschemasoracleidcsextensiondynamic_group_membership_rule" {
  default = "membershipRule"
}

variable "group_urnietfparamsscimschemasoracleidcsextensiondynamic_group_membership_type" {
  default = "static"
}

variable "group_urnietfparamsscimschemasoracleidcsextensiongroup_group_creation_mechanism" {
  default = "api"
}

variable "group_urnietfparamsscimschemasoracleidcsextensiongroup_group_description" {
  default = "description"
}

variable "group_urnietfparamsscimschemasoracleidcsextensiongroup_group_owners_display" {
  default = "display"
}

variable "group_urnietfparamsscimschemasoracleidcsextensiongroup_group_owners_type" {
  default = "User"
}

variable "group_urnietfparamsscimschemasoracleidcsextensiongroup_group_owners_value" {
  default = "value"
}

variable "group_urnietfparamsscimschemasoracleidcsextensionposix_group_gid_number" {
  default = 500
}

variable "group_urnietfparamsscimschemasoracleidcsextensionrequestable_group_requestable" {
  default = false
}


resource "oci_identity_domains_group" "test_group" {
  #Required
  display_name  = var.group_display_name
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:core:2.0:Group"]

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.group_authorization
  external_id    = "externalId"
  /* #provide user's id and/or ocid to add to this group
  members {
    #Required
    type  = var.group_members_type
    value = var.group_members_value

    #Optional
    ocid = var.group_members_ocid
  }
  */
  non_unique_display_name = var.group_non_unique_display_name
  tags {
    #Required
    key   = var.group_tags_key
    value = var.group_tags_value
  }
  urnietfparamsscimschemasoracleidcsextension_oci_tags {

    #Optional
    /* #create tagNamespace to use defined tags
    defined_tags {
      #Required
      key       = var.group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key
      namespace = var.group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace
      value     = var.group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value
    }
    */
    freeform_tags {
      #Required
      key   = var.group_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key
      value = var.group_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value
    }
  }
  urnietfparamsscimschemasoracleidcsextensiondynamic_group {

    #Optional
    #membership_rule can't be set for static groups
    # membership_rule = var.group_urnietfparamsscimschemasoracleidcsextensiondynamic_group_membership_rule
    membership_type = var.group_urnietfparamsscimschemasoracleidcsextensiondynamic_group_membership_type
  }
  urnietfparamsscimschemasoracleidcsextensiongroup_group {

    #Optional
    creation_mechanism = var.group_urnietfparamsscimschemasoracleidcsextensiongroup_group_creation_mechanism
    description        = var.group_urnietfparamsscimschemasoracleidcsextensiongroup_group_description
    /* #set value to id of user/app
    owners {
      #Required
      type  = var.group_urnietfparamsscimschemasoracleidcsextensiongroup_group_owners_type
      value = var.group_urnietfparamsscimschemasoracleidcsextensiongroup_group_owners_value
    }
    */
  }
  urnietfparamsscimschemasoracleidcsextensionposix_group {

    #Optional
    gid_number = var.group_urnietfparamsscimschemasoracleidcsextensionposix_group_gid_number
  }
  urnietfparamsscimschemasoracleidcsextensionrequestable_group {

    #Optional
    requestable = var.group_urnietfparamsscimschemasoracleidcsextensionrequestable_group_requestable
  }
  lifecycle {
    ignore_changes = [schemas]
  }
}

data "oci_identity_domains_groups" "test_groups" {
  #Required
  # provider the idcs url of domain
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  group_count    = var.group_group_count
  group_filter   = var.group_group_filter
  attribute_sets = []
  attributes     = ""
  authorization  = var.group_authorization
  start_index    = var.group_start_index
}

