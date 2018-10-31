output "key_id" {
  value = "${oci_kms_key.test_key.id}"
}

output "volumes" {
  value = "${data.oci_core_volumes.test_volumes.volumes}"
}
