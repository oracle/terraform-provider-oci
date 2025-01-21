// use oci_database_db_node_snapshot_management resource to create dbnode snapshot(s)
resource "oci_database_db_node_snapshot_management" "test_db_node_snapshot_management" {
  #Required
  exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  source_dbnode_ids = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.node_resource[*].node_id
  name                = "snapshot-tf-test"
  #Optional
  defined_tags = {
    "example-tag-namespace-all.example-tag" = "value"
  }
  freeform_tags = {
    "Department" = "Finance"
  }
  lifecycle {
    ignore_changes = [defined_tags, freeform_tags]
  }
}

// use oci_database_db_node_snapshot resource to mount / unmount / delete dbnode snapshot
resource "oci_database_db_node_snapshot" "test_db_node_snapshot" {
  #Required
  dbnode_snapshot_id = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].id
  mount_dbnode_id    = "null"
}

data "oci_database_db_node_snapshots" "test_db_node_snapshots" {
  #Required
  compartment_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.compartment_id

  #Optional
  cluster_id       = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.exadb_vm_cluster_id
  name             = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].name
  source_dbnode_id = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].source_dbnode_id
  state            = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].state

  ## Example: filter db_node_snapshots by name
  #filter {
  #  name  = "name"
  #  regex = true
  #  values = ["^\\w+-${oci_database_db_node_snapshot_management.test_db_node_snapshot_management.name}$"]
  #}
  #
  ## Example: Get all but Terminated db_node_snapshots
  #filter {
  #  name = "state"
  #  values = ["CREATING", "AVAILABLE", "FAILED", "MOUNTED", "MOUNTING", "UNMOUNTING"]
  #}
}

data "oci_database_db_node_snapshot" "test_db_node_snapshot" {
  #Required
  dbnode_snapshot_id = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].id
}