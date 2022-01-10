// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


data "oci_identity_tag_standard_tag_namespace_templates" "test_tag_standard_tag_namespace_template" {
  #Required
  compartment_id = var.compartment_id
}
