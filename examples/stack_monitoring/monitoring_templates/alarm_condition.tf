// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_stack_monitoring_monitoring_template_alarm_condition" "test_monitoring_template_alarm_condition_example" {
  # Required
  condition_type = "FIXED"
  conditions {
    body = "Garbage collection throughput exceeds the warning threshold value"
    query = "GarbageCollectionThroughput[10m].mean() > 0.3"
    severity = "WARNING"
  }
  metric_name = "GarbageCollectionThroughput"
  monitoring_template_id = oci_stack_monitoring_monitoring_template.test_monitoring_template_example.id
  namespace = "oracle_appmgmt"
  resource_type = "ocid1.stackmonitoringresourcetype.apache_tomcat"
}

data "oci_stack_monitoring_monitoring_template_alarm_condition" "test_monitoring_template_alarm_condition_example" {
  # Required
  alarm_condition_id = oci_stack_monitoring_monitoring_template_alarm_condition.test_monitoring_template_alarm_condition_example.id
  monitoring_template_id = oci_stack_monitoring_monitoring_template_alarm_condition.test_monitoring_template_alarm_condition_example.monitoring_template_id
}
