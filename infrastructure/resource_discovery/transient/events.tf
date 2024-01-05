// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_events_rule" "events_rule_rd" {
  #Required
  actions {
    #Required
    actions {
      #Required
      action_type = "ONS"
      is_enabled  = true

      #Optional
      description = "description"
      topic_id    = "${oci_ons_notification_topic.test_notification_topic_events_rd.id}"
    }

    actions {
      #Required
      action_type = "OSS"
      is_enabled  = true

      #Optional
      description = "description"
      stream_id   = "${oci_streaming_stream.stream_rd.id}"
    }

    actions {
      #Required
      action_type = "FAAS"
      is_enabled  = true

      #Optional
      description = "description"
      function_id = "${oci_functions_function.functions_function_rd.id}"
    }
  }

  compartment_id = "${var.compartment_ocid}"
  condition      = "{\"eventType\": \"com.oraclecloud.dbaas.autonomous.database.backup.end\"}"
  description    = "description"
  display_name   = "eventsRuleRD"
  is_enabled     = true
}

data "oci_events_rules" "test_events_rules_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "testEventsRulesRD"
  state        = "ACTIVE"
}

resource "oci_ons_notification_topic" "test_notification_topic_events_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  name           = "testNotificationTopicEventsRD"
}
