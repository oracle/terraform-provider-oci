// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "notification_setting_authorization" {
  default = "authorization"
}

variable "notification_setting_event_settings_enabled" {
  default = false
}

variable "notification_setting_from_email_address_display_name" {
  default = "displayName"
}

variable "notification_setting_from_email_address_validate" {
  default = "email"
}

variable "notification_setting_from_email_address_value" {
  default = "value@email.com"
}

variable "notification_setting_notification_enabled" {
  default = false
}

variable "notification_setting_send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email" {
  default = false
}

variable "notification_setting_send_notifications_to_secondary_email" {
  default = false
}

variable "notification_setting_tags_key" {
  default = "key"
}

variable "notification_setting_tags_value" {
  default = "value"
}

variable "notification_setting_test_mode_enabled" {
  default = false
}

variable "notification_setting_test_recipients" {
  default = []
}


resource "oci_identity_domains_notification_setting" "test_notification_setting" {
  #Required
  event_settings {
    #Required
    event_id = "admin.user.create.success"

    #Optional
    enabled = var.notification_setting_event_settings_enabled
  }
  from_email_address {
    #Required
    validate = var.notification_setting_from_email_address_validate
    value    = var.notification_setting_from_email_address_value

    #Optional
    display_name = var.notification_setting_from_email_address_display_name
  }
  idcs_endpoint           = data.oci_identity_domain.test_domain.url
  notification_enabled    = var.notification_setting_notification_enabled
  notification_setting_id = "NotificationSettings"
  schemas                 = ["urn:ietf:params:scim:schemas:oracle:idcs:NotificationSettings"]

  #Optional
  attribute_sets                                                                   = ["all"]
  attributes                                                                       = ""
  authorization                                                                    = var.notification_setting_authorization
  external_id                                                                      = "externalId"
  #use the latest if not provided
  # resource_type_schema_version = var.notification_setting_resource_type_schema_version
  send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email = var.notification_setting_send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email
  send_notifications_to_secondary_email                                            = var.notification_setting_send_notifications_to_secondary_email
  tags {
    #Required
    key   = var.notification_setting_tags_key
    value = var.notification_setting_tags_value
  }
  test_mode_enabled = var.notification_setting_test_mode_enabled
  test_recipients   = var.notification_setting_test_recipients
}

data "oci_identity_domains_notification_settings" "test_notification_settings" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.notification_setting_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.notification_setting_resource_type_schema_version
}

