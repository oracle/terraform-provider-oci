// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

output "key_id" {
  value = oci_kms_key.test_key.id
}

output "volumes" {
  value = data.oci_core_volumes.test_volumes.volumes
}

