data "oci_database_autonomous_vm_clusters" "test_autonomous_vm_clusters" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name              = "autonomousVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  state                     = "AVAILABLE"
}

data "oci_database_exadata_infrastructures" "test_exadata_infrastructures" {
  #Required
  compartment_id = var.compartment_ocid
}

