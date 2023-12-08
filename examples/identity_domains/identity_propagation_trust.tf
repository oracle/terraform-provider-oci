// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "identity_propagation_trust_identity_propagation_trust_count" {
  default = 10
}

variable "identity_propagation_trust_identity_propagation_trust_filter" {
  default = ""
}

variable "identity_propagation_trust_active" {
  default = false
}

variable "identity_propagation_trust_allow_impersonation" {
  default = false
}

variable "identity_propagation_trust_authorization" {
  default = "authorization"
}

variable "identity_propagation_trust_client_claim_name" {
  default = "clientClaimName"
}

variable "identity_propagation_trust_description" {
  default = "description"
}

variable "identity_propagation_trust_impersonation_service_users_ocid" {
  default = "ocid"
}

variable "identity_propagation_trust_impersonation_service_users_rule" {
  default = "rule"
}

variable "identity_propagation_trust_impersonation_service_users_value" {
  default = "value"
}

variable "identity_propagation_trust_issuer" {
  default = "issuer"
}

variable "identity_propagation_trust_keytab_secret_ocid" {
  default = "secretOcid"
}

variable "identity_propagation_trust_keytab_secret_version" {
  default = 10
}

variable "identity_propagation_trust_name" {
  default = "name"
}

variable "identity_propagation_trust_public_certificate" {
  default = "publicCertificate"
}

variable "identity_propagation_trust_public_key_endpoint" {
  default = "publicKeyEndpoint"
}

variable "identity_propagation_trust_start_index" {
  default = 1
}

variable "identity_propagation_trust_subject_claim_name" {
  default = "subjectClaimName"
}

variable "identity_propagation_trust_subject_mapping_attribute" {
  default = "subjectMappingAttribute"
}

variable "identity_propagation_trust_subject_type" {
  default = "User"
}

variable "identity_propagation_trust_tags_key" {
  default = "key"
}

variable "identity_propagation_trust_tags_value" {
  default = "value"
}

variable "identity_propagation_trust_type" {
  default = "JWT"
}

resource "oci_identity_domains_user" "test_identity_propagation_trust_user" {
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas = [    "urn:ietf:params:scim:schemas:core:2.0:User"]
  name {
    family_name = "familyName"
  }
  user_name = "testIdentityPropagationTrustUser"

  emails {
    type = "work"
    value = "value@email.com"
    primary = true
  }

  lifecycle {
    ignore_changes = [schemas]
  }
}

resource "oci_identity_domains_identity_propagation_trust" "test_identity_propagation_trust" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  issuer        = var.identity_propagation_trust_issuer
  name          = var.identity_propagation_trust_name
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:IdentityPropagationTrust"]
  type          = var.identity_propagation_trust_type

  #Optional
  account_id          = "accountId"
  active              = var.identity_propagation_trust_active
  allow_impersonation = var.identity_propagation_trust_allow_impersonation
  attribute_sets      = ["all"]
  attributes          = ""
  authorization       = var.identity_propagation_trust_authorization
  client_claim_name   = var.identity_propagation_trust_client_claim_name
  /* if client_claim_name is set, client_claim_values cannot be empty */
  client_claim_values = ["clientClaimValues"]
  description         = var.identity_propagation_trust_description
  /* provide users' id */
  impersonation_service_users {
    #Required
    rule  = var.identity_propagation_trust_impersonation_service_users_rule
    value = oci_identity_domains_user.test_identity_propagation_trust_user.id

    #Optional
    /* #if setting ocid here, make sure it is the ocid of the same resource that "value" is referring to. */
    ocid = oci_identity_domains_user.test_identity_propagation_trust_user.ocid
  }
  keytab {
    #Required
    secret_ocid = var.identity_propagation_trust_keytab_secret_ocid

    #Optional
    secret_version = var.identity_propagation_trust_keytab_secret_version
  }
  /* oauth_clients must be defined for IdentityPropagationTrust */
  oauth_clients                = ["oauthClients"]
  public_certificate           = var.identity_propagation_trust_public_certificate
  public_key_endpoint          = var.identity_propagation_trust_public_key_endpoint
  #use the latest if not provided
  # resource_type_schema_version = var.identity_propagation_trust_resource_type_schema_version
  subject_claim_name           = var.identity_propagation_trust_subject_claim_name
  subject_mapping_attribute    = var.identity_propagation_trust_subject_mapping_attribute
  subject_type                 = var.identity_propagation_trust_subject_type
  tags {
    #Required
    key   = var.identity_propagation_trust_tags_key
    value = var.identity_propagation_trust_tags_value
  }
}

data "oci_identity_domains_identity_propagation_trusts" "test_identity_propagation_trusts" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  identity_propagation_trust_count  = var.identity_propagation_trust_identity_propagation_trust_count
  identity_propagation_trust_filter = var.identity_propagation_trust_identity_propagation_trust_filter
  attribute_sets                    = ["all"]
  attributes                        = ""
  authorization                     = var.identity_propagation_trust_authorization
  #use the latest if not provided
  # resource_type_schema_version      = ""
  start_index                       = var.identity_propagation_trust_start_index
}

