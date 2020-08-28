// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_database_autonomous_vm_cluster" "test_autonomous_vm_cluster" {
  #Required
  compartment_id            = var.compartment_id
  display_name              = "autonomousVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network.id

  #Optional
  is_local_backup_enabled = "false"
  license_model           = "LICENSE_INCLUDED"
  time_zone               = "US/Pacific"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  }

  freeform_tags = {
    "Department" = "Finance"
  }
}

data "oci_database_autonomous_vm_clusters" "test_autonomous_vm_clusters" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name              = "autonomousVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  state                     = "AVAILABLE"
}

