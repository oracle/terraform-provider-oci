data "oci_identity_availability_domains" "test_availability_domain" {
  compartment_id = var.compartment_id
}

data "oci_database_cloud_vm_clusters" "test_cloud_vm_cluster" {
  compartment_id = var.compartment_id
  filter {
    name   = "id"
    values = [oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id]
  }
}

data "oci_database_db_homes" "test_db_home" {
  compartment_id = var.compartment_id
  vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id
  filter {
    name   = "db_system_id"
    values = [oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id]
  }
}

data "oci_database_databases" "test_databases" {
  compartment_id = var.compartment_id
  db_home_id = oci_database_db_home.test_db_home.id
  filter {
    name   = "db_name"
    values = [oci_database_database.primary_database.db_name]
  }
}

data "oci_database_database" "test_database" {
  database_id = oci_database_database.primary_database.id
}
