---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_connection_credential"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_connection_credential"
description: |-
  Provides details about a specific Database Tools Connection Credential in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_connection_credential
This data source provides details about a specific Database Tools Connection Credential resource in Oracle Cloud Infrastructure Database Tools Runtime service.

Get a credential

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_connection_credential" "test_database_tools_connection_credential" {
	#Required
	credential_key = var.database_tools_connection_credential_credential_key
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
}
```

## Argument Reference

The following arguments are supported:

* `credential_key` - (Required) The name of the credential
* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.


## Attributes Reference

The following attributes are exported:

* `enabled` - Indicates whether this credential is enabled (TRUE) or not (FALSE)
* `key` - Name of the credential
* `key_type` - Indicates whether this refers to a public synonym or not.
* `owner` - Owner of the credential
* `related_resource` - A related resource for a credential.
	* `identifier` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related resource.
	* `type` - The related resource type.
* `user_name` - Name of the user that will be used to log in to the remote database or the remote or local operating system
* `windows_domain` - For a Windows target, the Windows domain to use when logging in

