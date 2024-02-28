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

####################### External DB System #########################

variable "compartment_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-compartmentId-Value"
}

variable "external_db_system_discovery_display_name" {
  default = "EXAMPLE-displayName-Value"
}

variable "agent_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-agentId-Value"
}

variable "external_db_system_database_management_config_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "external_db_system_display_name" {
  default = "EXAMPLE-displayName-Value"
}

variable "db_host_name" {
  default = "EXAMPLE-hostName-Value"
}

variable "db_port" {
  default = "1521"
}

variable "db_service_name" {
  default = "EXAMPLE-service-Value"
}

variable "db_user_name" {
  default = "EXAMPLE-userName-Value"
}

variable "db_password_secret_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-passwordSecretId-Value"
}

variable "db_credential_name" {
  default = "EXAMPLE-dbCredName-Value"
}

variable "asm_host_name" {
  default = "EXAMPLE-hostName-Value"
}

variable "asm_port" {
  default = "1521"
}

variable "asm_service_name" {
  default = "EXAMPLE-service-Value"
}

variable "asm_user_name" {
  default = "EXAMPLE-userName-Value"
}

variable "asm_password_secret_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-passwordSecretId-Value"
}

variable "asm_credential_name" {
  default = "EXAMPLE-asmCredName-Value"
}

variable "external_asm_connector_display_name" {
  default = "asmConnectorName"
}

variable "external_listener_connector_display_name" {
  default = "listenerConnectorName"
}

variable "connector_agent_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-agentId-Value"
}

variable "local_listener_name" {
  default = "LISTENER_NAME"
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

# Create a new ExternalDbSystemDiscovery resource and discover an external DB System and its components.
# Also add a connector to the discovered Oracle Container Database (CDB).
resource "oci_database_management_external_db_system_discovery" "test_external_db_system_discovery" {
  #Required
  agent_id                        = var.agent_id
  compartment_id                  = var.compartment_id

  #Optional
  display_name = var.external_db_system_discovery_display_name
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.db_system_discovery_defined_tags_value
  }
  freeform_tags = var.db_system_discovery_freeform_tags

  # Patch the Discovery resource and add connector to the database component
  patch_operations {
    operation = "MERGE"
    selection = "discoveredComponents[?componentType == 'DATABASE'] | [0]"
    value {
      connector {
        agent_id = var.agent_id
        connection_info {
          component_type = "DATABASE"
          connection_credentials {
            credential_name = var.db_credential_name
            credential_type = "DETAILS"
            password_secret_id = var.db_password_secret_id
            role = "NORMAL"
            user_name = var.db_user_name
          }
          connection_string {
            host_name = var.db_host_name
            port = var.db_port
            protocol = "TCP"
            service = var.db_service_name
          }
        }
        connector_type = "MACS"
        display_name = "tersiDBconnector"
      }
      is_selected_for_monitoring = "true"
    }
  }

#  # Deselect the Pluggable Databases (PDBs).
#  patch_operations {
#    operation = "MERGE"
#    selection = "discoveredComponents[?componentType == 'DATABASE'].pluggableDatabases[]"
#    value {
#      is_selected_for_monitoring = "false"
#    }
#  }
}

# List ExternalDbSystemDiscovery resources
data "oci_database_management_external_db_system_discoveries" "test_external_db_system_discoveries" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.external_db_system_discovery_display_name
}

# Create a new ExternalDbSystem resource from ExternalDbSystemDiscovery resource
resource "oci_database_management_external_db_system" "test_external_db_system" {
  #Required
  compartment_id         = var.compartment_id
  db_system_discovery_id = oci_database_management_external_db_system_discovery.test_external_db_system_discovery.id

  #Optional
  database_management_config {
    #Required
    license_model = var.external_db_system_database_management_config_license_model
  }
  display_name = var.external_db_system_display_name
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.db_system_defined_tags_value
  }
  freeform_tags = var.db_system_freeform_tags
}

# List ExternalDbSystem resources
data "oci_database_management_external_db_systems" "test_external_db_systems" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.external_db_system_display_name
}

# Create a new Management Agent based ExternalDbSystemConnector
resource "oci_database_management_external_db_system_connector" "test_external_asm_connector" {
  #Required
  connector_type        = "MACS"
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id
  agent_id              = var.connector_agent_id

  #Optional
  display_name = var.external_asm_connector_display_name
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

# Add connector to an ExternalAsm resource
resource "oci_database_management_external_asm" "test_external_asm" {
  external_asm_id = data.oci_database_management_external_asms.test_external_asms.external_asm_collection.0.items.0.id
  external_connector_id = oci_database_management_external_db_system_connector.test_external_asm_connector.id
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.asm_defined_tags_value
  }
  freeform_tags = var.asm_freeform_tags
}

# List ExternalAsms in ExternalDbSystem
data "oci_database_management_external_asms" "test_external_asms" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# Get configuration details for ExternalAsm
data "oci_database_management_external_asm_configuration" "test_external_asm_configuration" {
  #Required
  external_asm_id = oci_database_management_external_asm.test_external_asm.id
}

# List ASM disk groups in ExternalAsm
data "oci_database_management_external_asm_disk_groups" "test_external_asm_disk_groups" {
  #Required
  external_asm_id = oci_database_management_external_asm.test_external_asm.id
}

# List ASM users in ExternalAsm
data "oci_database_management_external_asm_users" "test_external_asm_users" {
  #Required
  external_asm_id = oci_database_management_external_asm.test_external_asm.id
}

# List ExternalAsmInstances in ExternalAsm
data "oci_database_management_external_asm_instances" "test_external_asm_instances" {
  #Required
  external_asm_id = oci_database_management_external_asm.test_external_asm.id

  #Optional
  compartment_id  = var.compartment_id
}

# List ExternalClusters in ExternalDbSystem
data "oci_database_management_external_clusters" "test_external_clusters" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List ExternalClusterInstances in ExternalCluster
data "oci_database_management_external_cluster_instances" "test_external_cluster_instances" {
  #Required
  external_cluster_id = data.oci_database_management_external_clusters.test_external_clusters.external_cluster_collection.0.items.0.id

  #Optional
  compartment_id      = var.compartment_id
}

# List ExternalDatabases in ExternalDbSystem
data "oci_database_management_external_databases" "test_external_databases" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List ExternalDbHomes in ExternalDbSystem
data "oci_database_management_external_db_homes" "test_external_db_homes" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List ExternalDbNodes in ExternalDbSystem
data "oci_database_management_external_db_nodes" "test_external_db_nodes" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# Creating listener connector
resource "oci_database_management_external_db_system_connector" "test_external_listener_connector" {
  #Required
  connector_type        = "MACS"
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id
  agent_id              = var.connector_agent_id

  #Optional
  display_name = var.external_listener_connector_display_name
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.db_system_connector_defined_tags_value
  }
  freeform_tags = var.db_system_connector_freeform_tags
}

# List ExternalListeners in ExternalDbSystem
data "oci_database_management_external_listeners" "test_external_listeners" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
  # display_name =  var.local_listener_name
}

# Add connector to an ExternalListener resource
resource "oci_database_management_external_listener" "test_external_listener" {
  #Required
  external_listener_id = data.oci_database_management_external_listeners.test_external_listeners.external_listener_collection.0.items.0.id
  external_connector_id = oci_database_management_external_db_system_connector.test_external_listener_connector.id
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.listener_defined_tags_value
  }
  freeform_tags = var.listener_freeform_tags
}

# List ExternalListenerServices for ExternalListener
data "oci_database_management_external_listener_services" "test_external_listener_services" {
  #Required
  external_listener_id = oci_database_management_external_listener.test_external_listener.id
  #managed_database_id  = oci_database_management_external_listener.test_external_listener.serviced_databases.0.id
  managed_database_id  = data.oci_database_management_external_databases.test_external_databases.external_database_collection.0.items.0.id
}

# List ExternalDbSystemConnector resources in ExternalDbSystem
data "oci_database_management_external_db_system_connectors" "test_external_db_system_connectors" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
  display_name          = var.external_asm_connector_display_name
  depends_on = [oci_database_management_external_db_system_connector.test_external_asm_connector]
}

data "oci_database_management_managed_databases_asm_properties" "test_managed_databases_asm_properties" {
  #Required
  managed_database_id = data.oci_database_management_external_databases.test_external_databases.external_database_collection.0.items.0.id

  #Optional
  name = var.managed_databases_asm_property_name
}

data "oci_database_management_managed_databases_asm_property" "test_managed_databases_asm_property" {
  #Required
  managed_database_id = data.oci_database_management_external_databases.test_external_databases.external_database_collection.0.items.0.id

  #Optional
  name = var.managed_databases_asm_property_name
}

## Disable database management for ExternalDbSystem
#resource "oci_database_management_external_db_system_database_managements_management" "test_external_db_system_database_managements_management" {
#  #Required
#  external_db_system_id      = oci_database_management_external_db_system.test_external_db_system.id
#  enable_database_management = false
#}

# Enable Stack Monitoring for ExternalDbSystem
resource "oci_database_management_external_db_system_stack_monitorings_management" "test_external_db_system_stack_monitoring_management" {
  #Required
  external_db_system_id      = oci_database_management_external_db_system.test_external_db_system.id
  enable_stack_monitoring    = true
  is_enabled                 = true
}

## Disable Stack Monitoring for ExternalDbSystem
#resource "oci_database_management_external_db_system_stack_monitorings_management" "test_external_db_system_disable_stack_monitoring_management" {
#  #Required
#  external_db_system_id      = oci_database_management_external_db_system.test_external_db_system.id
#  enable_stack_monitoring    = false
#}
