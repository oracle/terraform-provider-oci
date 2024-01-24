// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {

}

variable "publication_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "publication_is_agreement_acknowledged" {
  default = true
}

variable "publication_listing_type" {
  default = "COMMUNITY"
}

variable "publication_long_description" {
  default = "longDescription"
}

variable "publication_name" {
  default = "name"
}

variable "publication_names" {
  default = ["name"]
}

variable "publication_operating_systems" {
  default = []
}

variable "publication_package_details_eula_eula_type" {
  default = "TEXT"
}

variable "publication_package_details_eula_license_text" {
  default = "licenseText"
}

variable "publication_package_details_operating_system_name" {
  default = "name"
}

variable "publication_package_details_package_type" {
  default = "IMAGE"
}

variable "publication_package_details_package_version" {
  default = "packageVersion"
}

variable "publication_short_description" {
  default = "shortDescription"
}

variable "publication_support_contacts_email" {
  default = "email"
}

variable "publication_support_contacts_name" {
  default = "name"
}

variable "publication_support_contacts_phone" {
  default = "phone"
}

variable "publication_support_contacts_subject" {
  default = "subject"
}

variable "image_id" {

}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_marketplace_publication" "test_publication" {
  #Required
  compartment_id            = var.compartment_id
  is_agreement_acknowledged = var.publication_is_agreement_acknowledged
  listing_type              = var.publication_listing_type
  name                      = var.publication_name
  package_details {
    #Required
    eula {
      #Required
      eula_type = var.publication_package_details_eula_eula_type

      #Optional
      license_text = var.publication_package_details_eula_license_text
    }
    operating_system {

      #Optional
      name = var.publication_package_details_operating_system_name
    }
    package_type    = var.publication_package_details_package_type
    package_version = var.publication_package_details_package_version

    #Required
    image_id = var.image_id
  }
  short_description = var.publication_short_description
  support_contacts {

    #Optional
    email   = var.publication_support_contacts_email
    name    = var.publication_support_contacts_name
    phone   = var.publication_support_contacts_phone
    subject = var.publication_support_contacts_subject
  }

  #Optional
  freeform_tags    = var.publication_freeform_tags
  long_description = var.publication_long_description
}

data "oci_marketplace_publications" "test_publications" {
  #Required
  compartment_id = var.compartment_id
  listing_type   = var.publication_listing_type

  #Optional
  name              = var.publication_names
  operating_systems = var.publication_operating_systems
  publication_id    = oci_marketplace_publication.test_publication.id
}

