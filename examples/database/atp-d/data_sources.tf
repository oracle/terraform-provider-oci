data "oci_identity_availability_domain" "ad" {
  compartment_id = var.compartment_ocid
  ad_number      = 1
}

data "oci_database_cloud_autonomous_vm_cluster" "test_cloud_autonomous_vm_cluster" {
  cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
}

data "oci_database_cloud_autonomous_vm_clusters" "test_cloud_autonomous_vm_clusters" {
  compartment_id = var.compartment_ocid
}

data "oci_database_autonomous_container_databases" "test_autonomous_container_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
  availability_domain                  = data.oci_identity_availability_domain.ad.name
  display_name                         = "example-container-database"
  state                                = "AVAILABLE"
}

data "oci_database_autonomous_databases" "autonomous_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
  display_name                     = oci_database_autonomous_database.test_autonomous_database.display_name
  db_workload                      = "OLTP"
}


data "oci_database_cloud_exadata_infrastructure_un_allocated_resource" "test_cloud_exadata_infrastructure_un_allocated_resources" {
  #Required
  cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id

  #Optional
  db_servers = var.cloud_exadata_infrastructure_un_allocated_resource_db_servers
}

data "oci_database_autonomous_container_database_resource_usage" "test_autonomous_container_database_resource_usages" {
  #Required
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
}

data "oci_database_cloud_autonomous_vm_cluster_acd_resource_usages" "test_cloud_autonomous_vm_cluster_acd_resource_usages" {
  #Required
  cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id

  #Optional
  compartment_id = var.compartment_ocid
}

data "oci_database_cloud_autonomous_vm_cluster_resource_usage" "test_cloud_autonomous_vm_cluster_resource_usages" {
  #Required
  cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
}

data "oci_database_autonomous_database_software_image" "test_autonomous_database_software_image" {
  autonomous_database_software_image_id = oci_database_autonomous_database_software_image.autonomous_database_software_image.id
}