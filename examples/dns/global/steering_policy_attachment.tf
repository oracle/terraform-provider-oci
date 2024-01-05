// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "steering_policy_attachment_display_name" {
  default = "Test-Steering-Policy-Attachment"
}

variable "steering_policy_attachment_state" {
  default = "ACTIVE"
}

variable "steering_policy_attachment_time_created_greater_than_or_equal_to" {
  default = "2018-01-01T00:00:00.000Z"
}

variable "steering_policy_attachment_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}

resource "oci_dns_steering_policy_attachment" "test_steering_policy_attachment" {
  #Required
  domain_name        = oci_dns_record.record-a.domain
  steering_policy_id = oci_dns_steering_policy.test_steering_policy.id
  zone_id            = oci_dns_zone.zone1.id

  #Optional
  display_name = var.steering_policy_attachment_display_name
}

data "oci_dns_steering_policy_attachments" "test_steering_policy_attachments" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_dns_steering_policy_attachment.test_steering_policy_attachment.display_name
  domain       = oci_dns_steering_policy_attachment.test_steering_policy_attachment.domain_name

  #domain_contains                       = oci_dns_steering_policy_attachment.test_steering_policy_attachment.domain_name
  id                                    = oci_dns_steering_policy_attachment.test_steering_policy_attachment.id
  state                                 = var.steering_policy_attachment_state
  steering_policy_id                    = oci_dns_steering_policy.test_steering_policy.id
  time_created_greater_than_or_equal_to = var.steering_policy_attachment_time_created_greater_than_or_equal_to
  time_created_less_than                = var.steering_policy_attachment_time_created_less_than
  zone_id                               = oci_dns_zone.zone2.id
}

