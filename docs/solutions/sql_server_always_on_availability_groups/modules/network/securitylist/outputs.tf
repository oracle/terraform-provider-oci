output "dmz_id" {
  value = "${oci_core_security_list.dmz.id}"
}

output "admin_id" {
  value = "${oci_core_security_list.admin.id}"
}

output "sql_id" {
  value = "${oci_core_security_list.sql.id}"
}
