
output "primary_exadata_infrustructure_id" {
  value = data.oci_database_exadata_infrastructure.primary_exadata_infrastructure.id
}

output "standby_exadata_infrustructure_id" {
  value = data.oci_database_exadata_infrastructure.standby_exadata_infrastructure.id
}

output "primary_autonomous_vm_cluster_id" {
  value = data.oci_database_autonomous_vm_cluster.primary_autonomous_vm_cluster.id
}

output "standby_autonomous_vm_cluster_id" {
  value = data.oci_database_autonomous_vm_cluster.standby_autonomous_vm_cluster.id
}

output "primary_acd_id" {
  value = data.oci_database_autonomous_container_database.primary_autonomous_container_database.id
}

output "standby_acd_id" {
  value = data.oci_database_autonomous_container_database.standby_autonomous_container_database.id
}

output "primary_adb_id" {
  value = data.oci_database_autonomous_database.parimary_adb.id
}

output "primary_autonomous_dg_id" {
  value = data.oci_database_autonomous_container_database_dataguard_association.primary_autonomous_dg_association.id
}

output "standby_autonomous_dg_id" {
  value = data.oci_database_autonomous_container_database_dataguard_association.standby_autonomous_dg_association.id
}

output "primary_autonomous_dg_role" {
  value = data.oci_database_autonomous_container_database_dataguard_association.primary_autonomous_dg_association.role
}

output "standby_autonomous_dg_role" {
  value = data.oci_database_autonomous_container_database_dataguard_association.standby_autonomous_dg_association.role
}