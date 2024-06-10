variable "autonomous_virtual_machine_state" {
  default = "AVAILABLE"
}

data "oci_database_autonomous_virtual_machines" "test_autonomous_virtual_machines" {
  #Required
  autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
  compartment_id           = var.compartment_ocid

  #Optional
  state = var.autonomous_virtual_machine_state
}
