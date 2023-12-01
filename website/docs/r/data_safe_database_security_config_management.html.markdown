---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_database_security_config_management"
sidebar_current: "docs-oci-resource-data_safe-database_security_config_management"
description: |-
  Provides the Database Security Config Management resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_database_security_config_management
This resource provides the Database Security Config Management resource in Oracle Cloud Infrastructure Data Safe service.

Updates the database security configuration.

## Example Usage

```hcl
resource "oci_data_safe_database_security_config_management" "test_database_security_config_management" {
	#Required
	compartment_id = var.compartment_id
	target_id = oci_data_safe_target_database.test_target_database.id
	
	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.database_security_config_management_description
	display_name = var.database_security_config_management_display_name
	freeform_tags = {"Department"= "Finance"}
	sql_firewall_config {

		#Optional
		exclude_job = var.database_security_config_management_sql_firewall_config_exclude_job
		status = var.database_security_config_management_sql_firewall_config_status
		violation_log_auto_purge = var.database_security_config_management_sql_firewall_config_violation_log_auto_purge
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment containing the database security config.
* `target_id` - (Required) Unique target identifier.
* `database_security_config_id` - (Required) The OCID of the database security configuration resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the security policy.
* `display_name` - (Optional) (Updatable) The display name of the database security config. The name does not have to be unique, and it is changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `sql_firewall_config` - (Optional) (Updatable) Details to update the SQL firewall config. 
	* `exclude_job` - (Optional) (Updatable) Specifies whether the firewall should include or exclude the database internal job activities.
	* `status` - (Optional) (Updatable) Specifies whether the firewall is enabled or disabled on the target database.
	* `violation_log_auto_purge` - (Optional) (Updatable) Specifies whether Data Safe should automatically purge the violation logs  from the database after collecting the violation logs and persisting on Data Safe. 
* `refresh_trigger` - (Optional) (Updatable) An optional property when incremented triggers Refresh. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the database security config.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the database security config.
* `display_name` - The display name of the database security config.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the database security config.
* `lifecycle_details` - Details about the current state of the database security config in Data Safe.
* `sql_firewall_config` - The SQL firewall related configurations. 
	* `exclude_job` - Specifies whether the firewall should include or exclude the database internal job activities.
	* `status` - Specifies if the firewall is enabled or disabled on the target database.
	* `time_status_updated` - The most recent time when the firewall status is updated, in the format defined by RFC3339.
	* `violation_log_auto_purge` - Specifies whether Data Safe should automatically purge the violation logs  from the database after collecting the violation logs and persisting on Data Safe. 
* `state` - The current state of the database security config.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The target OCID corresponding to the database security config.
* `time_created` - The time that the database security config was created, in the format defined by RFC3339.
* `time_last_refreshed` - The last date and time the database security config was refreshed, in the format defined by RFC3339.
* `time_updated` - The date and time the database security configuration was last updated, in the format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Security Config Management
	* `update` - (Defaults to 20 minutes), when updating the Database Security Config Management
	* `delete` - (Defaults to 20 minutes), when destroying the Database Security Config Management


## Import

Import is not supported for this resource.

