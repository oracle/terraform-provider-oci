// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "public_key" {
  default = "-----BEGIN PUBLIC KEY-----MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuYNxKqyNSTPApIVh1xiR3914Q8Ex+goi8kbMUjMa/b47A12SGdh18SAsZTTkld09MGhIswyv2Eln5MQKyupf646zk0E0kxH4llpfSAtUEaa5bxRXhko5BejvimMy4hCMn+kYkzAre7CoAw97rZ96L+TgkqdtwYXl0JzE4xYwfM7OqkH9/3TIeiX4q8kVDi0CsHMGbBo4gMIIunLoEn27ej/Vm6Nbkgl8AnJaWZq8gG8y6ojDLrJhnTK4IVYZ3XYx2uxz/E5VcjMaTdWVjKVCS4F2yK9hFbL1G2KDDh8k3G7dFDFwGI6qxwidbZW7JtcXQWu0Qx0tBNdB28VlsDWZEQIDAQAB-----END PUBLIC KEY-----"
}

variable "scope" {
  default = "urn:oracle:db::id::*"
}

resource "oci_identity_data_plane_generate_scoped_access_token" "test_scoped_access_token" {
  #Required
  public_key = var.public_key
  scope    = var.scope
}