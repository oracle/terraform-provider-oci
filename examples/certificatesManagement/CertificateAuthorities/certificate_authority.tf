// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "certificate_authority_certificate_authority_config_config_type" {
  default = "ROOT_CA_GENERATED_INTERNALLY"
}

variable "certificate_authority_certificate_authority_config_signing_algorithm" {
  default = "SHA512_WITH_RSA"
}

variable "certificate_authority_certificate_authority_config_subject_common_name" {
  default = "Sample_CA"
}

variable "certificate_authority_certificate_authority_config_subject_country" {
  default = "IE"
}

variable "certificate_authority_certificate_authority_config_subject_distinguished_name_qualifier" {
  default = "distinguishedNameQualifier"
}

variable "certificate_authority_certificate_authority_config_subject_domain_component" {
  default = "domainComponent"
}

variable "certificate_authority_certificate_authority_config_subject_generation_qualifier" {
  default = "Sr."
}

variable "certificate_authority_certificate_authority_config_subject_given_name" {
  default = "givenName"
}

variable "certificate_authority_certificate_authority_config_subject_initials" {
  default = "LI"
}

variable "certificate_authority_certificate_authority_config_subject_locality_name" {
  default = "Dublin"
}

variable "certificate_authority_certificate_authority_config_subject_organization" {
  default = "Company Ltd."
}

variable "certificate_authority_certificate_authority_config_subject_organizational_unit" {
  default = "R&D"
}

variable "certificate_authority_certificate_authority_config_subject_pseudonym" {
  default = "pseudonym"
}

variable "certificate_authority_certificate_authority_config_subject_serial_number" {
  default = "rvZ6gMLf4Pc3hoGu"
}

variable "certificate_authority_certificate_authority_config_subject_state_or_province_name" {
  default = "Leinster"
}

variable "certificate_authority_certificate_authority_config_subject_street" {
  default = "Suffolk St"
}

variable "certificate_authority_certificate_authority_config_subject_surname" {
  default = "surname"
}

variable "certificate_authority_certificate_authority_config_subject_title" {
  default = "title"
}

variable "certificate_authority_certificate_authority_config_subject_user_id" {
  default = "user_id"
}

variable "certificate_authority_certificate_authority_config_validity_time_of_validity_not_after" {
  default = "2031-07-05T21:10:29.999Z"
}

variable "certificate_authority_certificate_authority_config_validity_time_of_validity_not_before" {
  default = "2021-07-05T21:10:29.999Z"
}

variable "certificate_authority_certificate_authority_config_version_name" {
  default = "versionName"
}

variable "certificate_authority_certificate_authority_rules_certificate_authority_max_validity_duration" {
  default = "P1825D"
}

variable "certificate_authority_certificate_authority_rules_leaf_certificate_max_validity_duration" {
  default = "P365D"
}

variable "certificate_authority_certificate_authority_rules_rule_type" {
  default = "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"
}

variable "certificate_authority_certificate_revocation_list_details_custom_formatted_urls" {
  default = []
}

variable "certificate_authority_certificate_revocation_list_details_object_storage_config_object_storage_namespace" {
  default = "example-namespace"
}

variable "certificate_authority_certificate_revocation_list_details_object_storage_config_object_storage_object_name_format" {
  default = "versionName{}.crl"
}

variable "certificate_authority_defined_tags_value" {
  default = "value"
}

variable "certificate_authority_description" {
  default = "My Example CA"
}

variable "certificate_authority_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "certificate_authority_name" {
  default = "Sample_CA"
}

variable "certificate_authority_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_certificates_management_certificate_authority" "test_certificate_authority" {
  #Required
  certificate_authority_config {
    #Required
    config_type = var.certificate_authority_certificate_authority_config_config_type
    subject {

      #Optional
      common_name                  = var.certificate_authority_certificate_authority_config_subject_common_name
      country                      = var.certificate_authority_certificate_authority_config_subject_country
      distinguished_name_qualifier = var.certificate_authority_certificate_authority_config_subject_distinguished_name_qualifier
      domain_component             = var.certificate_authority_certificate_authority_config_subject_domain_component
      generation_qualifier         = var.certificate_authority_certificate_authority_config_subject_generation_qualifier
      given_name                   = var.certificate_authority_certificate_authority_config_subject_given_name
      initials                     = var.certificate_authority_certificate_authority_config_subject_initials
      locality_name                = var.certificate_authority_certificate_authority_config_subject_locality_name
      organization                 = var.certificate_authority_certificate_authority_config_subject_organization
      organizational_unit          = var.certificate_authority_certificate_authority_config_subject_organizational_unit
      pseudonym                    = var.certificate_authority_certificate_authority_config_subject_pseudonym
      serial_number                = var.certificate_authority_certificate_authority_config_subject_serial_number
      state_or_province_name       = var.certificate_authority_certificate_authority_config_subject_state_or_province_name
      street                       = var.certificate_authority_certificate_authority_config_subject_street
      surname                      = var.certificate_authority_certificate_authority_config_subject_surname
      title                        = var.certificate_authority_certificate_authority_config_subject_title
      user_id                      = oci_identity_user.test_user.id
    }

    #Optional
    issuer_certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
    signing_algorithm               = var.certificate_authority_certificate_authority_config_signing_algorithm
    validity {

      #Optional
      time_of_validity_not_after  = var.certificate_authority_certificate_authority_config_validity_time_of_validity_not_after
      time_of_validity_not_before = var.certificate_authority_certificate_authority_config_validity_time_of_validity_not_before
    }
    version_name = var.certificate_authority_certificate_authority_config_version_name
  }

  compartment_id = var.compartment_id
  kms_key_id     = oci_kms_key.test_key.id
  name           = var.certificate_authority_name

  #Optional
  certificate_authority_rules {
    #Required
    rule_type = var.certificate_authority_certificate_authority_rules_rule_type

    #Optional
    certificate_authority_max_validity_duration = var.certificate_authority_certificate_authority_rules_certificate_authority_max_validity_duration
    leaf_certificate_max_validity_duration      = var.certificate_authority_certificate_authority_rules_leaf_certificate_max_validity_duration
  }
  certificate_revocation_list_details {
    #Required
    object_storage_config {
      #Required
      object_storage_bucket_name        = oci_objectstorage_bucket.test_bucket.name
      object_storage_object_name_format = var.certificate_authority_certificate_revocation_list_details_object_storage_config_object_storage_object_name_format

      #Optional
      object_storage_namespace = oci_objectstorage_bucket.test_bucket.namespace
    }

    #Optional
    custom_formatted_urls = var.certificate_authority_certificate_revocation_list_details_custom_formatted_urls
  }
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.certificate_authority_defined_tags_value)
  description   = var.certificate_authority_description
  freeform_tags = var.certificate_authority_freeform_tags
}

data "oci_certificates_management_certificate_authorities" "test_certificate_authorities" {

  #Optional
  certificate_authority_id        = oci_certificates_management_certificate_authority.test_certificate_authority.id
  compartment_id                  = var.compartment_id
  issuer_certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
  name                            = var.certificate_authority_name
  state                           = var.certificate_authority_state
}

