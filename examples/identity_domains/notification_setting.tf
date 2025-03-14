// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "notification_setting_authorization" {
  default = "authorization"
}

variable "notification_setting_event_settings_enabled" {
  default = true
}

variable "notification_setting_from_email_address_display_name" {
  default = "Oracle"
}

variable "notification_setting_from_email_address_validate" {
  default = "domain"
}

variable "notification_setting_from_email_address_value" {
  default = "no-reply@oracle.com"
}

variable "notification_setting_notification_enabled" {
  default = true
}

variable "notification_setting_send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email" {
  default = false
}

variable "notification_setting_send_notifications_to_secondary_email" {
  default = true
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

variable event_settings_list {
  default = [
    # "admin.user.create.success",
    "admin.approval.create.success",
    "admin.workflow.request.expiry",
    "admin.workflow.request.rejected",
    "admin.me.register.success",
    "admin.me.register.activation.required",
    "admin.user.federated.create.success",
    "admin.user.authentication.delegated.create.success",
    "admin.user.initiate.activation.success",
    "admin.user.authentication.delegated.initiate.activation.success",
    "admin.me.password.reset.request.success",
    "admin.me.email.verify.recovery.success",
    "admin.me.email.verify.primary.success",
    "admin.me.email.verify.secondary.success",
    "admin.me.email.update.recovery.success",
    "admin.me.email.update.primary.success",
    "admin.me.email.update.secondary.success",
    "admin.me.password.change.success",
    "admin.me.password.reset.success",
    "admin.user.password.change.success",
    "admin.user.password.reset.success",
    "job.running",
    "job.cancelled",
    "job.succeeded",
    "job.failed",
    "admin.quota.exceeded",
    "admin.user.activated.success",
    "admin.user.deactivated.success",
    "admin.me.locked.success",
    "admin.me.passwordrecoverylocked.success",
    "admin.me.unlocked.success",
    "admin.user.update.success",
    "admin.user.replace.success",
    "admin.domain.validation.create.success",
    "admin.email.validation.create.success",
    "authentication.request",
    "authentication.factor.enrollment.request",
    "admin.me.mfa.locked.success",
    "admin.me.mfa.federated.locked.success",
    "authentication.bypasscode.notification",
    "admin.login.to.enable.kerberos.authentication.request",
    "admin.me.request.submit.success",
    "admin.me.request.fulfillment.success",
    "admin.me.request.fulfillment.failure",
    "authentication.emailotp.notification",
    "authentication.phonecall.notification",
    "admin.adaptive.newdevice.detected",
    "job.managed.object.sync",
    "admin.bridge.unreachable",
    "admin.bridge.reachable",
    "admin.bridge.newbinary.available",
    "admin.user.email.verify.primary.success",
    "admin.user.email.verify.recovery.success",
    "admin.user.email.verify.secondary.success",
    "authentication.email.link.notification",
    "admin.idbridge.sync.success",
    "admin.idbridge.sync.failure",
    "admin.domain.create.secondary.success",
    "admin.user.password.propagation.failure",
    "saml.sp.signing.cert.expiration.warning",
    "saml.idp.signing.cert.expiration.warning",
    "admin.ociconsolesignonpolicyconsent.modified",
    "admin.ociconsolesignonpolicyconsent.restored",
    "admin.ociconsolesignonpolicyconsent.autorecord"
  ]
}


resource "oci_identity_domains_notification_setting" "test_notification_setting" {
  #Required
  event_settings {
    #Required
    event_id = "admin.user.create.success"

    #Optional
    enabled = var.notification_setting_event_settings_enabled
  }

  dynamic event_settings {
    for_each = var.event_settings_list

    content {
      event_id = event_settings.value
      enabled  = true
    }
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
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.notification_setting_authorization
  external_id    = "externalId"
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
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.notification_setting_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.notification_setting_resource_type_schema_version
}

