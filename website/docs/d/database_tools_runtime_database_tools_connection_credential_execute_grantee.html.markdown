---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_connection_credential_execute_grantee"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_connection_credential_execute_grantee"
description: |-
  Provides details about a specific Database Tools Connection Credential Execute Grantee in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_connection_credential_execute_grantee
This data source provides details about a specific Database Tools Connection Credential Execute Grantee resource in Oracle Cloud Infrastructure Database Tools Runtime service.

Get a credential execute grantee

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_connection_credential_execute_grantee" "test_database_tools_connection_credential_execute_grantee" {
	#Required
	credential_key = var.database_tools_connection_credential_execute_grantee_credential_key
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
	execute_grantee_key = var.database_tools_connection_credential_execute_grantee_execute_grantee_key
}
```

## Argument Reference

The following arguments are supported:

* `credential_key` - (Required) The name of the credential
* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.
* `execute_grantee_key` - (Required) The name of the user granted the EXECUTE privilege on the credential.


## Attributes Reference

The following attributes are exported:

* `key` - The name of the user to grant the EXECUTE privilege on the credential.

