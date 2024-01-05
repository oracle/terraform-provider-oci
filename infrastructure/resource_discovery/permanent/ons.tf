// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "random_string" "topicname" {
  length  = 10
  special = false
}

resource "oci_ons_notification_topic" "notification_topic_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  name           = "${random_string.topicname.result}"
}

resource "oci_ons_subscription" "subscription_email_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  endpoint       = "RobotNotExistRd@oracle.com"
  protocol       = "EMAIL"
  topic_id       = "${oci_ons_notification_topic.notification_topic_rd.id}"
}
