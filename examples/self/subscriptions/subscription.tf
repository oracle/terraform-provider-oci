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
  display_name   = "SELF-DEXREQ-9206-Test-${formatdate("YYYYMMDDhhmmss", timestamp())}"
  source_type    = "THIRD_PARTY"
  realm          = "OC1"
  region         = var.region
  subscription_details {
    pricing_plan {
      plan_type         = "USAGE_BASED"
      plan_name         = "p1"
      plan_description  = "Fixes to be done\n\nConsumer screens -\n1. Description should be multi line for all screens for plan and dimensions, to avoid horizontal scrolling the description. If the desciption is big word wrap should be done.\n2. Fix the size of the columns, if description is big then other columns are getting shrinked.\n3. Price should be 100 USD , right now $ is coming -> Check on all Screens\n4. Review the content for all the pages\n5. After selecting the plan, Plan Details is visible. We might have to check whether it is required or not. -> Can be done between LA and GA\n6. Remove additional space above plan details and actual data\n7. After removing search and filter there is a big white space coming. Need to check on this\n\nCreate Subscription screen -\n1. No error is coming when without selecting plan , and clicking Subscribe button.\n2. Select one pricing plan by default.\n3. On viewing plan details -\n  i.  Plan should be Plan Name\n  ii. Order should be Plan Name, Description, Base Price, Type, Billing Frequency, Plan Duration\n4. Heading should be Create subscription instead of Subscription followed by listing\n5. Remove metric type from dimension\n6. Required place holder should be consistent for all pages\n7. Auto Renew button is not present\n8. If there are 2 terms and condition - 1 partner and other oracle, right now only 1 is visible\n\nSubscription details -\n1. Compartment name is not present\n2. Change the warning to info\n3. Subscription Plan can be rename to Plan details\n4. Order of fields should be same as that of create subscription\n5. Subscription with Pending Cancellation status is not visible\n6. Duration is not consistent with other screen, right now 1 year is coming\n7. Auto Renew button is not present\n8. Cancellation button is missing -> At a time only button should be visible either activate or cancellation\n9. More Subscription details should be present like billing frequency etc.\n\nSubscriptions (List Page) -\n1. Plan Type should be added in column and filter\n2. Subscription name should be word wrap or 3 dots should come\n3. Statuses in filtering should be user friendly and should show only relevant not all.\n4. Annual Price should be only Price and value should come\n5. Plan should be changed to Plan Name\n5. Order -> Subscription Name, Status, Type, Plan Name, Price, rest same\n\nBE -\n1. Subscription should be activated immediately\n\nCheck with Kartik\n1. How to show Semi annual\n2. When viewing Plan details table layout can be improved -> instead of vertical (multiple rows, 2 columns) can be horizontal (1 row, multiple columns)"
      billing_frequency = "ANNUAL"
      plan_duration     = "ANNUAL"
      rates {
        currency = "USD"
        rate     = 0
      }
      dimensions {
        dimension_key               = "5327465d-2151-49c6-8eac-e671170660e9"
        dimension_name              = "d1"
        dimension_description       = "d1"
        metric_type                 = "EACH"
        dimension_billing_frequency = "ANNUAL"
        rates {
          currency = "USD"
          rate     = 2.999
        }
      }
      dimensions {
        dimension_key               = "79fc947a-2605-4606-9af4-aabee779c3c0"
        dimension_name              = "d333"
        dimension_description       = "Fixes to be done\n\nConsumer screens -\n1. Description should be multi line for all screens for plan and dimensions, to avoid horizontal scrolling the description. If the desciption is big word wrap should be done.\n2. Fix the size of the columns, if description is big then other columns are getting shrinked.\n3. Price should be 100 USD , right now $ is coming -> Check on all Screens\n4. Review the content for all the pages\n5. After selecting the plan, Plan Details is visible. We might have to check whether it is required or not. -> Can be done between LA and GA\n6. Remove additional space above plan details and actual data\n7. After removing search and filter there is a big white space coming. Need to check on this\n\nCreate Subscription screen -\n1. No error is coming when without selecting plan , and clicking Subscribe button.\n2. Select one pricing plan by default.\n3. On viewing plan details -\n  i.  Plan should be Plan Name\n  ii. Order should be Plan Name, Description, Base Price, Type, Billing Frequency, Plan Duration\n4. Heading should be Create subscription instead of Subscription followed by listing\n5. Remove metric type from dimension\n6. Required place holder should be consistent for all pages\n7. Auto Renew button is not present\n8. If there are 2 terms and condition - 1 partner and other oracle, right now only 1 is visible\n\nSubscription details -\n1. Compartment name is not present\n2. Change the warning to info\n3. Subscription Plan can be rename to Plan details\n4. Order of fields should be same as that of create subscription\n5. Subscription with Pending Cancellation status is not visible\n6. Duration is not consistent with other screen, right now 1 year is coming\n7. Auto Renew button is not present\n8. Cancellation button is missing -> At a time only button should be visible either activate or cancellation\n9. More Subscription details should be present like billing frequency etc.\n\nSubscriptions (List Page) -\n1. Plan Type should be added in column and filter\n2. Subscription name should be word wrap or 3 dots should come\n3. Statuses in filtering should be user friendly and should show only relevant not all.\n4. Annual Price should be only Price and value should come\n5. Plan should be changed to Plan Name\n5. Order -> Subscription Name, Status, Type, Plan Name, Price, rest same\n\nBE -\n1. Subscription should be activated immediately\n\nCheck with Kartik\n1. How to show Semi annual\n2. When viewing Plan details table layout can be improved -> instead of vertical (multiple rows, 2 columns) can be horizontal (1 row, multiple columns)"
        metric_type                 = "EACH"
        dimension_billing_frequency = "ANNUAL"
        rates {
          currency = "USD"
          rate     = 3.999999
        }
      }
      dimensions {
        dimension_key               = "225ec907-a883-4e78-b6f6-eea07108b354"
        dimension_name              = "dee"
        dimension_description       = "Fixes to be done\n\nConsumer screens -\n1. Description should be multi line for all screens for plan and dimensions, to avoid horizontal scrolling the description. If the desciption is big word wrap should be done.\n2. Fix the size of the columns, if description is big then other columns are getting shrinked.\n3. Price should be 100 USD , right now $ is coming -> Check on all Screens\n4. Review the content for all the pages\n5. After selecting the plan, Plan Details is visible. We might have to check whether it is required or not. -> Can be done between LA and GA\n6. Remove additional space above plan details and actual data\n7. After removing search and filter there is a big white space coming. Need to check on this\n\nCreate Subscription screen -\n1. No error is coming when without selecting plan , and clicking Subscribe button.\n2. Select one pricing plan by default.\n3. On viewing plan details -\n  i.  Plan should be Plan Name\n  ii. Order should be Plan Name, Description, Base Price, Type, Billing Frequency, Plan Duration\n4. Heading should be Create subscription instead of Subscription followed by listing\n5. Remove metric type from dimension\n6. Required place holder should be consistent for all pages\n7. Auto Renew button is not present\n8. If there are 2 terms and condition - 1 partner and other oracle, right now only 1 is visible\n\nSubscription details -\n1. Compartment name is not present\n2. Change the warning to info\n3. Subscription Plan can be rename to Plan details\n4. Order of fields should be same as that of create subscription\n5. Subscription with Pending Cancellation status is not visible\n6. Duration is not consistent with other screen, right now 1 year is coming\n7. Auto Renew button is not present\n8. Cancellation button is missing -> At a time only button should be visible either activate or cancellation\n9. More Subscription details should be present like billing frequency etc.\n\nSubscriptions (List Page) -\n1. Plan Type should be added in column and filter\n2. Subscription name should be word wrap or 3 dots should come\n3. Statuses in filtering should be user friendly and should show only relevant not all.\n4. Annual Price should be only Price and value should come\n5. Plan should be changed to Plan Name\n5. Order -> Subscription Name, Status, Type, Plan Name, Price, rest same\n\nBE -\n1. Subscription should be activated immediately\n\nCheck with Kartik\n1. How to show Semi annual\n2. When viewing Plan details table layout can be improved -> instead of vertical (multiple rows, 2 columns) can be horizontal (1 row, multiple columns)"
        metric_type                 = "EACH"
        dimension_billing_frequency = "ANNUAL"
        rates {
          currency = "USD"
          rate     = 108
        }
      }
    }
    partner_registration_url = "https://www.google.com"
    currency                 = "USD"
    amount                   = 0
    billing_details {
      pricing_plan_key = "4f787c1e-1c83-4a78-901c-24e74f4c4b5b"
      sku              = "MP10327"
      metric_type      = "INSTANCE_HOURS"
      billing_model    = "USAGE_BASED"
      rate_allocation  = 1
      has_gov_sku      = false
      meters {
        name            = "MP_INSP_MP"
        rate_allocation = 1
      }
    }
    is_auto_renew = true
  }
}

data "oci_self_subscriptions" "test_subscriptions" {
  compartment_id = var.compartment_id
  display_name   = oci_self_subscription.test_subscription.display_name
}
