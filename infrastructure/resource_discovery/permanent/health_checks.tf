// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_health_checks_http_monitor" "http_monitor_rd" {
  #Required
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "httpMonitorRD"
  interval_in_seconds = "${var.http_monitor_interval_in_seconds}"
  protocol            = "${var.http_monitor_protocol}"
  targets             = "${var.http_monitor_targets}"

  #Optional
  freeform_tags       = "${var.http_monitor_freeform_tags}"
  is_enabled          = "${var.http_monitor_is_enabled}"
  method              = "${var.http_monitor_method}"
  path                = "${var.http_monitor_path}"
  port                = "${var.http_monitor_port}"
  timeout_in_seconds  = "${var.http_monitor_timeout_in_seconds}"
  vantage_point_names = "${var.http_monitor_vantage_point_names}"
}

resource "oci_health_checks_ping_monitor" "ping_monitor_rd" {
  #Required
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "pingMonitorRD"
  interval_in_seconds = "${var.ping_monitor_interval_in_seconds}"
  protocol            = "${var.ping_monitor_protocol}"
  targets             = "${var.ping_monitor_targets}"

  #Optional
  freeform_tags = "${var.ping_monitor_freeform_tags}"

  is_enabled          = "${var.ping_monitor_is_enabled}"
  port                = "${var.ping_monitor_port}"
  timeout_in_seconds  = "${var.ping_monitor_timeout_in_seconds}"
  vantage_point_names = "${var.ping_monitor_vantage_point_names}"
}
