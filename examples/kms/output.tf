// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

output "key_id" {
  value = "${oci_kms_key.test_key.id}"
}

output "volumes" {
  value = "${data.oci_core_volumes.test_volumes.volumes}"
}
