// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_kms_verify" "test_verify" {
  crypto_endpoint      = var.crypto_endpoint
  key_id               = var.test_rsa_key_id
  key_version_id       = var.test_rsa_key_version
  message              = "message"
  signing_algorithm    = "SHA_224_RSA_PKCS1_V1_5"
  signature            = oci_kms_sign.test_sign.signature
}
