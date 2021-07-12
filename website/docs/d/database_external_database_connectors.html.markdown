---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_external_database_connectors"
sidebar_current: "docs-oci-datasource-database-external_database_connectors"
description: |-
  Provides the list of External Database Connectors in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_external_database_connectors
This data source provides the list of External Database Connectors in Oracle Cloud Infrastructure Database service.

Gets a list of the external database connectors in the specified compartment.


## Example Usage

```hcl
data "oci_database_external_database_connectors" "test_external_database_connectors" {
	#Required
	compartment_id = var.compartment_id
	external_database_id = oci_database_database.test_database.id

	#Optional
	display_name = var.external_database_connector_display_name
	state = var.external_database_connector_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `external_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database whose connectors will be listed. 
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state.


## Attributes Reference

The following attributes are exported:

* `external_database_connectors` - The list of external_database_connectors.

### ExternalDatabaseConnector Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_credentials` - Credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector crendentials. 
	* `credential_name` - The name of the credential information that used to connect to the database. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

		For example: inventorydb.abc112233445566778899 
	* `credential_type` - The type of credential used to connect to the database.
	* `password` - The password that will be used to connect to the database.
	* `role` - The role of the user that will be connecting to the database.
	* `username` - The username that will be used to connect to the database.
* `connection_status` - The status of connectivity to the external database.
* `connection_string` - The Oracle Database connection string. 
	* `hostname` - The host name of the database.
	* `port` - The port used to connect to the database.
	* `protocol` - The protocol used to connect to the database.
	* `service` - The name of the service alias used to connect to the database.
* `connector_agent_id` - The ID of the agent used for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
* `connector_type` - The type of connector used by the external database resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). The name does not have to be unique. 
* `external_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the external database connector resource.
* `time_connection_status_last_updated` - The date and time the `connectionStatus` of this external connector was last updated.
* `time_created` - The date and time the external connector was created.

