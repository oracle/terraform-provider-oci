// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "compartment_id" {
  description = "Compartment in which to create the subscription."
  type        = string
}

variable "tenant_id" {
  description = "Tenancy that owns the subscription."
  type        = string
}

variable "seller_id" {
  description = "Marketplace publisher OCID for the listing."
  type        = string
}

variable "product_id" {
  description = "Marketplace listing OCID."
  type        = string
}

variable "region" {
  default = "us-ashburn-1"
}

variable "config_file_profile" {
  default = "terraform-federation-test"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  region              = var.region
}

resource "oci_self_subscription" "test_subscription" {
  compartment_id = var.compartment_id
  tenant_id      = var.tenant_id
  seller_id      = var.seller_id
  product_id     = var.product_id
  display_name   = "S-Open-AI-Test-Listing-${formatdate("YYYYMMDDhhmmss", timestamp())}"
  source_type    = "THIRD_PARTY"
  realm          = "OC1"
  region         = var.region

  subscription_details {
    amount        = 89
    currency      = "USD"
    is_auto_renew = true

    partner_registration_url = "https://www.google.com/"

    billing_details {
      pricing_plan_key = "16ff7edf-a17f-4f77-88b8-1a09bef6df8f"
      billing_model    = "FLAT_RATE"
      sku              = "MP10257"
      metric_type      = "INSTANCE_HOURS"
      rate_allocation  = 1
      has_gov_sku      = false

      meters {
        name            = "MP_INSP_JX"
        rate_allocation = 1
      }
    }

    billing_details {
      pricing_plan_key = "16ff7edf-a17f-4f77-88b8-1a09bef6df8f"
      billing_model    = "USAGE_BASED"
      sku              = "MP10258"
      metric_type      = "INSTANCE_HOURS"
      rate_allocation  = 1
      has_gov_sku      = false

      meters {
        name            = "MP_INSP_JY"
        rate_allocation = 1
      }
    }

    pricing_plan {
      plan_type         = "HYBRID"
      plan_name         = "Pro-Enterprice"
      plan_description  = "Enterprise marketplace listing with licensed-user and output-token usage dimensions."
      billing_frequency = "ANNUAL"
      plan_duration     = "ANNUAL"

      rates {
        currency = "USD"
        rate     = 89
      }

      dimensions {
        dimension_key               = "f5325c5f-a0fd-4a10-b026-abe93e3e5963"
        dimension_name              = "Starter-Monthly-Licensed-User-Subscription"
        dimension_description       = "Starter-Monthly-Licensed-User-Subscription"
        metric_type                 = "EACH"
        dimension_billing_frequency = "ANNUAL"
        included_quantity           = 98

        rates {
          currency = "USD"
          rate     = 87
        }
      }

      dimensions {
        dimension_key               = "373936ee-23cc-4f06-92bc-4e8f926714e8"
        dimension_name              = "Starter-Output-Token-Generation-Volume"
        dimension_description       = "Starter-Output-Token-Generation-Volume"
        metric_type                 = "EACH"
        dimension_billing_frequency = "ANNUAL"
        included_quantity           = 909

        rates {
          currency = "USD"
          rate     = 768
        }
      }
    }
  }
}

data "oci_self_subscriptions" "test_subscriptions" {
  compartment_id = var.compartment_id
  display_name   = oci_self_subscription.test_subscription.display_name
}
