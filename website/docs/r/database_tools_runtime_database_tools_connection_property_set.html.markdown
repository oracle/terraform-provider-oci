---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_connection_property_set"
sidebar_current: "docs-oci-resource-database_tools_runtime-database_tools_connection_property_set"
description: |-
  Provides the Database Tools Connection Property Set resource in Oracle Cloud Infrastructure Database Tools Runtime service
---

# oci_database_tools_runtime_database_tools_connection_property_set
This resource provides the Database Tools Connection Property Set resource in Oracle Cloud Infrastructure Database Tools Runtime service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database_tools_runtime

Update a property set

## Example Usage

```hcl
resource "oci_database_tools_runtime_database_tools_connection_property_set" "test_database_tools_connection_property_set" {
	#Required
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
	key = var.database_tools_connection_property_set_key
	property_set_key = var.database_tools_connection_property_set_property_set_key

	#Optional
	authentication_substitutions = var.database_tools_connection_property_set_authentication_substitutions
	autonomous_database_resource_principal_status = var.database_tools_connection_property_set_autonomous_database_resource_principal_status
	credential_key = var.database_tools_connection_property_set_credential_key
	function_id = oci_functions_function.test_function.id
	identity_provider {
		#Required
		type = var.database_tools_connection_property_set_identity_provider_type

		#Optional
		configs = var.database_tools_connection_property_set_identity_provider_configs
	}
	instance_dbms_credential_enabled = var.database_tools_connection_property_set_instance_dbms_credential_enabled
	invoke_endpoint = var.database_tools_connection_property_set_invoke_endpoint
	object_storage_bucket_compartment_id = oci_identity_compartment.test_compartment.id
	object_storage_endpoint = var.database_tools_connection_property_set_object_storage_endpoint
	object_storage_namespace = var.database_tools_connection_property_set_object_storage_namespace
	print_server_type = var.database_tools_connection_property_set_print_server_type
}
```

## Argument Reference

The following arguments are supported:

* `authentication_substitutions` - (Applicable when key=APEX_FA_INTEGRATION) (Updatable) APEX FA Integration key-value pairs.
* `autonomous_database_resource_principal_status` - (Applicable when key=APEX_DOCUMENT_GENERATOR) (Updatable) The status of the Autonomous Database Serverless Resource Principal (OCI$RESOURCE_PRINCIPAL)
* `credential_key` - (Applicable when key=APEX_DOCUMENT_GENERATOR) (Updatable) The name of the credential used by APEX to manage Object Storage Buckets and Objects as well as invoke the Document Generator function.
* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.
* `function_id` - (Applicable when key=APEX_DOCUMENT_GENERATOR) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Document Generator function
* `identity_provider` - (Required when key=ORACLE_DATABASE_EXTERNAL_AUTHENTICATION) (Updatable) External identity type provider
	* `configs` - (Required when type=AZURE_AD) (Updatable) External identity provider configuration parameters. Simple key-value pair Example: { "tenant_id": "...", "application_id_uri": "...", ... } 
	* `type` - (Required) (Updatable) External identity type provider.  Supported values - OCI_IAM, AZURE_AD, NONE.
* `instance_dbms_credential_enabled` - (Applicable when key=APEX_FA_INTEGRATION) (Updatable) Specifies whether database credentials can be used in all workspaces on the APEX instance. Supported values include: "Y", "N" and empty string.
* `invoke_endpoint` - (Applicable when key=APEX_DOCUMENT_GENERATOR) (Updatable) The base endpoint URL to use to invoke the Document Generator function
* `key` - (Required) (Updatable) The name of the property set
* `object_storage_bucket_compartment_id` - (Applicable when key=APEX_DOCUMENT_GENERATOR) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Object Storage Buckets managed by APEX
* `object_storage_endpoint` - (Applicable when key=APEX_DOCUMENT_GENERATOR) (Updatable) Object Storage Endpoint
* `object_storage_namespace` - (Applicable when key=APEX_DOCUMENT_GENERATOR) (Updatable) The Object Storage Namespace containing the Object Storage Buckets managed by APEX
* `print_server_type` - (Applicable when key=APEX_DOCUMENT_GENERATOR) (Updatable) The print server type
* `property_set_key` - (Required) The name of the property set


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `authentication_substitutions` - APEX FA Integration key-value pairs.
* `autonomous_database_resource_principal_status` - The status of the Autonomous Database Serverless Resource Principal (OCI$RESOURCE_PRINCIPAL)
* `credential_key` - The name of the credential used by APEX to manage Object Storage Buckets and Objects as well as invoke the Document Generator function
* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Document Generator function
* `identity_provider` - External identity type provider
	* `configs` - External identity provider configuration parameters. Simple key-value pair Example: { "tenant_id": "...", "application_id_uri": "...", ... } 
	* `type` - External identity type provider.  Supported values - OCI_IAM, AZURE_AD, NONE.
* `instance_dbms_credential_enabled` - Specifies whether database credentials can be used in all workspaces on the APEX instance. Supported values include: "Y", "N" and empty string.
* `invoke_endpoint` - The base endpoint URL to use to invoke the Document Generator function
* `is_mutable` - Indicates whether the property set is mutable or not
* `key` - The name of the property set
* `object_storage_bucket_compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Object Storage Buckets managed by APEX
* `object_storage_endpoint` - Object Storage Endpoint
* `object_storage_namespace` - The Object Storage Namespace containing the Object Storage Buckets managed by APEX
* `prerequisites_check` - The results of a prerequisites check for APEX FA integration
	* `status` - Status indicating the outcome of the prerequisites check.
	* `status_details` - Messages describing the prerequisites check outcome. Messages can provide actionable information when the status indicates a failure.
* `print_server_type` - The print server type
* `user_key` - The APEX engine schema name
* `version` - The version of APEX

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Connection Property Set
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Connection Property Set
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Connection Property Set


## Import

DatabaseToolsConnectionPropertySets can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_runtime_database_tools_connection_property_set.test_database_tools_connection_property_set "databaseToolsConnections/{databaseToolsConnectionId}/propertySets/{propertySetKey}" 
```

