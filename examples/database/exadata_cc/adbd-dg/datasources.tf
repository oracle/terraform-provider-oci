############################
#### Begin Data Sources ####

# get/list all exadata infrastructures in a compartment
data "oci_database_exadata_infrastructures" "all_exadata_infrastructures" {
  #Required
  compartment_id = var.compartment_ocid
}

# get/list all autonomous vm clusters in a compartment
data "oci_database_autonomous_vm_clusters" "all_autonomous_vm_clusters" {
  #Required
  compartment_id = var.compartment_ocid
}

# get exadata infrastructure with given id
data "oci_database_exadata_infrastructure" "primary_exadata_infrastructure" {
  #Required
  exadata_infrastructure_id = oci_database_exadata_infrastructure.primary_exadata_infrastructure.id
}

data "oci_database_exadata_infrastructure" "standby_exadata_infrastructure" {
  #Required
  exadata_infrastructure_id = oci_database_exadata_infrastructure.standby_exadata_infrastructure.id
}

# get autonomous vm cluster with given id
data "oci_database_autonomous_vm_cluster" "primary_autonomous_vm_cluster" {
  #Required
  autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.primary_autonomous_vm_cluster.id
}

data "oci_database_autonomous_vm_cluster" "standby_autonomous_vm_cluster" {
  #Required
  autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.standby_autonomous_vm_cluster.id
}

# get/list all ACDs in a given AVM and compartment id
data "oci_database_autonomous_container_databases" "primary_autonomous_vm_cluster_acds" {
  compartment_id           = var.compartment_ocid
  autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.primary_autonomous_vm_cluster.id
}

data "oci_database_autonomous_container_database" "primary_autonomous_container_database" {
  autonomous_container_database_id = oci_database_autonomous_container_database.dg_autonomous_container_database.id
}

data "oci_database_autonomous_database" "parimary_adb" {
  autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
}

data "oci_database_autonomous_container_database" "standby_autonomous_container_database" {
  autonomous_container_database_id = data.oci_database_autonomous_container_database_dataguard_associations.primary_autonomous_dg_associations.autonomous_container_database_dataguard_associations[0].peer_autonomous_container_database_id
}

# get/list all dg associations for a given ACD
data "oci_database_autonomous_container_database_dataguard_associations" "primary_autonomous_dg_associations" {
  autonomous_container_database_id = oci_database_autonomous_container_database.dg_autonomous_container_database.id
}

# get dg association
# added depends_on so that we fetch latest dg info after switchover/failover/reinstate operations are done
data "oci_database_autonomous_container_database_dataguard_association" "primary_autonomous_dg_association" {
  autonomous_container_database_id                       = oci_database_autonomous_container_database.dg_autonomous_container_database.id
  autonomous_container_database_dataguard_association_id = data.oci_database_autonomous_container_database_dataguard_associations.primary_autonomous_dg_associations.autonomous_container_database_dataguard_associations[0].id
  depends_on                                             = [oci_database_autonomous_container_database_dataguard_association_operation.reinstate]
}

data "oci_database_autonomous_container_database_dataguard_association" "standby_autonomous_dg_association" {
  autonomous_container_database_id                       = data.oci_database_autonomous_container_database.standby_autonomous_container_database.id
  autonomous_container_database_dataguard_association_id = data.oci_database_autonomous_container_database_dataguard_associations.primary_autonomous_dg_associations.autonomous_container_database_dataguard_associations[0].peer_autonomous_container_database_dataguard_association_id
  depends_on                                             = [oci_database_autonomous_container_database_dataguard_association_operation.reinstate]
}

#### End Data Sources ####
##########################