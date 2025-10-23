// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_fleet_drs_files" "test_fleet_drs_files" {
  #Required
  fleet_id = var.fleet_ocid
}

# You need actual ID value for drs_file_key 
# data "oci_jms_fleet_drs_file" "test_fleet_drs_file" {
#   #Required
#   fleet_id     = var.fleet_ocid
#   drs_file_key = "example-drs-file-key"
# }
