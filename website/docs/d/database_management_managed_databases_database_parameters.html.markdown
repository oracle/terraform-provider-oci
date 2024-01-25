---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_databases_database_parameters"
sidebar_current: "docs-oci-datasource-database_management-managed_databases_database_parameters"
description: |-
  Provides the list of Managed Databases Database Parameters in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_databases_database_parameters
This data source provides the list of Managed Databases Database Parameters in Oracle Cloud Infrastructure Database Management service.

Gets the list of database parameters for the specified Managed Database. The parameters are listed in alphabetical order, along with their current values.


## Example Usage

```hcl
data "oci_database_management_managed_databases_database_parameters" "test_managed_databases_database_parameters" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	is_allowed_values_included = var.managed_databases_database_parameter_is_allowed_values_included
	name = var.managed_databases_database_parameter_name
	opc_named_credential_id = var.managed_databases_database_parameter_opc_named_credential_id
	source = var.managed_databases_database_parameter_source
}
```

## Argument Reference

The following arguments are supported:

* `is_allowed_values_included` - (Optional) When true, results include a list of valid values for parameters (if applicable). 
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return all parameters that have the text given in their names. 
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.
* `source` - (Optional) The source used to list database parameters. `CURRENT` is used to get the database parameters that are currently in effect for the database instance. `SPFILE` is used to list parameters from the server parameter file. Default is `CURRENT`. 


## Attributes Reference

The following attributes are exported:

* `database_parameters_collection` - The list of database_parameters_collection.

### ManagedDatabasesDatabaseParameter Reference

The following attributes are exported:

* `database_name` - The name of the Managed Database.
* `database_sub_type` - The subtype of the Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, or a Non-container Database. 
* `database_type` - The type of Oracle Database installation.
* `database_version` - The Oracle Database version.
* `items` - An array of DatabaseParameterSummary objects.
	* `allowed_values` - A list of allowed values for this parameter.
		* `is_default` - Indicates whether the given ordinal value is the default value for the parameter.
		* `ordinal` - The ordinal number in the list (1-based).
		* `value` - The parameter value at ordinal.
	* `category` - The parameter category.
	* `constraint` - Applicable in case of Oracle Real Application Clusters (Oracle RAC) databases. A `UNIQUE` parameter is one which is unique to each Oracle Real Application Clusters (Oracle RAC) instance. For example, the parameter `INSTANCE_NUMBER` must have different values in each instance. An `IDENTICAL` parameter must have the same value for every instance. For example, the parameter `DB_BLOCK_SIZE` must have the same value in all instances. 
	* `container_id` - The ID of the database container to which the data pertains. Possible values include:
		* `0`: This value is used for data that pertain to the entire CDB. This value is also used for data in non-CDBs.
		* `1`: This value is used for data that pertain to only the root container.
		* `n`: Where n is the applicable container ID for the data. 
	* `description` - The description of the parameter.
	* `display_value` - The parameter value in a user-friendly format. For example, if the `value` property shows the value 262144 for a big integer parameter, then the `displayValue` property will show the value 256K. 
	* `is_adjusted` - Indicates whether Oracle adjusted the input value to a more suitable value.
	* `is_basic` - Indicates whether the parameter is a basic parameter (`TRUE`) or not (`FALSE`).
	* `is_default` - Indicates whether the parameter is set to the default value (`TRUE`) or the parameter value was specified in the parameter file (`FALSE`). 
	* `is_deprecated` - Indicates whether the parameter has been deprecated (`TRUE`) or not (`FALSE`).
	* `is_instance_modifiable` - For parameters that can be changed with `ALTER SYSTEM`, indicates whether the value of the parameter can be different for every instance (`TRUE`) or whether the parameter must have the same value for all Real Application Clusters instances (`FALSE`). For other parameters, this is always `FALSE`. 
	* `is_modified` - Indicates how the parameter was modified. If an `ALTER SYSTEM` was performed, the value will be `MODIFIED`. 
	* `is_pdb_modifiable` - Indicates whether the parameter can be modified on a per-PDB basis (`TRUE`) or not (`FALSE`). In a non-CDB, the value of this property is `null`. 
	* `is_session_modifiable` - Indicates whether the parameter can be changed with `ALTER SESSION` (`TRUE`) or not (`FALSE`) 
	* `is_specified` - Indicates whether the parameter was specified in the server parameter file (`TRUE`) or not (`FALSE`). Applicable only when the parameter source is `SPFILE`. 
	* `is_system_modifiable` - Indicates whether the parameter can be changed with `ALTER SYSTEM` and when the change takes effect:
		* IMMEDIATE: Parameter can be changed with `ALTER SYSTEM` regardless of the type of parameter file used to start the instance. The change takes effect immediately.
		* DEFERRED: Parameter can be changed with `ALTER SYSTEM` regardless of the type of parameter file used to start the instance. The change takes effect in subsequent sessions.
		* FALSE: Parameter cannot be changed with `ALTER SYSTEM` unless a server parameter file was used to start the instance. The change takes effect in subsequent instances. 
	* `name` - The parameter name.
	* `number` - The parameter number.
	* `ordinal` - The position (ordinal number) of the parameter value. Useful only for parameters whose values are lists of strings. 
	* `sid` - The database instance SID for which the parameter is defined.
	* `type` - The parameter type.
	* `update_comment` - The comments associated with the most recent update.
	* `value` - The parameter value.

