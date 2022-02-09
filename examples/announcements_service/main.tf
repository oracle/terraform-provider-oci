// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "announcement_subscription_defined_tags_value" {
  default = "value"
}

variable "announcement_subscription_description" {
  default = "description"
}

variable "announcement_subscription_display_name" {
  default = "displayName"
}

variable "announcement_subscription_filter_groups_filters_type" {
  default = "COMPARTMENT_ID"
}

variable "announcement_subscription_filter_groups_filters_value" {
  default = "value"
}

variable "announcement_subscription_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "announcement_subscription_id" {
  default = "id"
}

variable "announcement_subscription_state" {
  default = "ACTIVE"
}

variable "announcement_subscriptions_filter_group_filters_type" {
  default = "PLATFORM_TYPE"
}

variable "announcement_subscriptions_filter_group_filters_value" {
  default = "IAAS"
}

variable "announcement_subscriptions_filter_group_name" {
  default = "fg-name"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

// Announcement Subscription Resource
resource "oci_announcements_service_announcement_subscription" "test_announcement_subscription" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.announcement_subscription_display_name
  ons_topic_id   = oci_ons_notification_topic.test_notification_topic.id

  #Optional
  defined_tags = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.announcement_subscription_defined_tags_value)
  description  = var.announcement_subscription_description
  filter_groups {
    #Required
    filters {
      #Required
      type  = var.announcement_subscription_filter_groups_filters_type
      value = var.announcement_subscription_filter_groups_filters_value
    }
  }
  freeform_tags = var.announcement_subscription_freeform_tags
}

// Announcement Subscription Change Compartment Resource
resource "oci_announcements_service_announcement_subscriptions_actions_change_compartment" "test_announcement_subscriptions_actions_change_compartment" {
  #Required
  announcement_subscription_id = oci_announcements_service_announcement_subscription.test_announcement_subscription.id
  compartment_id               = var.compartment_id
}

// Announcement Subscription Filter Group Resource
resource "oci_announcements_service_announcement_subscriptions_filter_group" "test_announcement_subscriptions_filter_group" {
  #Required
  announcement_subscription_id = oci_announcements_service_announcement_subscription.test_announcement_subscription.id
  filters {
    #Required
    type  = var.announcement_subscriptions_filter_group_filters_type
    value = var.announcement_subscriptions_filter_group_filters_value
  }
  name = var.announcement_subscriptions_filter_group_name
}

data "oci_announcements_service_announcement_subscriptions" "test_announcement_subscriptions" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.announcement_subscription_display_name
  id           = var.announcement_subscription_id
  state        = var.announcement_subscription_state
}
