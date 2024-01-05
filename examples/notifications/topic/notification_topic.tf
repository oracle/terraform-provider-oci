// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "random_string" "topicname" {
  length  = 10
  special = false
}

resource "oci_ons_notification_topic" "test_notification_topic" {
  #Required
  compartment_id = var.compartment_ocid
  name           = random_string.topicname.result

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.notification_topic_defined_tags_value
  }
  description   = var.notification_topic_description
  freeform_tags = var.notification_topic_freeform_tags
}

data "oci_ons_notification_topics" "test_notification_topics" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  id    = oci_ons_notification_topic.test_notification_topic.id
  name  = oci_ons_notification_topic.test_notification_topic.name
  state = var.notification_topic_state
}

