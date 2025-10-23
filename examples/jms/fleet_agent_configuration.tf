// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_jms_fleet_agent_configuration" "test_fleet_agent_configuration" {
  #Required
  fleet_id = var.fleet_ocid

  #Optional
  agent_polling_interval_in_minutes                  = 10
  is_capturing_ip_address_and_fqdn_enabled           = false
  is_collecting_managed_instance_metrics_enabled     = false
  is_collecting_usernames_enabled                    = false
  is_libraries_scan_enabled                          = false
  java_usage_tracker_processing_frequency_in_minutes = 10
  jre_scan_frequency_in_minutes                      = 180  # must be >= 180
  linux_configuration {
    #Required
    exclude_paths = ["/user/private1", "/opt/private1"]
    include_paths = ["/user", "/opt"]
  }
  mac_os_configuration {
    #Required
    exclude_paths = ["/home/private1"]
    include_paths = ["/home"]
  }
  windows_configuration {
    #Required
    exclude_paths = ["c:\\windows\\private1", "d:\\data\\private1"]
    include_paths = ["c:\\windows", "d:\\data"]
  }
  work_request_validity_period_in_days = 10
}

data "oci_jms_fleet_agent_configuration" "test_fleet_agent_configuration" {
  #Required
  fleet_id = var.fleet_ocid
}
