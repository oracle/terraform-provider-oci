// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
  default = "tenancy"
}
variable "user_ocid" {
  default = "user"
}
variable "fingerprint" {
  default = ""
}
variable "private_key_path" {
  default = "private_key"
}
variable "region" {
  default = "us-ashburn-1"
}
variable "compartment_id" {
  default = "compartment"
}

variable "product_id" {
  type = string
  default = "product_id"
}

variable "seller_id" {
  type = string
  default = "seller_id"
}

variable "tenant_id" {
  type = string
  default = "tenant_id"
}

# defined tags are "namespace.tagName" = "value"
variable "subscription_defined_tags" {
  type    = map(string)
  default = {}
}

variable "subscription_additional_details_key" {
  default = "key"
}

variable "subscription_additional_details_value" {
  default = "value"
}

resource "random_id" "suffix" {
  byte_length = 4
}

locals {
  subscription_display_name = "test-subscription-15Mar-1"
}

variable "subscription_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "subscription_id" {
  default = "id"
}

variable "subscription_lifecycle_details" {
  default = "CREATED"
}

variable "subscription_realm" {
  default = "OC1"
}

variable "subscription_region" {
  default = "us-ashburn-1"
}

variable "subscription_source_type" {
  default = "THIRD_PARTY"
}

variable "subscription_subscription_details_amount" {
  default = 1.0
}

variable "subscription_subscription_details_billing_details_has_gov_sku" {
  default = false
}

variable "subscription_subscription_details_billing_details_meters_extended_metadata_key" {
  default = "key"
}

variable "subscription_subscription_details_billing_details_meters_extended_metadata_value" {
  default = "value"
}

variable "subscription_subscription_details_billing_details_meters_name" {
  default = "name"
}

variable "subscription_subscription_details_billing_details_meters_rate_allocation" {
  default = 1.0
}

variable "subscription_subscription_details_billing_details_metric_type" {
  default = "OCPU_HOURS"
}

variable "subscription_subscription_details_billing_details_rate_allocation" {
  default = 1.0
}

variable "subscription_subscription_details_billing_details_sku" {
  default = "sku"
}

variable "subscription_subscription_details_currency" {
  default = "USD"
}

variable "subscription_subscription_details_is_auto_renew" {
  default = false
}

variable "subscription_subscription_details_partner_registration_url" {
  default = "https://oracle.com"
}

variable "subscription_subscription_details_pricing_plan_billing_frequency" {
  default = "YEARLY"
}

variable "subscription_subscription_details_pricing_plan_plan_description" {
  default = "test description"
}

variable "subscription_subscription_details_pricing_plan_plan_duration" {
  default = "ANNUAL"
}

variable "subscription_subscription_details_pricing_plan_plan_name" {
  default = "Base"
}

variable "subscription_subscription_details_pricing_plan_plan_type" {
  default = "FIXED"
}

variable "subscription_subscription_details_pricing_plan_rates_currency" {
  default = "USD"
}

variable "subscription_subscription_details_pricing_plan_rates_rate" {
  default = 5000
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_self_subscription" "test_subscription" {
  #Required
  compartment_id = var.compartment_id
  product_id     = var.product_id
  seller_id      = var.seller_id
  subscription_details {
    #Required
    billing_details {
      #Required
      meters {
        #Required
        name            = var.subscription_subscription_details_billing_details_meters_name
        rate_allocation = var.subscription_subscription_details_billing_details_meters_rate_allocation

        #Optional
        extended_metadata {
          #Required
          key   = var.subscription_subscription_details_billing_details_meters_extended_metadata_key
          value = var.subscription_subscription_details_billing_details_meters_extended_metadata_value
        }
      }
      metric_type     = var.subscription_subscription_details_billing_details_metric_type
      rate_allocation = var.subscription_subscription_details_billing_details_rate_allocation
      sku             = var.subscription_subscription_details_billing_details_sku

      #Optional
      has_gov_sku = var.subscription_subscription_details_billing_details_has_gov_sku
    }
    partner_registration_url = var.subscription_subscription_details_partner_registration_url
    pricing_plan {
      #Required
      billing_frequency = var.subscription_subscription_details_pricing_plan_billing_frequency
      plan_name         = var.subscription_subscription_details_pricing_plan_plan_name
      plan_type         = var.subscription_subscription_details_pricing_plan_plan_type
      rates {
        #Required
        currency = var.subscription_subscription_details_pricing_plan_rates_currency
        rate     = var.subscription_subscription_details_pricing_plan_rates_rate
      }

      #Optional
      plan_description = var.subscription_subscription_details_pricing_plan_plan_description
      plan_duration    = var.subscription_subscription_details_pricing_plan_plan_duration
    }

    #Optional
    amount        = var.subscription_subscription_details_amount
    currency      = var.subscription_subscription_details_currency
    is_auto_renew = var.subscription_subscription_details_is_auto_renew
  }
  tenant_id = var.tenant_id

  #Optional
  additional_details {
    #Required
    key   = var.subscription_additional_details_key
    value = var.subscription_additional_details_value
  }
  defined_tags  = var.subscription_defined_tags
  display_name  = local.subscription_display_name
  freeform_tags = var.subscription_freeform_tags
  realm         = var.subscription_realm
  region        = var.subscription_region
  source_type   = var.subscription_source_type
}

data "oci_self_subscriptions" "test_subscriptions" {

  #Optional
  compartment_id    = var.compartment_id
  display_name      = local.subscription_display_name
  id                = var.subscription_id
  lifecycle_details = var.subscription_lifecycle_details
}

