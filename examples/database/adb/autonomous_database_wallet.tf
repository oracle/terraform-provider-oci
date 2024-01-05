// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "random_string" "autonomous_database_wallet_password" {
  length  = 16
  special = true
}

resource "oci_database_autonomous_database_wallet" "autonomous_database_wallet" {
  autonomous_database_id = oci_database_autonomous_database.autonomous_database.id
  password               = random_string.autonomous_database_wallet_password.result
  base64_encode_content  = "true"
}

resource "local_file" "autonomous_database_wallet_file" {
  content_base64 = oci_database_autonomous_database_wallet.autonomous_database_wallet.content
  filename       = "${path.module}/autonomous_database_wallet.zip"
}

output "autonomous_database_wallet_password" {
  value = random_string.autonomous_database_wallet_password.result
}

