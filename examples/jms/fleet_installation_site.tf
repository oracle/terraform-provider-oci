// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_fleet_installation_sites" "test_fleet_installation_sites" {
  #Required
  fleet_id = var.fleet_ocid

  #Optional
  application_id      = "example-application-id"
  installation_path   = "/example/java/installation/path"
  path_contains       = "example"
  jre_version         = "17.0.0"
  jre_vendor          = "example"
  jre_distribution    = "example"
  jre_security_status = "UNKNOWN"
  managed_instance_id = var.managed_instance_ocid
  os_family           = []
  time_end            = var.time_end
  time_start          = var.time_start
}
