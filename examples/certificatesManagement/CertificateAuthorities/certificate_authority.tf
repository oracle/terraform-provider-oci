// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "kms_key_id" {}
variable "root_certificate_pem" {}

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
      user_id                      = var.certificate_authority_certificate_authority_config_subject_user_id
    }

    #Optional
    signing_algorithm               = var.certificate_authority_certificate_authority_config_signing_algorithm
    version_name = var.certificate_authority_certificate_authority_config_version_name
  }

  compartment_id = var.compartment_id
  kms_key_id     = var.kms_key_id
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


resource "oci_certificates_management_certificate_authority" "externally_managed_root_ca" {
  certificate_authority_config {
    certificate_pem = "${var.root_certificate_pem}"
    config_type = "ROOT_CA_MANAGED_EXTERNALLY"
  }
  compartment_id = "${var.compartment_id}"
  description = "description"
  external_key_description = "externally managed root's key description"
  freeform_tags = {
    "bar-key" = "value"
  }
  certificate_authority_rules {
    certificate_authority_max_validity_duration = "P100D"
    leaf_certificate_max_validity_duration = "P50D"
    rule_type = "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"
  }
  name = "externally-managed-root"
}

resource "oci_certificates_management_certificate_authority" "subordinate_managed_internally_issued_by_external_ca" {
  certificate_authority_config {
    config_type = "SUBORDINATE_CA_MANAGED_INTERNALLY_ISSUED_BY_EXTERNAL_CA"
    issuer_certificate_authority_id = "${oci_certificates_management_certificate_authority.externally_managed_root_ca.id}"
    signing_algorithm = "SHA256_WITH_RSA"
    subject {
      common_name = "www.example.com"
      country = "US"
      distinguished_name_qualifier = "distinguishedNameQualifier"
      domain_component = "domainComponent"
      generation_qualifier = "JR"
      given_name = "Sir"
      initials = "HAM"
      locality_name = "Seattle"
      organization = "WarehouseOrg"
      organizational_unit = "Products"
      pseudonym = "pseudonym"
      serial_number = "serialNumber"
      state_or_province_name = "Washington"
      street = "123 Main Street"
      surname = "Last"
      title = "Sr"
      user_id = "Neytiri"
    }
    version_name = "versionName"
  }

  certificate_authority_rules {
    certificate_authority_max_validity_duration = "P1000D"
    leaf_certificate_max_validity_duration = "P500D"
    rule_type = "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"
  }

  certificate_authority_rules {
    rule_type = "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"
    path_length_constraint = 3
    name_constraint {
      excluded_subtree {
        type  = "DNS"
        value = "bad.com"
      }

      permitted_subtree {
        type  = "DNS"
        value = "good.example.com"
      }
    }
  }
  compartment_id = "${var.compartment_id}"
  description = "description"
  freeform_tags = {
    "bar-key" = "value"
  }
  kms_key_id = var.kms_key_id
  name = "SubCAManagedInternallyIssuedByExternalCA"
}


data "oci_certificates_management_certificate_authorities" "test_certificate_authorities" {

  #Optional
  certificate_authority_id        = oci_certificates_management_certificate_authority.test_certificate_authority.id
  state                           = var.certificate_authority_state
}

data "oci_certificates_management_certificate_authority" "certificate_authority_data_source" {
  certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
}


data "oci_certificates_management_certificate_authority_version" "test_certificate_authority_version" {

  certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
  certificate_authority_version_number = 1
}

