// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_osmanagement_managed_instance_group" "managed_instance_group_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TF-managed-instance-group-rd"

  #Optional
  description = "TF Managed instance group"
}

resource "oci_osmanagement_software_source" "software_source_rd" {
  #Required
  arch_type      = "X86_64"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TF-software-source-rd"

  #Optional
  checksum_type = "SHA1"
  description   = "TF software source"
}
