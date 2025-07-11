// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

####################### Cloud DB System #########################

variable "compartment_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-compartmentId-Value"
}

variable "cloud_rac1_agent_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-agentId-Value"
}

variable "cloud_rac2_agent_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-agentId-Value"
}

variable "cloud_db_system_display_name" {
  default = "EXAMPLE-displayName-Value"
}

variable "dbaas_dbsystem_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-dbsystemId-Value"
}

variable "dbaas_dbsystem_deployment_type" {
  default = "VM"
}


variable "asm_host_name" {
  default = "EXAMPLE-asm-host-Value"
}

variable "asm_port" {
  default = "1521"
}

variable "asm_service_name" {
  default = "EXAMPLE-asm-service-Value"
}

variable "asm_user_name" {
  default = "EXAMPLE-asm-user-Value"
}

variable "asm_password_secret_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-secretId-Value"
}

variable "asm_credential_name" {
  default = "EXAMPLE-asmCredName-Value"
}

variable "cloud_asm_connector_display_name" {
  default = "asmConnectorName"
}


variable "cloud_cluster_rac1_connector_display_name" {
  default = "cluster1ConnectorName"
}

variable "cloud_cluster_rac2_connector_display_name" {
  default = "cluster2ConnectorName"
}

variable "cloud_listener_connector_display_name" {
  default = "listenerConnectorName"
}

variable "local_listener_name" {
  default = "LISTENER_cloudrac1"
}

variable "managed_databases_asm_property_name" {
  default = "DATA"
}

variable "db_system_discovery_defined_tags_value" {
  default = "db_system_discovery_tag_value"
}

variable "db_system_discovery_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "db_system_defined_tags_value" {
  default = "db_system_tag_value"
}

variable "db_system_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "asm_defined_tags_value" {
  default = "asm_tag_value"
}

variable "asm_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "listener_defined_tags_value" {
  default = "listener_tag_value"
}

variable "listener_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "db_system_connector_defined_tags_value" {
  default = "db_system_connector_tag_value"
}

variable "db_system_connector_freeform_tags" {
  default = { "bar-key" = "value" }
}

# Create a new Tag Namespace.
resource "oci_identity_tag_namespace" "tag_namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "example-tag-namespace-all"
}

# Create a new Tag definition in the above Tag Namespace.
resource "oci_identity_tag" "tag1" {
  #Required
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1.id
}

/*
# Create a new CloudDbSystemDiscovery resource and discover the cloud DB System and its components.
# Also add a connector to the discovered Oracle Container Database (CDB).
resource "oci_database_management_cloud_db_system_discovery" "test_cloud_db_system_discovery" {
  #Required
  agent_id                        = var.cloud_rac1_agent_id
  compartment_id                  = var.compartment_id
  dbaas_parent_infrastructure_id  = var.dbaas_dbsystem_id 
  display_name = var.cloud_db_system_display_name
  deployment_type = var.dbaas_dbsystem_deployment_type

  #Optional
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.db_system_discovery_defined_tags_value
  }
  freeform_tags = var.db_system_discovery_freeform_tags

  # Deselect the Pluggable Databases (PDBs).
  patch_operations {
    operation = "MERGE"
    selection = "discoveredComponents[?componentType == 'DATABASE'].pluggableDatabases[]"
    value {
      is_selected_for_monitoring = "false"
    }
  }
}

# List CloudDbSystemDiscovery resources
data "oci_database_management_cloud_db_system_discoveries" "test_cloud_db_system_discoveries" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  #display_name = 
}

# Create a new CloudDbSystem resource from CloudDbSystemDiscovery resource
resource "oci_database_management_cloud_db_system" "test_cloud_db_system" {
  #Required
  compartment_id         = var.compartment_id
  db_system_discovery_id = oci_database_management_cloud_db_system_discovery.test_cloud_db_system_discovery.id

  #Optional
  database_management_config {
    #Required
    is_enabled = "true"
  }
  display_name = var.cloud_db_system_display_name
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.db_system_defined_tags_value
  }
  freeform_tags = var.db_system_freeform_tags
}
*/

# List CloudDbSystem resources
data "oci_database_management_cloud_db_systems" "test_cloud_db_systems" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.cloud_db_system_display_name
}

# Get CloudDbSystem resource
data "oci_database_management_cloud_db_system" "test_cloud_db_system" {
  #Required
  cloud_db_system_id = var.dbaas_dbsystem_id
}

# List CloudAsms in CloudDbSystem
data "oci_database_management_cloud_asms" "test_cloud_asms" {
  #Required
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id
  #Optional
  compartment_id        = var.compartment_id
}

/* ASM Connector

# Create a new Management Agent based CloudDbSystemConnector
resource "oci_database_management_cloud_db_system_connector" "test_cloud_asm_connector" {
  #Required
  connector_type        = "MACS"
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id
  agent_id              = var.cloud_rac1_agent_id

  #Optional
  display_name = var.cloud_asm_connector_display_name
  connection_info {
    component_type = "ASM"
    connection_credentials {
      credential_name = var.asm_credential_name
      credential_type = "DETAILS"
      password_secret_id = var.asm_password_secret_id
      role = "SYSASM"
      user_name = var.asm_user_name
    }
    connection_string {
      hosts = [var.asm_host_name]
      port = var.asm_port
      protocol = "TCP"
      service = var.asm_service_name
    }
  }
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.db_system_connector_defined_tags_value
  }
  freeform_tags = var.db_system_connector_freeform_tags
  lifecycle {
    ignore_changes = [connection_info]
  }
}

# Add connector to an cloud Asm resource
resource "oci_database_management_cloud_asm" "test_cloud_asm" {
  cloud_asm_id = data.oci_database_management_cloud_asms.test_cloud_asms.cloud_asm_collection.0.items.0.id
  cloud_connector_id = oci_database_management_cloud_db_system_connector.test_cloud_asm_connector.id
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.asm_defined_tags_value
  }
  freeform_tags = var.asm_freeform_tags
}
*/

# Get CloudASM resource
data "oci_database_management_cloud_asm" "test_cloud_asm_get"  {
  #Required
  cloud_asm_id = data.oci_database_management_cloud_asms.test_cloud_asms.cloud_asm_collection.0.items.0.id
}

# Get configuration details for CloudAsm
data "oci_database_management_cloud_asm_configuration" "test_cloud_asm_configuration" {
  #Required
  cloud_asm_id = data.oci_database_management_cloud_asm.test_cloud_asm_get.id
}

# List ASM disk groups in CloudAsm
data "oci_database_management_cloud_asm_disk_groups" "test_cloud_asm_disk_groups" {
  #Required
  cloud_asm_id = data.oci_database_management_cloud_asm.test_cloud_asm_get.id
}

# List ASM users in CloudAsm
data "oci_database_management_cloud_asm_users" "test_cloud_asm_users" {
  #Required
  cloud_asm_id = data.oci_database_management_cloud_asm.test_cloud_asm_get.id
}

# List CloudAsmInstances in CloudAsm
data "oci_database_management_cloud_asm_instances" "test_cloud_asm_instances" {
  #Required
  cloud_asm_id = data.oci_database_management_cloud_asm.test_cloud_asm_get.id

  #Optional
  compartment_id  = var.compartment_id
}


# List CloudClusters in CloudDbSystem
data "oci_database_management_cloud_clusters" "test_cloud_clusters" {
  #Required
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List CloudClusterInstances in CloudCluster
data "oci_database_management_cloud_cluster_instances" "test_cloud_cluster_instances" {
  #Required
  cloud_cluster_id = data.oci_database_management_cloud_clusters.test_cloud_clusters.cloud_cluster_collection.0.items.0.id

  #Optional
  compartment_id      = var.compartment_id
}

/*
# Create cluster conenctor for rac1
resource "oci_database_management_cloud_db_system_connector" "test_cloud_cluster_connector_rac1" {
  #Required
  connector_type        = "MACS"
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id
  agent_id              = var.cloud_rac1_agent_id
  display_name = var.cloud_cluster_rac1_connector_display_name

  #Optional
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.db_system_connector_defined_tags_value
  }
  freeform_tags = var.db_system_connector_freeform_tags
  lifecycle {
    ignore_changes = [connection_info]
  }
}

# Create cluster conenctor for rac2
resource "oci_database_management_cloud_db_system_connector" "test_cloud_cluster_connector_rac2" {
  #Required
  connector_type        = "MACS"
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id
  agent_id              = var.cloud_rac2_agent_id
  display_name = var.cloud_cluster_rac2_connector_display_name

  #Optional
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.db_system_connector_defined_tags_value
  }
  freeform_tags = var.db_system_connector_freeform_tags
  lifecycle {
    ignore_changes = [connection_info]
  }
}


# Add connector to an cloud cluster instance resource, Change the number after items if required
resource "oci_database_management_cloud_cluster_instance" "test_cloud_cluster_instance1" {
  cloud_cluster_instance_id = data.oci_database_management_cloud_cluster_instances.test_cloud_cluster_instances.cloud_cluster_instance_collection.0.items.0.id
  cloud_connector_id = oci_database_management_cloud_db_system_connector.test_cloud_cluster_connector_rac1.id
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.asm_defined_tags_value
  }
  freeform_tags = var.asm_freeform_tags
}

# Add connector to an cloud cluster instance resource, Change the number after items if required
resource "oci_database_management_cloud_cluster_instance" "test_cloud_cluster_instance2" {
  cloud_cluster_instance_id = data.oci_database_management_cloud_cluster_instances.test_cloud_cluster_instances.cloud_cluster_instance_collection.0.items.1.id
  cloud_connector_id = oci_database_management_cloud_db_system_connector.test_cloud_cluster_connector_rac2.id
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.asm_defined_tags_value
  }
  freeform_tags = var.asm_freeform_tags
}
*/

# List CloudDatabases in CloudDbSystem
data "oci_database_management_cloud_databases" "test_cloud_databases" {
  #Required
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List CloudDbHomes in CloudDbSystem
data "oci_database_management_cloud_db_homes" "test_cloud_db_homes" {
  #Required
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List CloudDbNodes in CloudDbSystem
data "oci_database_management_cloud_db_nodes" "test_cloud_db_nodes" {
  #Required
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List CloudListeners in CloudDbSystem
data "oci_database_management_cloud_listeners" "test_cloud_listeners" {
  #Required
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id

  #Optional
  compartment_id        = var.compartment_id
  display_name =  var.local_listener_name
}

# Get CloudListener resource
data "oci_database_management_cloud_listener" "test_cloud_listener_get"  {
  #Required
  cloud_listener_id = data.oci_database_management_cloud_listeners.test_cloud_listeners.cloud_listener_collection.0.items.0.id
}

/*
# Creating listener connector
resource "oci_database_management_cloud_db_system_connector" "test_cloud_listener_connector" {
  #Required
  connector_type        = "MACS"
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id
  agent_id              = var.cloud_rac1_agent_id

  #Optional
  display_name = var.cloud_listener_connector_display_name
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.db_system_connector_defined_tags_value
  }
  freeform_tags = var.db_system_connector_freeform_tags
}

# Add connector to an ExternalListener resource
resource "oci_database_management_cloud_listener" "test_cloud_listener" {
  #Required
  cloud_listener_id = data.oci_database_management_cloud_listener.test_cloud_listener_get.id
  cloud_connector_id = oci_database_management_cloud_db_system_connector.test_cloud_listener_connector.id
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.listener_defined_tags_value
  }
  freeform_tags = var.listener_freeform_tags
}
*/

# List CloudListenerServices for CloudListener
data "oci_database_management_cloud_listener_services" "test_cloud_listener_services" {
  #Required
  cloud_listener_id = data.oci_database_management_cloud_listener.test_cloud_listener_get.id
  managed_database_id  = data.oci_database_management_cloud_databases.test_cloud_databases.cloud_database_collection.0.items.0.id
}


# List CloudDbSystemConnector resources in CloudDbSystem
data "oci_database_management_cloud_db_system_connectors" "test_cloud_db_system_connectors" {
  #Required
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id

  #Optional
  compartment_id        = var.compartment_id
  display_name          = var.cloud_asm_connector_display_name
}


#Enable database management for ExternalDbSystem
resource "oci_database_management_cloud_db_system_cloud_database_managements_management" "test_cloud_db_system_database_managements_management" {
  #Required
  cloud_db_system_id = data.oci_database_management_cloud_db_system.test_cloud_db_system.id
  enable_cloud_database_management = true
  is_enabled = true
}





