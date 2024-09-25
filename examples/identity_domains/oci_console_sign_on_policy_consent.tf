// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "oci_console_sign_on_policy_consent_oci_console_sign_on_policy_consent_count" {
  default = 10
}

variable "oci_console_sign_on_policy_consent_oci_console_sign_on_policy_consent_filter" {
  default = ""
}

variable "oci_console_sign_on_policy_consent_authorization" {
  default = "authorization"
}

variable "oci_console_sign_on_policy_consent_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "oci_console_sign_on_policy_consent_start_index" {
  default = 1
}


data "oci_identity_domains_oci_console_sign_on_policy_consents" "test_oci_console_sign_on_policy_consents" {

  idcs_endpoint = data.oci_identity_domain.test_domain.url
  #Optional
  oci_console_sign_on_policy_consent_count  = var.oci_console_sign_on_policy_consent_oci_console_sign_on_policy_consent_count
  oci_console_sign_on_policy_consent_filter = var.oci_console_sign_on_policy_consent_oci_console_sign_on_policy_consent_filter
  authorization                             = var.oci_console_sign_on_policy_consent_authorization
  #resource_type_schema_version              = var.oci_console_sign_on_policy_consent_resource_type_schema_version
  start_index                               = var.oci_console_sign_on_policy_consent_start_index
}

data "oci_identity_domains_oci_console_sign_on_policy_consent" "test_oci_console_sign_on_policy_consent" {

  idcs_endpoint = data.oci_identity_domain.test_domain.url
  oci_console_sign_on_policy_consent_id = data.oci_identity_domains_oci_console_sign_on_policy_consents.test_oci_console_sign_on_policy_consents.resources.0.id
}

output "consent_resource" {
  value = {
    id = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.id
    change_type = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.change_type
    client_ip = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.client_ip
    reason = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.reason
    justification = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.justification
    notification_recipients = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.notification_recipients
    consent_signed_by = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.consent_signed_by
    time_consent_signed = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.time_consent_signed
    modified_resource = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.modified_resource
    policy_resource = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.policy_resource
    time_consent_signed = data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent.time_consent_signed
  }
}