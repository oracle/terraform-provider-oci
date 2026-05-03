---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_connection_credential_public_synonyms"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_connection_credential_public_synonyms"
description: |-
  Provides the list of Database Tools Connection Credential Public Synonyms in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_connection_credential_public_synonyms
This data source provides the list of Database Tools Connection Credential Public Synonyms in Oracle Cloud Infrastructure Database Tools Runtime service.

Get a list of all public synonyms for the given credential

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_connection_credential_public_synonyms" "test_database_tools_connection_credential_public_synonyms" {
	#Required
	credential_key = var.database_tools_connection_credential_public_synonym_credential_key
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
}
```

## Argument Reference

The following arguments are supported:

* `credential_key` - (Required) The name of the credential
* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.


## Attributes Reference

The following attributes are exported:

* `credential_public_synonym_collection` - The list of credential_public_synonym_collection.

### DatabaseToolsConnectionCredentialPublicSynonym Reference

The following attributes are exported:

* `key` - The name of the public synonym for the credential

