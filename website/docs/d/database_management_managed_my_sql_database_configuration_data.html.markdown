---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database_configuration_data"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database_configuration_data"
description: |-
  Provides the list of Managed My Sql Database Configuration Data in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database_configuration_data
This data source provides the list of Managed My Sql Database Configuration Data in Oracle Cloud Infrastructure Database Management service.

Retrieves Configuration Data for given MySQL Instance.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database_configuration_data" "test_managed_my_sql_database_configuration_data" {
	#Required
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_my_sql_database_id` - (Required) The OCID of ManagedMySqlDatabase.


## Attributes Reference

The following attributes are exported:

* `my_sql_configuration_data_collection` - The list of my_sql_configuration_data_collection.

### ManagedMySqlDatabaseConfigurationData Reference

The following attributes are exported:

* `items` - List of ConfigurationDataSummary.
	* `default_value` - default value of variable
	* `description` - Description of the variable
	* `host_set` - Host from where this value was set. Empty for MySql Database System
	* `is_configurable` - Whether this variable is configurable
	* `is_dynamic` - Whether variable can be set dynamically or not
	* `is_init` - whether variable is set at server startup
	* `max_value` - Maximum value of variable
	* `min_value` - Minimum value of variable
	* `name` - The name of variable
	* `path` - If the variable was set from an option file, VARIABLE_PATH is the path name of that file. Otherwise, the value is the empty string.
	* `possible_values` - Comma separated list of possible values for the variable in value:valueDescription format
	* `source` - The source from which the variable was most recently set
	* `supported_versions` - Comma separated list of MySql versions where this variable is supported
	* `time_set` - Time when value was set
	* `type` - type of variable
	* `user_set` - User who set this value. Empty for MySql Database System
	* `value` - The value of variable

