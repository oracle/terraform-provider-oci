---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_listener"
sidebar_current: "docs-oci-resource-database_management-cloud_listener"
description: |-
  Provides the Cloud Listener resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_cloud_listener
This resource provides the Cloud Listener resource in Oracle Cloud Infrastructure Database Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-management/latest/CloudListener

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databasemanagement

Updates the cloud listener specified by `cloudListenerId`.


## Example Usage

```hcl
resource "oci_database_management_cloud_listener" "test_cloud_listener" {
	#Required
	cloud_listener_id = oci_database_management_cloud_listener.test_cloud_listener.id

	#Optional
	cloud_connector_id = oci_database_management_cloud_connector.test_cloud_connector.id
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `cloud_connector_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
* `cloud_listener_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud listener.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the cloud listener defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `adr_home_directory` - The directory that stores tracing and logging incidents when Automatic Diagnostic Repository (ADR) is enabled.
* `cloud_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
* `cloud_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB home.
* `cloud_db_node_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB node.
* `cloud_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the listener is a part of.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the cloud listener.
* `dbaas_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in DBaas service.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the cloud listener. The name does not have to be unique.
* `endpoints` - The list of protocol addresses the listener is configured to listen on.
	* `host` - The host name or IP address.
	* `key` - The unique name of the service.
	* `port` - The port number.
	* `protocol` - The listener protocol.
	* `services` - The list of services registered with the listener.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `host_name` - The name of the host on which the cloud listener is running.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud listener.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `listener_alias` - The listener alias.
* `listener_ora_location` - The location of the listener configuration file listener.ora.
* `listener_type` - The type of listener.
* `log_directory` - The destination directory of the listener log file.
* `oracle_home` - The Oracle home location of the listener.
* `serviced_asms` - The list of ASMs that are serviced by the listener.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the cloud ASM resides.
	* `display_name` - The user-friendly name for the cloud ASM. The name does not have to be unique.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM.
* `serviced_databases` - The list of databases that are serviced by the listener.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the cloud database resides.
	* `database_sub_type` - The subtype of Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database. 
	* `database_type` - The type of Oracle Database installation.
	* `db_unique_name` - The unique name of the cloud database.
	* `dbaas_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Dbaas Oracle Cloud Infrastructure resource matching the discovered DB system component.
	* `display_name` - The user-friendly name for the database. The name does not have to be unique.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud database.
	* `is_managed` - Indicates whether the database is a Managed Database or not.
* `state` - The current lifecycle state of the cloud listener.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cloud listener was created.
* `time_updated` - The date and time the cloud listener was last updated.
* `trace_directory` - The destination directory of the listener trace file.
* `version` - The listener version.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Listener
	* `update` - (Defaults to 20 minutes), when updating the Cloud Listener
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Listener


## Import

CloudListeners can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_cloud_listener.test_cloud_listener "id"
```

