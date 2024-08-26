---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_exadata_insight"
sidebar_current: "docs-oci-resource-opsi-exadata_insight"
description: |-
 Provides the Exadata Insight resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_exadata_insight
This resource provides the Exadata Insight resource in Oracle Cloud Infrastructure Opsi service.

Create an Exadata insight resource for an Exadata system in Operations Insights. The Exadata system will be enabled in Operations Insights. Exadata-related metric collection and analysis will be started.


## Example Usage

```hcl
resource "oci_opsi_exadata_insight" "test_exadata_insight" {
  #Required
  compartment_id = var.compartment_id
  enterprise_manager_bridge_id = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
  enterprise_manager_entity_identifier = var.exadata_insight_enterprise_manager_entity_identifier
  enterprise_manager_identifier = var.exadata_insight_enterprise_manager_identifier
  entity_source = var.exadata_insight_entity_source

  #Optional
  defined_tags = {"foo-namespace.bar-key"= "value"}
  freeform_tags = {"bar-key"= "value"}
  is_auto_sync_enabled = var.exadata_insight_is_auto_sync_enabled
  member_vm_cluster_details {

    #Optional
    compartment_id = var.compartment_id
    dbm_private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
    member_database_details {

      #Optional
      compartment_id = var.compartment_id
      connection_credential_details {
        #Required
        credential_type = var.exadata_insight_member_vm_cluster_details_member_database_details_connection_credential_details_credential_type

        #Optional
        credential_source_name = var.exadata_insight_member_vm_cluster_details_member_database_details_connection_credential_details_credential_source_name
        password_secret_id = oci_vault_secret.test_secret.id
        role = var.exadata_insight_member_vm_cluster_details_member_database_details_connection_credential_details_role
        user_name = oci_identity_user.test_user.name
        wallet_secret_id = oci_vault_secret.test_secret.id
      }
      connection_details {

        #Optional
        host_name = var.exadata_insight_member_vm_cluster_details_member_database_details_connection_details_host_name
        hosts {

          #Optional
          host_ip = var.exadata_insight_member_vm_cluster_details_member_database_details_connection_details_hosts_host_ip
          port = var.exadata_insight_member_vm_cluster_details_member_database_details_connection_details_hosts_port
        }
        port = var.exadata_insight_member_vm_cluster_details_member_database_details_connection_details_port
        protocol = var.exadata_insight_member_vm_cluster_details_member_database_details_connection_details_protocol
        service_name = oci_core_service.test_service.name
      }
      credential_details {
        #Required
        credential_type = var.exadata_insight_member_vm_cluster_details_member_database_details_credential_details_credential_type

        #Optional
        credential_source_name = var.exadata_insight_member_vm_cluster_details_member_database_details_credential_details_credential_source_name
        password_secret_id = oci_vault_secret.test_secret.id
        role = var.exadata_insight_member_vm_cluster_details_member_database_details_credential_details_role
        user_name = oci_identity_user.test_user.name
        wallet_secret_id = oci_vault_secret.test_secret.id
      }
      database_id = oci_database_database.test_database.id
      database_resource_type = var.exadata_insight_member_vm_cluster_details_member_database_details_database_resource_type
      dbm_private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
      defined_tags = var.exadata_insight_member_vm_cluster_details_member_database_details_defined_tags
      deployment_type = var.exadata_insight_member_vm_cluster_details_member_database_details_deployment_type
      entity_source = var.exadata_insight_member_vm_cluster_details_member_database_details_entity_source
      freeform_tags = var.exadata_insight_member_vm_cluster_details_member_database_details_freeform_tags
      management_agent_id = oci_management_agent_management_agent.test_management_agent.id
      opsi_private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
      service_name = oci_core_service.test_service.name
      system_tags = var.exadata_insight_member_vm_cluster_details_member_database_details_system_tags
    }
    opsi_private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
    vm_cluster_type = var.exadata_insight_member_vm_cluster_details_vm_cluster_type
    vmcluster_id = oci_opsi_vmcluster.test_vmcluster.id
  }
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier of Exadata insight
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `enterprise_manager_bridge_id` - (Required) OPSI Enterprise Manager Bridge OCID
* `enterprise_manager_entity_identifier` - (Required) Enterprise Manager Entity Unique Identifier
* `enterprise_manager_identifier` - (Required) Enterprise Manager Unique Identifier
* `entity_source` - (Required) (Updatable) Source of the Exadata system.
* `exadata_infra_id` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Infrastructure.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `is_auto_sync_enabled` - (Applicable when entity_source=EM_MANAGED_EXTERNAL_EXADATA) (Updatable) Set to true to enable automatic enablement and disablement of related targets from Enterprise Manager. New resources (e.g. Database Insights) will be placed in the same compartment as the related Exadata Insight.
* `member_vm_cluster_details` - (Applicable when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA)
    * `compartment_id` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
    * `dbm_private_endpoint_id` - (Applicable when entity_source=PE_COMANAGED_EXADATA) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint
    * `member_database_details` - (Applicable when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) The databases that belong to the VM Cluster
        * `compartment_id` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) (Updatable) Compartment Identifier of database
        * `connection_credential_details` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA) User credential details to connect to the database.
            * `credential_source_name` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA) Credential source name that had been added in Management Agent wallet. This is supplied in the External Database Service.
            * `credential_type` - (Required) Credential type.
            * `password_secret_id` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) The secret [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) mapping to the database credentials.
            * `role` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) database user role.
            * `user_name` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) database user name.
            * `wallet_secret_id` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the database keystore contents are stored. This is used for TCPS support in BM/VM/ExaCS cases.
        * `connection_details` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) Connection details to connect to the database. HostName, protocol, and port should be specified.
            * `host_name` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA) Name of the listener host that will be used to create the connect string to the database.
            * `hosts` - (Required when entity_source=PE_COMANAGED_EXADATA) List of hosts and port for private endpoint accessed database resource.
                * `host_ip` - (Applicable when entity_source=PE_COMANAGED_EXADATA) Host IP used for connection requests for Cloud DB resource.
                * `port` - (Applicable when entity_source=PE_COMANAGED_EXADATA) Listener port number used for connection requests for rivate endpoint accessed db resource.
            * `port` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA) Listener port number used for connection requests.
            * `protocol` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) Protocol used for connection requests for private endpoint accssed database resource.
            * `service_name` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) Database service name used for connection requests.
        * `credential_details` - (Required when entity_source=PE_COMANAGED_EXADATA) User credential details to connect to the database.
            * `credential_source_name` - (Required when entity_source=PE_COMANAGED_EXADATA) Credential source name that had been added in Management Agent wallet. This is supplied in the External Database Service.
            * `credential_type` - (Required) Credential type.
            * `password_secret_id` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) The secret [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) mapping to the database credentials.
            * `role` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) database user role.
            * `user_name` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) database user name.
            * `wallet_secret_id` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the database keystore contents are stored. This is used for TCPS support in BM/VM/ExaCS cases.
        * `database_id` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
        * `database_resource_type` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) Oracle Cloud Infrastructure database resource type
        * `dbm_private_endpoint_id` - (Applicable when entity_source=PE_COMANAGED_EXADATA) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint
        * `defined_tags` - (Applicable when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
        * `deployment_type` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) Database Deployment Type (EXACS will be supported in the future)
        * `entity_source` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) Source of the database entity.
        * `freeform_tags` - (Applicable when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        * `management_agent_id` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
        * `opsi_private_endpoint_id` - (Applicable when entity_source=PE_COMANAGED_EXADATA) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
        * `service_name` - (Required when entity_source=PE_COMANAGED_EXADATA) Database service name used for connection requests.
        * `system_tags` - (Applicable when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}`
    * `opsi_private_endpoint_id` - (Applicable when entity_source=PE_COMANAGED_EXADATA) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
    * `vm_cluster_type` - (Applicable when entity_source=MACS_MANAGED_CLOUD_EXADATA) Exadata VMCluster type
    * `vmcluster_id` - (Required when entity_source=MACS_MANAGED_CLOUD_EXADATA | PE_COMANAGED_EXADATA) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster.
* `status` - (Optional) (Updatable) Status of the resource. Example: "ENABLED", "DISABLED". Resource can be either enabled or disabled by updating the value of status field to either "ENABLED" or "DISABLED"


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier of the Exadata insight resource
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `enterprise_manager_bridge_id` - OPSI Enterprise Manager Bridge OCID
* `enterprise_manager_entity_display_name` - Enterprise Manager Entity Display Name
* `enterprise_manager_entity_identifier` - Enterprise Manager Entity Unique Identifier
* `enterprise_manager_entity_name` - Enterprise Manager Entity Name
* `enterprise_manager_entity_type` - Enterprise Manager Entity Type
* `enterprise_manager_identifier` - Enterprise Manager Unique Identifier
* `entity_source` - Source of the Exadata system.
* `exadata_display_name` - The user-friendly name for the Exadata system. The name does not have to be unique.
* `exadata_name` - The Exadata system name. If the Exadata systems managed by Enterprise Manager, the name is unique amongst the Exadata systems managed by the same Enterprise Manager.
* `exadata_rack_type` - Exadata rack type.
* `exadata_type` - Operations Insights internal representation of the the Exadata system type.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `id` - Exadata insight identifier
* `is_auto_sync_enabled` - Set to true to enable automatic enablement and disablement of related targets from Enterprise Manager. New resources (e.g. Database Insights) will be placed in the same compartment as the related Exadata Insight.
* `is_virtualized_exadata` - true if virtualization is used in the Exadata system
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the Exadata insight.
* `status` - Indicates the status of an Exadata insight in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the the Exadata insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the Exadata insight was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Exadata Insight
* `update` - (Defaults to 20 minutes), when updating the Exadata Insight
* `delete` - (Defaults to 20 minutes), when destroying the Exadata Insight


## Import

ExadataInsights can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_exadata_insight.test_exadata_insight "id"
```
