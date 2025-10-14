---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_db_system"
sidebar_current: "docs-oci-resource-database_management-cloud_db_system"
description: |-
  Provides the Cloud Db System resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_cloud_db_system
This resource provides the Cloud Db System resource in Oracle Cloud Infrastructure Database Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-management/latest/CloudDbSystem

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databasemanagement

Creates a cloud DB system and its related resources.


## Example Usage

```hcl
resource "oci_database_management_cloud_db_system" "test_cloud_db_system" {
	#Required
	compartment_id = var.compartment_id
	db_system_discovery_id = oci_database_management_db_system_discovery.test_db_system_discovery.id

	#Optional
	database_management_config {
		#Required
		is_enabled = var.cloud_db_system_database_management_config_is_enabled

		#Optional
		metadata = var.cloud_db_system_database_management_config_metadata
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.cloud_db_system_display_name
	freeform_tags = {"Department"= "Finance"}
	stack_monitoring_config {
		#Required
		is_enabled = var.cloud_db_system_stack_monitoring_config_is_enabled

		#Optional
		metadata = var.cloud_db_system_stack_monitoring_config_metadata
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the cloud DB system resides.
* `database_management_config` - (Optional) The configuration details of Database Management for a cloud DB system.
	* `is_enabled` - (Required) The status of the associated service.
	* `metadata` - (Optional) The associated service-specific inputs in JSON string format, which Database Management can identify.
* `db_system_discovery_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system discovery.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) The user-friendly name for the DB system. The name does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `stack_monitoring_config` - (Optional) The configuration details of Stack Monitoring for a cloud DB system.
	* `is_enabled` - (Required) The status of the associated service.
	* `metadata` - (Optional) The associated service-specific inputs in JSON string format, which Database Management can identify.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_management_config` - The configuration details of Database Management for a cloud DB system.
	* `is_enabled` - The status of the associated service.
	* `metadata` - The associated service-specific inputs in JSON string format, which Database Management can identify.
* `db_system_discovery_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system discovery.
* `dbaas_parent_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent cloud DB Infrastructure. For VM Dbsystems , it will be the DBSystem Id. For ExaCS and ExaCC,  it will be the cloudVmClusterId and vmClusterId respectively. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `deployment_type` - The deployment type of cloud dbsystem.
* `discovery_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used during the discovery of the DB system.
* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `home_directory` - The Oracle Grid home directory in case of cluster-based DB system and Oracle home directory in case of single instance-based DB system. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
* `is_cluster` - Indicates whether the DB system is a cluster DB system or not.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `stack_monitoring_config` - The configuration details of Stack Monitoring for a cloud DB system.
	* `is_enabled` - The status of the associated service.
	* `metadata` - The associated service-specific inputs in JSON string format, which Database Management can identify.
* `state` - The current lifecycle state of the cloud DB system resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cloud DB system was created.
* `time_updated` - The date and time the cloud DB system was last updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Db System
	* `update` - (Defaults to 20 minutes), when updating the Cloud Db System
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Db System


## Import

CloudDbSystems can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_cloud_db_system.test_cloud_db_system "id"
```

