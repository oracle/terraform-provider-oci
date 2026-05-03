---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_connection_credential_public_synonym"
sidebar_current: "docs-oci-resource-database_tools_runtime-database_tools_connection_credential_public_synonym"
description: |-
  Provides the Database Tools Connection Credential Public Synonym resource in Oracle Cloud Infrastructure Database Tools Runtime service
---

# oci_database_tools_runtime_database_tools_connection_credential_public_synonym
This resource provides the Database Tools Connection Credential Public Synonym resource in Oracle Cloud Infrastructure Database Tools Runtime service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database_tools_runtime

Creates a public synonym for the given credentials

## Example Usage

```hcl
resource "oci_database_tools_runtime_database_tools_connection_credential_public_synonym" "test_database_tools_connection_credential_public_synonym" {
	#Required
	credential_key = var.database_tools_connection_credential_public_synonym_credential_key
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
	key = var.database_tools_connection_credential_public_synonym_key
}
```

## Argument Reference

The following arguments are supported:

* `credential_key` - (Required) The name of the credential
* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.
* `key` - (Required) The name of the public synonym for the credential


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `key` - The name of the public synonym for the credential

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Connection Credential Public Synonym
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Connection Credential Public Synonym
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Connection Credential Public Synonym


## Import

DatabaseToolsConnectionCredentialPublicSynonyms can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_runtime_database_tools_connection_credential_public_synonym.test_database_tools_connection_credential_public_synonym "databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}/publicSynonyms/{publicSynonymKey}" 
```

