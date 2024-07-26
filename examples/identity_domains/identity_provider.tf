// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "identity_provider_identity_provider_count" {
  default = 10
}

variable "identity_provider_identity_provider_filter" {
  default = ""
}

variable "identity_provider_assertion_attribute" {
  default = "assertionAttribute"
}

variable "identity_provider_authn_request_binding" {
  default = "Redirect"
}

variable "identity_provider_authorization" {
  default = "authorization"
}

variable "identity_provider_correlation_policy_display" {
  default = "display"
}

variable "identity_provider_correlation_policy_ref" {
  default = "ref"
}

variable "identity_provider_correlation_policy_type" {
  default = "Policy"
}

variable "identity_provider_correlation_policy_value" {
  default = "value"
}

variable "identity_provider_description" {
  default = "description"
}

#provide the IDP encryption cert
variable "identity_provider_encryption_certificate" {
  default = ""
}

variable "identity_provider_icon_url" {
  default = "https://something.com/iconUrl.png"
}

variable "identity_provider_idp_sso_url" {
  default = "https://idpSsoUrl.com"
}

variable "identity_provider_include_signing_cert_in_signature" {
  default = false
}

variable "identity_provider_jit_user_prov_assigned_groups_display" {
  default = "display"
}

variable "identity_provider_jit_user_prov_assigned_groups_value" {
  default = "value"
}

variable "identity_provider_jit_user_prov_attribute_update_enabled" {
  default = false
}

variable "identity_provider_jit_user_prov_attributes_value" {
  default = "value"
}

variable "identity_provider_jit_user_prov_create_user_enabled" {
  default = false
}

variable "identity_provider_jit_user_prov_enabled" {
  default = false
}

variable "identity_provider_jit_user_prov_group_assertion_attribute_enabled" {
  default = false
}

variable "identity_provider_jit_user_prov_group_assignment_method" {
  default = "Overwrite"
}

variable "identity_provider_jit_user_prov_group_mapping_mode" {
  default = "implicit"
}

variable "identity_provider_jit_user_prov_group_mappings_idp_group" {
  default = "idpGroup"
}

variable "identity_provider_jit_user_prov_group_mappings_value" {
  default = "value"
}

variable "identity_provider_jit_user_prov_group_saml_attribute_name" {
  default = "jitUserProvGroupSAMLAttributeName"
}

variable "identity_provider_jit_user_prov_group_static_list_enabled" {
  default = false
}

variable "identity_provider_jit_user_prov_ignore_error_on_absent_groups" {
  default = false
}

variable "identity_provider_logout_binding" {
  default = "Redirect"
}

variable "identity_provider_logout_enabled" {
  default = false
}

variable "identity_provider_logout_request_url" {
  default = "https://logoutRequestUrl.com"
}

variable "identity_provider_logout_response_url" {
  default = "https://logoutResponseUrl.com"
}

#provide the IDP metadata
variable "identity_provider_metadata" {
  default = ""
}

variable "identity_provider_name_id_format" {
  default = "nameIdFormat"
}

variable "identity_provider_partner_name" {
  default = "partnerName"
}

#provide the IDP partner provider id
variable "identity_provider_partner_provider_id" {
  default = ""
}

variable "identity_provider_requested_authentication_context" {
  default = []
}

variable "identity_provider_require_force_authn" {
  default = false
}

variable "identity_provider_requires_encrypted_assertion" {
  default = false
}

variable "identity_provider_saml_ho_krequired" {
  default = false
}

variable "identity_provider_service_instance_identifier" {
  default = "serviceInstanceIdentifier"
}

variable "identity_provider_shown_on_login_page" {
  default = false
}

variable "identity_provider_signature_hash_algorithm" {
  default = "SHA-1"
}

#provide the IDP signing cert
variable "identity_provider_signing_certificate" {
  default = ""
}

variable "identity_provider_start_index" {
  default = 1
}

variable "identity_provider_tags_key" {
  default = "key"
}

variable "identity_provider_tags_value" {
  default = "value"
}

variable "identity_provider_type" {
  default = "SAML"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_access_token_url" {
  default = "accessTokenUrl"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_account_linking_enabled" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_admin_scope" {
  default = []
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_authz_url" {
  default = "authzUrl"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_auto_redirect_enabled" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_client_credential_in_payload" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_clock_skew_in_seconds" {
  default = 10
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_consumer_key" {
  default = "consumerKey"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_consumer_secret" {
  default = "consumerSecret"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_discovery_url" {
  default = "discoveryUrl"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_id_attribute" {
  default = "idAttribute"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_jit_prov_assigned_groups_value" {
  default = "value"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_jit_prov_group_static_list_enabled" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_social_jit_provisioning_enabled" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_profile_url" {
  default = "profileUrl"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_redirect_url" {
  default = "redirectUrl"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_registration_enabled" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_scope" {
  default = []
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_service_provider_name" {
  default = "serviceProviderName"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_status" {
  default = "created"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_cert_match_attribute" {
  default = "certMatchAttribute"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_check_on_ocsp_failure_enabled" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_enabled" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_location" {
  default = "crlLocation"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_reload_duration" {
  default = 10
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_allow_unknown_response_status" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_enable_signed_response" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_enabled" {
  default = false
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_responder_url" {
  default = "ocspResponderURL"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_revalidate_time" {
  default = 10
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_server_name" {
  default = "ocspServerName"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_trust_cert_chain" {
  default = []
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_other_cert_match_attribute" {
  default = "otherCertMatchAttribute"
}

variable "identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_user_match_attribute" {
  default = "userMatchAttribute"
}

variable "identity_provider_user_mapping_method" {
  default = "NameIDToUserAttribute"
}

variable "identity_provider_user_mapping_store_attribute" {
  default = "userName"
}


resource "oci_identity_domains_identity_provider" "test_identity_provider" {
  #Required
  enabled       = false
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  partner_name  = var.identity_provider_partner_name
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:IdentityProvider"]

  #Optional
  assertion_attribute   = var.identity_provider_assertion_attribute
  attribute_sets        = ["all"]
  attributes            = ""
  authn_request_binding = var.identity_provider_authn_request_binding
  authorization         = var.identity_provider_authorization
  /* #set value to a Policy id to reference a Policy
  correlation_policy {
    #Required
    type  = var.identity_provider_correlation_policy_type
    value = var.identity_provider_correlation_policy_value
  }
  */
  description                       = var.identity_provider_description
  encryption_certificate            = var.identity_provider_encryption_certificate
  external_id                       = "externalId"
  icon_url                          = var.identity_provider_icon_url
  idp_sso_url                       = var.identity_provider_idp_sso_url
  include_signing_cert_in_signature = var.identity_provider_include_signing_cert_in_signature
  /* #set value to a Group id to reference a Group
  jit_user_prov_assigned_groups {
    #Required
    value = var.identity_provider_jit_user_prov_assigned_groups_value
  }
  */
  jit_user_prov_attribute_update_enabled = var.identity_provider_jit_user_prov_attribute_update_enabled
  /* #set value to a MappedAttribute id to reference a MappedAttribute
  jit_user_prov_attributes {
    #Required
    value = var.identity_provider_jit_user_prov_attributes_value
  }
  */
  jit_user_prov_create_user_enabled               = var.identity_provider_jit_user_prov_create_user_enabled
  jit_user_prov_enabled                           = var.identity_provider_jit_user_prov_enabled
  jit_user_prov_group_assertion_attribute_enabled = var.identity_provider_jit_user_prov_group_assertion_attribute_enabled
  jit_user_prov_group_assignment_method           = var.identity_provider_jit_user_prov_group_assignment_method
  jit_user_prov_group_mapping_mode                = var.identity_provider_jit_user_prov_group_mapping_mode
  /* #set value to a Group id to reference a Group
  jit_user_prov_group_mappings {
    #Required
    idp_group = var.identity_provider_jit_user_prov_group_mappings_idp_group
    value     = var.identity_provider_jit_user_prov_group_mappings_value
  }
  */
  jit_user_prov_group_saml_attribute_name     = var.identity_provider_jit_user_prov_group_saml_attribute_name
  jit_user_prov_group_static_list_enabled     = var.identity_provider_jit_user_prov_group_static_list_enabled
  jit_user_prov_ignore_error_on_absent_groups = var.identity_provider_jit_user_prov_ignore_error_on_absent_groups
  logout_binding                              = var.identity_provider_logout_binding
  logout_enabled                              = var.identity_provider_logout_enabled
  logout_request_url                          = var.identity_provider_logout_request_url
  logout_response_url                         = var.identity_provider_logout_response_url
  metadata                                    = var.identity_provider_metadata
  name_id_format                              = var.identity_provider_name_id_format
  partner_provider_id                         = var.identity_provider_partner_provider_id
  requested_authentication_context            = var.identity_provider_requested_authentication_context
  require_force_authn                         = var.identity_provider_require_force_authn
  requires_encrypted_assertion                = var.identity_provider_requires_encrypted_assertion
  #use the latest if not provided
  # resource_type_schema_version                = var.identity_provider_resource_type_schema_version
  saml_ho_krequired           = var.identity_provider_saml_ho_krequired
  service_instance_identifier = var.identity_provider_service_instance_identifier
  shown_on_login_page         = var.identity_provider_shown_on_login_page
  signature_hash_algorithm    = var.identity_provider_signature_hash_algorithm
  signing_certificate         = var.identity_provider_signing_certificate
  succinct_id                 = "succinctId"
  tags {
    #Required
    key   = var.identity_provider_tags_key
    value = var.identity_provider_tags_value
  }
  type = var.identity_provider_type
  urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider {
    #Required
    account_linking_enabled = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_account_linking_enabled
    consumer_key            = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_consumer_key
    consumer_secret         = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_consumer_secret
    registration_enabled    = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_registration_enabled
    service_provider_name   = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_service_provider_name

    #Optional
    access_token_url             = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_access_token_url
    admin_scope                  = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_admin_scope
    authz_url                    = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_authz_url
    auto_redirect_enabled        = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_auto_redirect_enabled
    client_credential_in_payload = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_client_credential_in_payload
    clock_skew_in_seconds        = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_clock_skew_in_seconds
    discovery_url                = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_discovery_url
    id_attribute                 = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_id_attribute
    jit_prov_group_static_list_enabled                 = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_jit_prov_group_static_list_enabled
    social_jit_provisioning_enabled                 = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_social_jit_provisioning_enabled
    profile_url                  = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_profile_url
    redirect_url                 = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_redirect_url
    scope                        = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_scope
    status                       = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_status
  }
  urnietfparamsscimschemasoracleidcsextensionx509identity_provider {
    #Required
    cert_match_attribute      = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_cert_match_attribute
    signing_certificate_chain = ["signingCertificateChain"]
    user_match_attribute      = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_user_match_attribute

    #Optional
    crl_check_on_ocsp_failure_enabled  = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_check_on_ocsp_failure_enabled
    crl_enabled                        = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_enabled
    crl_location                       = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_location
    crl_reload_duration                = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_reload_duration
    ocsp_allow_unknown_response_status = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_allow_unknown_response_status
    ocsp_enable_signed_response        = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_enable_signed_response
    ocsp_enabled                       = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_enabled
    ocsp_responder_url                 = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_responder_url
    ocsp_revalidate_time               = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_revalidate_time
    ocsp_server_name                   = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_server_name
    ocsp_trust_cert_chain              = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_trust_cert_chain
    other_cert_match_attribute         = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_other_cert_match_attribute
  }
  user_mapping_method          = var.identity_provider_user_mapping_method
  user_mapping_store_attribute = var.identity_provider_user_mapping_store_attribute
  lifecycle {
    ignore_changes = [schemas]
  }
}

data "oci_identity_domains_identity_providers" "test_identity_providers" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  identity_provider_count  = var.identity_provider_identity_provider_count
  identity_provider_filter = var.identity_provider_identity_provider_filter
  attribute_sets           = []
  attributes               = ""
  authorization            = var.identity_provider_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.identity_provider_resource_type_schema_version
  start_index = var.identity_provider_start_index
}

