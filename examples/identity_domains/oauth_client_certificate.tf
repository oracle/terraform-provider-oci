// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "oauth_client_certificate_oauth_client_certificate_count" {
  default = 10
}

variable "oauth_client_certificate_oauth_client_certificate_filter" {
  default = ""
}

variable "oauth_client_certificate_authorization" {
  default = "authorization"
}

variable "oauth_client_certificate_certificate_alias" {
  default = "certificateAlias"
}

variable "oauth_client_certificate_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "oauth_client_certificate_start_index" {
  default = 1
}

# provide the x509 base 64 certificate
variable "oauth_client_certificate_x509base64certificate" {
  default = ""
}

resource "oci_identity_domains_oauth_client_certificate" "test_oauth_client_certificate" {
  #Required
  certificate_alias     = var.oauth_client_certificate_certificate_alias
  idcs_endpoint         = data.oci_identity_domain.test_domain.url
  schemas               = ["urn:ietf:params:scim:schemas:oracle:idcs:OAuthClientCertificate"]
  x509base64certificate = var.oauth_client_certificate_x509base64certificate

  #Optional
  authorization                = var.oauth_client_certificate_authorization
  external_id                  = "externalId"
#  resource_type_schema_version = var.oauth_client_certificate_resource_type_schema_version
}

data "oci_identity_domains_oauth_client_certificates" "test_oauth_client_certificates" {
  #Required
  idcs_endpoint         = data.oci_identity_domain.test_domain.url

  #Optional
  oauth_client_certificate_count  = var.oauth_client_certificate_oauth_client_certificate_count
  oauth_client_certificate_filter = var.oauth_client_certificate_oauth_client_certificate_filter
  authorization                   = var.oauth_client_certificate_authorization
#  resource_type_schema_version    = var.oauth_client_certificate_resource_type_schema_version
  start_index                     = var.oauth_client_certificate_start_index
}

