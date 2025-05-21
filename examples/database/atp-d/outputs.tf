output "autonomous_database_admin_password" {
  value = random_string.autonomous_database_admin_password.result
}

output "autonomous_database_high_connection_string" {
  value = lookup(
    oci_database_autonomous_database.test_autonomous_database.connection_strings[0].all_connection_strings,
    "high",
    "unavailable",
  )
}

output "autonomous_databases" {
  value = data.oci_database_autonomous_databases.autonomous_databases.autonomous_databases
}

output "autonomous_container_databases" {
  value = data.oci_database_autonomous_container_databases.test_autonomous_container_databases.autonomous_container_databases
}