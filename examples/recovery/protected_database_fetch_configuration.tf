// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "protected_database_fetch_configuration_configuration_type" {
  default = "ALL"
}

data "oci_recovery_protected_database_fetch_configuration" "test_protected_database_fetch_configuration" {
  #Required
  protected_database_id = oci_recovery_protected_database.test_protected_database.id

  #Optional
  configuration_type    = var.protected_database_fetch_configuration_configuration_type
  base64_encode_content = true
}

resource "local_sensitive_file" "downloaded_configuration" {
  filename       = "${path.module}/configuration.zip"
  content_base64 = data.oci_recovery_protected_database_fetch_configuration.test_protected_database_fetch_configuration.content
}
