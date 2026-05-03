---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_connection_credential_public_synonym"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_connection_credential_public_synonym"
description: |-
  Provides details about a specific Database Tools Connection Credential Public Synonym in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_connection_credential_public_synonym
This data source provides details about a specific Database Tools Connection Credential Public Synonym resource in Oracle Cloud Infrastructure Database Tools Runtime service.

Get a public synonym

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_connection_credential_public_synonym" "test_database_tools_connection_credential_public_synonym" {
	#Required
	credential_key = var.database_tools_connection_credential_public_synonym_credential_key
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
	public_synonym_key = var.database_tools_connection_credential_public_synonym_public_synonym_key
}
```

## Argument Reference

The following arguments are supported:

* `credential_key` - (Required) The name of the credential
* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.
* `public_synonym_key` - (Required) The name of the public synonym for the credential


## Attributes Reference

The following attributes are exported:

* `key` - The name of the public synonym for the credential

