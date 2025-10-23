// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_fleet_crypto_analysis_results" "test_fleet_crypto_analysis_results" {
  #Required
  fleet_id = var.fleet_ocid

  #Optional
  aggregation_mode                         = "JFR"
  managed_instance_id                      = var.managed_instance_ocid
  host_name                                = var.host_name
  finding_count                            = 10
  finding_count_greater_than               = 10
  non_compliant_finding_count              = 10
  non_compliant_finding_count_greater_than = 10
  time_end                                 = var.time_end
  time_start                               = var.time_start
}

data "oci_jms_fleet_crypto_analysis_result" "test_fleet_crypto_analysis_result" {
  #Required
  crypto_analysis_result_id = "example-crypto-analysis-result-id"
  fleet_id                  = var.fleet_ocid
}
