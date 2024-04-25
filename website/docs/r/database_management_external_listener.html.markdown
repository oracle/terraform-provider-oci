---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_listener"
sidebar_current: "docs-oci-resource-database_management-external_listener"
description: |-
  Provides the External Listener resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_listener
This resource provides the External Listener resource in Oracle Cloud Infrastructure Database Management service.

Updates the external listener specified by `externalListenerId`.


## Example Usage

```hcl
resource "oci_database_management_external_listener" "test_external_listener" {
	#Required
	external_listener_id = oci_database_management_external_listener.test_external_listener.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	external_connector_id = oci_database_management_external_connector.test_external_connector.id
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `external_connector_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external connector.
* `external_listener_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external listener.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the external listener defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `adr_home_directory` - The directory that stores tracing and logging incidents when Automatic Diagnostic Repository (ADR) is enabled.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the external listener.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the external listener. The name does not have to be unique.
* `endpoints` - The list of protocol addresses the listener is configured to listen on.
	* `host` - The host name or IP address.
	* `key` - The unique name of the service.
	* `port` - The port number.
	* `protocol` - The listener protocol.
	* `services` - The list of services registered with the listener.
* `external_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external connector.
* `external_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB home.
* `external_db_node_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB node.
* `external_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system that the listener is a part of.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `host_name` - The name of the host on which the external listener is running.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external listener.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `listener_alias` - The listener alias.
* `listener_ora_location` - The location of the listener configuration file listener.ora.
* `listener_type` - The type of listener.
* `log_directory` - The destination directory of the listener log file.
* `oracle_home` - The Oracle home location of the listener.
* `serviced_asms` - The list of ASMs that are serviced by the listener.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the external ASM resides.
	* `display_name` - The user-friendly name for the external ASM. The name does not have to be unique.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM.
* `serviced_databases` - The list of databases that are serviced by the listener.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the external database resides.
	* `database_sub_type` - The subtype of Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database. 
	* `database_type` - The type of Oracle Database installation.
	* `db_unique_name` - The unique name of the external database.
	* `display_name` - The user-friendly name for the database. The name does not have to be unique.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database.
	* `is_managed` - Indicates whether the database is a Managed Database or not.
* `state` - The current lifecycle state of the external listener.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the external listener was created.
* `time_updated` - The date and time the external listener was last updated.
* `trace_directory` - The destination directory of the listener trace file.
* `version` - The listener version.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Listener
	* `update` - (Defaults to 20 minutes), when updating the External Listener
	* `delete` - (Defaults to 20 minutes), when destroying the External Listener


## Import

ExternalListeners can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_external_listener.test_external_listener "id"
```

