---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_database_security_config"
sidebar_current: "docs-oci-datasource-data_safe-database_security_config"
description: |-
  Provides details about a specific Database Security Config in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_database_security_config
This data source provides details about a specific Database Security Config resource in Oracle Cloud Infrastructure Data Safe service.

Gets a database security configuration by identifier.

## Example Usage

```hcl
data "oci_data_safe_database_security_config" "test_database_security_config" {
	#Required
	database_security_config_id = oci_data_safe_database_security_config.test_database_security_config.id
}
```

## Argument Reference

The following arguments are supported:

* `database_security_config_id` - (Required) The OCID of the database security configuration resource.


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

