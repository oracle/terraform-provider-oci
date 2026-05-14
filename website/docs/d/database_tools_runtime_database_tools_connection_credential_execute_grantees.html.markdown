---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_connection_credential_execute_grantees"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_connection_credential_execute_grantees"
description: |-
  Provides the list of Database Tools Connection Credential Execute Grantees in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_connection_credential_execute_grantees
This data source provides the list of Database Tools Connection Credential Execute Grantees in Oracle Cloud Infrastructure Database Tools Runtime service.

Get a list of all execute grantees

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_connection_credential_execute_grantees" "test_database_tools_connection_credential_execute_grantees" {
	#Required
	credential_key = var.database_tools_connection_credential_execute_grantee_credential_key
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
}
```

## Argument Reference

The following arguments are supported:

* `credential_key` - (Required) The name of the credential
* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.


## Attributes Reference

The following attributes are exported:

* `credential_execute_grantee_collection` - The list of credential_execute_grantee_collection.

### DatabaseToolsConnectionCredentialExecuteGrantee Reference

The following attributes are exported:

* `key` - The name of the user to grant the EXECUTE privilege on the credential.

