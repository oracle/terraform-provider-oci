// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "certificate_certificate_config_certificate_profile_type" {
  default = "TLS_SERVER_OR_CLIENT"
}

variable "certificate_certificate_config_config_type" {
  default = "ISSUED_BY_INTERNAL_CA"
}

variable "certificate_certificate_config_csr_pem" {
  default = "-----BEGIN CERTIFICATE REQUEST-----<var>&lt;csrcontent&gt;</var>n-----END CERTIFICATE REQUEST-----"
}

variable "certificate_certificate_config_key_algorithm" {
  default = "RSA2048"
}

variable "certificate_certificate_config_signature_algorithm" {
  default = "SHA256_WITH_RSA"
}

variable "certificate_certificate_config_subject_common_name" {
  default = "Sample_Cert"
}

variable "certificate_certificate_config_subject_country" {
  default = "IE"
}

variable "certificate_certificate_config_subject_distinguished_name_qualifier" {
  default = "distinguishedNameQualifier"
}

variable "certificate_certificate_config_subject_domain_component" {
  default = "domainComponent"
}

variable "certificate_certificate_config_subject_generation_qualifier" {
  default = "Sr."
}

variable "certificate_certificate_config_subject_given_name" {
  default = "givenName"
}

variable "certificate_certificate_config_subject_initials" {
  default = "LI"
}

variable "certificate_certificate_config_subject_locality_name" {
  default = "Dublin"
}

variable "certificate_certificate_config_subject_organization" {
  default = "Company Ltd."
}

variable "certificate_certificate_config_subject_organizational_unit" {
  default = "R&D"
}

variable "certificate_certificate_config_subject_pseudonym" {
  default = "pseudonym"
}

variable "certificate_certificate_config_subject_serial_number" {
  default = "rvZ6gMLf4Pc3hoGu"
}

variable "certificate_certificate_config_subject_state_or_province_name" {
  default = "Leinster"
}

variable "certificate_certificate_config_subject_street" {
  default = "Suffolk St"
}

variable "certificate_certificate_config_subject_surname" {
  default = "surname"
}

variable "certificate_certificate_config_subject_title" {
  default = "title"
}

variable "certificate_certificate_config_subject_alternative_names_type" {
  default = "DNS"
}

variable "certificate_certificate_config_subject_alternative_names_value" {
  default = "value"
}

variable "certificate_certificate_config_validity_time_of_validity_not_after" {
  default = "2022-10-10T02:50:29.396Z"
}

variable "certificate_certificate_config_validity_time_of_validity_not_before" {
  default = "2021-10-09T02:50:29.396Z"
}

variable "certificate_certificate_config_version_name" {
  default = "myCertVersion"
}

variable "certificate_certificate_rules_advance_renewal_period" {
  default = "P60D"
}

variable "certificate_certificate_rules_renewal_interval" {
  default = "P365D"
}

variable "certificate_certificate_rules_rule_type" {
  default = "CERTIFICATE_RENEWAL_RULE"
}

variable "certificate_defined_tags_value" {
  default = "value"
}

variable "certificate_description" {
  default = "description"
}

variable "certificate_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "certificate_name" {
  default = "test-cert"
}

variable "certificate_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_certificates_management_certificate" "test_certificate" {
  #Required
  certificate_config {
    #Required
    config_type = var.certificate_certificate_config_config_type

    #Optional
    certificate_profile_type        = var.certificate_certificate_config_certificate_profile_type
    csr_pem                         = var.certificate_certificate_config_csr_pem
    issuer_certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
    key_algorithm                   = var.certificate_certificate_config_key_algorithm
    signature_algorithm             = var.certificate_certificate_config_signature_algorithm
    subject {

      #Optional
      common_name                  = var.certificate_certificate_config_subject_common_name
      country                      = var.certificate_certificate_config_subject_country
      distinguished_name_qualifier = var.certificate_certificate_config_subject_distinguished_name_qualifier
      domain_component             = var.certificate_certificate_config_subject_domain_component
      generation_qualifier         = var.certificate_certificate_config_subject_generation_qualifier
      given_name                   = var.certificate_certificate_config_subject_given_name
      initials                     = var.certificate_certificate_config_subject_initials
      locality_name                = var.certificate_certificate_config_subject_locality_name
      organization                 = var.certificate_certificate_config_subject_organization
      organizational_unit          = var.certificate_certificate_config_subject_organizational_unit
      pseudonym                    = var.certificate_certificate_config_subject_pseudonym
      serial_number                = var.certificate_certificate_config_subject_serial_number
      state_or_province_name       = var.certificate_certificate_config_subject_state_or_province_name
      street                       = var.certificate_certificate_config_subject_street
      surname                      = var.certificate_certificate_config_subject_surname
      title                        = var.certificate_certificate_config_subject_title
      user_id                      = oci_identity_user.test_user.id
    }
    subject_alternative_names {

      #Optional
      type  = var.certificate_certificate_config_subject_alternative_names_type
      value = var.certificate_certificate_config_subject_alternative_names_value
    }
    validity {

      #Optional
      time_of_validity_not_after  = var.certificate_certificate_config_validity_time_of_validity_not_after
      time_of_validity_not_before = var.certificate_certificate_config_validity_time_of_validity_not_before
    }
    version_name = var.certificate_certificate_config_version_name
  }
  compartment_id = var.compartment_id
  name           = var.certificate_name

  #Optional
  certificate_rules {
    #Required
    advance_renewal_period = var.certificate_certificate_rules_advance_renewal_period
    renewal_interval       = var.certificate_certificate_rules_renewal_interval
    rule_type              = var.certificate_certificate_rules_rule_type
  }
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.certificate_defined_tags_value)
  description   = var.certificate_description
  freeform_tags = var.certificate_freeform_tags
}

data "oci_certificates_management_certificates" "test_certificates" {

  #Optional
  certificate_id                  = oci_certificates_management_certificate.test_certificate.id
  compartment_id                  = var.compartment_id
  issuer_certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
  name                            = var.certificate_name
  state                           = var.certificate_state
}

