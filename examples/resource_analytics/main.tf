// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "resource_analytics_instance_subnet_id" {
  default = "ocid1.subnet..."
}

variable "resource_analytics_instance_adw_admin_password_password" {
  default = "BEstrO0ng_#11"
}

variable "resource_analytics_instance_adw_admin_password_vault_secret_id" {
  default = "ocid1.vaultsecret..."
}

variable "resource_analytics_instance_description" {
  default = "Resource Analytics Instance created with Terraform"
}

variable "resource_analytics_instance_display_name" {
  default = "Default Resource Analytics Instance"
}

variable "resource_analytics_instance_freeform_tags" {
  default = { "CreatedWith" = "Terraform" }
}

variable "resource_analytics_instance_is_mutual_tls_required" {
  default = false
}

variable "resource_analytics_instance_license_model" {
  default = "LICENSE_INCLUDED"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


###
# Create Resource Analytics Instance
###
resource "oci_resource_analytics_resource_analytics_instance" "test_resource_analytics_instance" {
  #Required
  timeouts {
    create = "1h"
    update = "1h"
    delete = "1h"
  }

  #Required
  adw_admin_password {
    #PlainTextPassword type
    password_type = "PLAIN_TEXT"
    password      = var.resource_analytics_instance_adw_admin_password_password

    # OR

    #VaultSecretPassword type
    # password_type = "VAULT_SECRET"
    # secret_id     = var.resource_analytics_instance_adw_admin_password_vault_secret_id
  }
  compartment_id = var.compartment_ocid
  subnet_id      = var.resource_analytics_instance_subnet_id

  #Optional
  description            = var.resource_analytics_instance_description
  display_name           = var.resource_analytics_instance_display_name
  freeform_tags          = var.resource_analytics_instance_freeform_tags
  is_mutual_tls_required = var.resource_analytics_instance_is_mutual_tls_required
  license_model          = var.resource_analytics_instance_license_model
}


###
# Add a Tenancy Attachment to monitor additional tenancies (optional)
###
variable "added_tenancy_attachment_ocid" {
  default = "ocid1.tenancy..."
}

variable "added_tenancy_attachment_description" {
  default = "Tenancy description"
}

resource "oci_resource_analytics_tenancy_attachment" "test_tenancy_attachment" {
  #Required
  resource_analytics_instance_id = oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id
  tenancy_id                     = var.added_tenancy_attachment_ocid

  #Optional
  description = var.added_tenancy_attachment_description
}


###
# Add all subscribed regions as Monitored Regions (or add at least one Monitored Region)
###
data "oci_identity_region_subscriptions" "test_region_subscriptions" {
  #Required
  tenancy_id = var.tenancy_ocid

  filter {
    name   = "state"
    values = ["READY"]
  }
}

resource "oci_resource_analytics_monitored_region" "test_monitored_regions" {
  depends_on = [oci_resource_analytics_tenancy_attachment.test_tenancy_attachment]

  for_each = toset([for region in data.oci_identity_region_subscriptions.test_region_subscriptions.region_subscriptions : region.region_name])

  #Required
  region_id                      = each.value
  resource_analytics_instance_id = oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id
}


###
# Enable OAC on Resource Analytics Instance (optional)
###
variable "resource_analytics_instance_oac_management_attachment_details_idcs_domain_id" {
  default = "ocid1.domain..."
}

variable "resource_analytics_instance_oac_management_attachment_details_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "resource_analytics_instance_oac_management_attachment_type" {
  default = "MANAGED"
}

variable "resource_analytics_instance_oac_management_enable_oac" {
  default = true
}

resource "oci_resource_analytics_resource_analytics_instance_oac_management" "test_resource_analytics_instance_enable_oac" {
  #Required
  timeouts {
    create = "2h"
    update = "2h"
    delete = "2h"
  }

  #Required
  resource_analytics_instance_id = oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id
  attachment_type                = var.resource_analytics_instance_oac_management_attachment_type
  enable_oac                     = var.resource_analytics_instance_oac_management_enable_oac

  #Required
  attachment_details {

    #Required
    idcs_domain_id = var.resource_analytics_instance_oac_management_attachment_details_idcs_domain_id

    #Optional
    license_model  = var.resource_analytics_instance_oac_management_attachment_details_license_model
    #Private OAC
    # network_details {
    #   subnet_id = oci_core_subnet.test_subnet.id
    # }
  }
}