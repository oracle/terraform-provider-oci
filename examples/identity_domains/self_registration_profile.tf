// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "self_registration_profile_self_registration_profile_count" {
  default = 10
}

variable "self_registration_profile_self_registration_profile_filter" {
  default = ""
}

variable "self_registration_profile_activation_email_required" {
  default = false
}

variable "self_registration_profile_active" {
  default = false
}

variable "self_registration_profile_after_submit_text_default" {
  default = true
}

variable "self_registration_profile_after_submit_text_locale" {
  default = "en-US"
}

variable "self_registration_profile_after_submit_text_value" {
  default = "value"
}

variable "self_registration_profile_allowed_email_domains" {
  default = []
}

variable "self_registration_profile_attribute_sets" {
  default = []
}

variable "self_registration_profile_attributes" {
  default = ""
}

variable "self_registration_profile_authorization" {
  default = "authorization"
}

variable "self_registration_profile_compartment_ocid" {
  default = "compartmentOcid"
}

variable "self_registration_profile_consent_text_default" {
  default = true
}

variable "self_registration_profile_consent_text_locale" {
  default = "en-US"
}

variable "self_registration_profile_consent_text_value" {
  default = "value"
}

variable "self_registration_profile_consent_text_present" {
  default = false
}

variable "self_registration_profile_default_groups_display" {
  default = "display"
}

variable "self_registration_profile_default_groups_ref" {
  default = "ref"
}

variable "self_registration_profile_default_groups_value" {
  default = "value"
}

variable "self_registration_profile_delete_in_progress" {
  default = false
}

variable "self_registration_profile_disallowed_email_domains" {
  default = []
}

variable "self_registration_profile_display_name_default" {
  default = true
}

variable "self_registration_profile_display_name_locale" {
  default = "en-US"
}

variable "self_registration_profile_display_name_value" {
  default = "value"
}

variable "self_registration_profile_domain_ocid" {
  default = "domainOcid"
}

variable "self_registration_profile_email_template_display" {
  default = "display"
}

variable "self_registration_profile_email_template_ref" {
  default = "ref"
}

variable "self_registration_profile_email_template_value" {
  default = "MeRegisterVerifyEmail"
}

variable "self_registration_profile_footer_logo" {
  default = "footerLogo"
}

variable "self_registration_profile_footer_text_default" {
  default = true
}

variable "self_registration_profile_footer_text_locale" {
  default = "en-US"
}

variable "self_registration_profile_footer_text_value" {
  default = "value"
}

variable "self_registration_profile_header_logo" {
  default = "headerLogo"
}

variable "self_registration_profile_header_text_default" {
  default = true
}

variable "self_registration_profile_header_text_locale" {
  default = "en-US"
}

variable "self_registration_profile_header_text_value" {
  default = "value"
}

variable "self_registration_profile_id" {
  default = "id"
}

variable "self_registration_profile_meta_created" {
  default = "created"
}

variable "self_registration_profile_meta_last_modified" {
  default = "lastModified"
}

variable "self_registration_profile_meta_location" {
  default = "location"
}

variable "self_registration_profile_meta_resource_type" {
  default = "resourceType"
}

variable "self_registration_profile_meta_version" {
  default = "version"
}

variable "self_registration_profile_name" {
  default = "name"
}

variable "self_registration_profile_number_of_days_redirect_url_is_valid" {
  default = 10
}

variable "self_registration_profile_ocid" {
  default = "ocid"
}

variable "self_registration_profile_redirect_url" {
  default = "https://www.oracle.com"
}

variable "self_registration_profile_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "self_registration_profile_show_on_login_page" {
  default = false
}

variable "self_registration_profile_start_index" {
  default = 10
}

variable "self_registration_profile_tags_key" {
  default = "key"
}

variable "self_registration_profile_tags_value" {
  default = "value"
}

variable "self_registration_profile_tenancy_ocid" {
  default = "tenancyOcid"
}

variable "self_registration_profile_user_attributes_deletable" {
  default = false
}

variable "self_registration_profile_user_attributes_fully_qualified_attribute_name" {
  default = "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber"
}

variable "self_registration_profile_user_attributes_metadata" {
  default = "metadata"
}

variable "self_registration_profile_user_attributes_seq_number" {
  default = 10
}

variable "self_registration_profile_user_attributes_value" {
  default = "employeeNumber"
}

resource "oci_identity_domains_group" "test_self_registration_profile_group" {
  #Required
  display_name  = "groupDisplayName"
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:core:2.0:Group"]

  lifecycle {
    ignore_changes = [schemas]
  }
}

resource "oci_identity_domains_self_registration_profile" "test_self_registration_profile" {
  #Required
  activation_email_required = var.self_registration_profile_activation_email_required
  consent_text_present      = var.self_registration_profile_consent_text_present
  display_name {
    #Required
    locale = var.self_registration_profile_display_name_locale
    value  = var.self_registration_profile_display_name_value

    #Optional
    default = var.self_registration_profile_display_name_default
  }
  email_template {
    #Required
    value = var.self_registration_profile_email_template_value

    #Optional
    #display = var.self_registration_profile_email_template_display
    #ref     = var.self_registration_profile_email_template_ref
  }
  idcs_endpoint                        = data.oci_identity_domain.test_domain.url
  name                                 = var.self_registration_profile_name
  number_of_days_redirect_url_is_valid = var.self_registration_profile_number_of_days_redirect_url_is_valid
  redirect_url                         = var.self_registration_profile_redirect_url
  schemas                              = ["urn:ietf:params:scim:schemas:oracle:idcs:SelfRegistrationProfile"]
  show_on_login_page                   = var.self_registration_profile_show_on_login_page

  #Optional
  active = var.self_registration_profile_active
  after_submit_text {
    #Required
    locale = var.self_registration_profile_after_submit_text_locale
    value  = var.self_registration_profile_after_submit_text_value

    #Optional
    default = var.self_registration_profile_after_submit_text_default
  }
  allowed_email_domains = var.self_registration_profile_allowed_email_domains
  attribute_sets        = ["all"]
  attributes            = ""
  authorization         = var.self_registration_profile_authorization
  consent_text {
    #Required
    locale = var.self_registration_profile_consent_text_locale
    value  = var.self_registration_profile_consent_text_value

    #Optional
    default = var.self_registration_profile_consent_text_default
  }
  default_groups {
    #Required
    value = oci_identity_domains_group.test_self_registration_profile_group.id
  }
  disallowed_email_domains = var.self_registration_profile_disallowed_email_domains
  footer_logo              = var.self_registration_profile_footer_logo
  footer_text {
    #Required
    locale = var.self_registration_profile_footer_text_locale
    value  = var.self_registration_profile_footer_text_value

    #Optional
    default = var.self_registration_profile_footer_text_default
  }
  header_logo = var.self_registration_profile_header_logo
  header_text {
    #Required
    locale = var.self_registration_profile_header_text_locale
    value  = var.self_registration_profile_header_text_value

    #Optional
    default = var.self_registration_profile_header_text_default
  }
#use the latest if not provided
# resource_type_schema_version = var.self_registration_profile_resource_type_schema_version
  tags {
    #Required
    key   = var.self_registration_profile_tags_key
    value = var.self_registration_profile_tags_value
  }
  user_attributes {
    #Required
    seq_number = var.self_registration_profile_user_attributes_seq_number
    value      = var.self_registration_profile_user_attributes_value

    #Optional
    fully_qualified_attribute_name = var.self_registration_profile_user_attributes_fully_qualified_attribute_name
  }
}

data "oci_identity_domains_self_registration_profiles" "test_self_registration_profiles" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  self_registration_profile_count  = var.self_registration_profile_self_registration_profile_count
  self_registration_profile_filter = var.self_registration_profile_self_registration_profile_filter
  attribute_sets                   = var.self_registration_profile_attribute_sets
  attributes                       = var.self_registration_profile_attributes
  authorization                    = var.self_registration_profile_authorization
#use the latest if not provided
# resource_type_schema_version     = var.self_registration_profile_resource_type_schema_version
  start_index                      = var.self_registration_profile_start_index
}
