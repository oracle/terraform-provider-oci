---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_database_security_configs"
sidebar_current: "docs-oci-datasource-data_safe-database_security_configs"
description: |-
  Provides the list of Database Security Configs in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_database_security_configs
This data source provides the list of Database Security Configs in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all database security configurations in Data Safe.

The ListDatabaseSecurityConfigs operation returns only the database security configurations in the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListDatabaseSecurityConfigs on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_database_security_configs" "test_database_security_configs" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.database_security_config_access_level
	compartment_id_in_subtree = var.database_security_config_compartment_id_in_subtree
	database_security_config_id = oci_data_safe_database_security_config.test_database_security_config.id
	display_name = var.database_security_config_display_name
	state = var.database_security_config_state
	target_id = oci_cloud_guard_target.test_target.id
	time_created_greater_than_or_equal_to = var.database_security_config_time_created_greater_than_or_equal_to
	time_created_less_than = var.database_security_config_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `database_security_config_id` - (Optional) An optional filter to return only resources that match the specified OCID of the database security configuration resource.
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `state` - (Optional) The current state of the database security configuration.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `database_security_config_collection` - The list of database_security_config_collection.

### DatabaseSecurityConfig Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the database security config.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the database security config.
* `display_name` - The display name of the database security config.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the database security config.
* `lifecycle_details` - Details about the current state of the database security config in Data Safe.
* `sql_firewall_config` - The SQL Firewall related configurations. 
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

