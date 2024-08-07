// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_agent_installers" "test_agent_installers" {
  #either one of these are necessary
  compartment_id        = var.compartment_ocid
  fleet_id              = var.fleet_ocid

  #optional
  os_family             = "LINUX"
  platform_architecture = "X86_64"
}
