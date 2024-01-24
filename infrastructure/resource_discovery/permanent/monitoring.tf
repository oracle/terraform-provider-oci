// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_monitoring_alarm" "alarm_rd" {
  #Required
  compartment_id        = "${var.compartment_ocid}"
  destinations          = ["${oci_ons_notification_topic.notification_topic_rd.id}"]
  display_name          = "${var.alarm_display_name}"
  is_enabled            = "${var.alarm_is_enabled}"
  metric_compartment_id = "${var.compartment_ocid}"
  namespace             = "${var.alarm_namespace}"
  query                 = "${var.alarm_query}"
  severity              = "${var.alarm_severity}"

  #Optional
  body                             = "${var.alarm_body}"
  metric_compartment_id_in_subtree = "${var.alarm_metric_compartment_id_in_subtree}"
  pending_duration                 = "${var.alarm_pending_duration}"
  repeat_notification_duration     = "${var.alarm_repeat_notification_duration}"
  resolution                       = "${var.alarm_resolution}"
  resource_group                   = "${var.alarm_resource_group}"

  suppression {
    #Required
    time_suppress_from  = "${var.alarm_suppression_time_suppress_from}"
    time_suppress_until = "${var.alarm_suppression_time_suppress_until}"

    #Optional
    description = "${var.alarm_suppression_description}"
  }
}
