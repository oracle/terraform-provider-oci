---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_connection_property_set"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_connection_property_set"
description: |-
  Provides details about a specific Database Tools Connection Property Set in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_connection_property_set
This data source provides details about a specific Database Tools Connection Property Set resource in Oracle Cloud Infrastructure Database Tools Runtime service.

Get a property set

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_connection_property_set" "test_database_tools_connection_property_set" {
	#Required
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
	property_set_key = var.database_tools_connection_property_set_property_set_key
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.
* `property_set_key` - (Required) The name of the property set


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

