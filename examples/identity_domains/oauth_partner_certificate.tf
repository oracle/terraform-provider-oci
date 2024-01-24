// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "oauth_partner_certificate_oauth_partner_certificate_count" {
  default = 10
}

variable "oauth_partner_certificate_oauth_partner_certificate_filter" {
  default = ""
}

variable "oauth_partner_certificate_authorization" {
  default = "authorization"
}

variable "oauth_partner_certificate_certificate_alias" {
  default = "certificateAlias"
}

variable "oauth_partner_certificate_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "oauth_partner_certificate_start_index" {
  default = 1
}

#provide the x509 base 64 certificate
variable "oauth_partner_certificate_x509base64certificate" {
  default = ""
}

resource "oci_identity_domains_oauth_partner_certificate" "test_oauth_partner_certificate" {
  #Required
  certificate_alias = var.oauth_partner_certificate_certificate_alias
  idcs_endpoint     = data.oci_identity_domain.test_domain.url
  schemas           = ["urn:ietf:params:scim:schemas:oracle:idcs:OAuthPartnerCertificate"]

  #Optional
  authorization                = var.oauth_partner_certificate_authorization
  external_id                  = "externalId"
#  resource_type_schema_version = var.oauth_partner_certificate_resource_type_schema_version
  x509base64certificate        = var.oauth_partner_certificate_x509base64certificate
}

data "oci_identity_domains_oauth_partner_certificates" "test_oauth_partner_certificates" {
  #Required
  idcs_endpoint     = data.oci_identity_domain.test_domain.url

  #Optional
  oauth_partner_certificate_count  = var.oauth_partner_certificate_oauth_partner_certificate_count
  oauth_partner_certificate_filter = var.oauth_partner_certificate_oauth_partner_certificate_filter
  authorization                    = var.oauth_partner_certificate_authorization
#  resource_type_schema_version     = var.oauth_partner_certificate_resource_type_schema_version
  start_index                      = var.oauth_partner_certificate_start_index
}

