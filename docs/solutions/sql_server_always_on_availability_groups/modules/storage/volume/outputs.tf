output "sql_db_id" {
  value = ["${oci_core_volume.db_block.*.id}"]
}

output "sql_log_id" {
  value = ["${oci_core_volume.db_log.*.id}"]
}

output "sql_backup_id" {
  value = ["${oci_core_volume.db_backup.*.id}"]
}

output "witness_id" {
  value = ["${oci_core_volume.witness_volume.*.id}"]
}
