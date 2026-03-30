---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_semantic_store"
sidebar_current: "docs-oci-resource-generative_ai-semantic_store"
description: |-
  Provides the Semantic Store resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_semantic_store
This resource provides the Semantic Store resource in Oracle Cloud Infrastructure Generative AI service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai/latest/SemanticStore

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai

Creates a SemanticStore.


## Example Usage

```hcl
resource "oci_generative_ai_semantic_store" "test_semantic_store" {
	#Required
	compartment_id = var.compartment_id
	data_source {
		#Required
		connection_type = var.semantic_store_data_source_connection_type
		enrichment_connection_id = oci_database_tools_database_tools_connection.test_connection.id
		querying_connection_id = oci_database_tools_database_tools_connection.test_connection.id
	}
	display_name = var.semantic_store_display_name
	schemas {
		#Required
		connection_type = var.semantic_store_schemas_connection_type
		schemas {
			#Required
			name = var.semantic_store_schemas_schemas_name
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.semantic_store_description
	freeform_tags = {"Department"= "Finance"}
	refresh_schedule {
		#Required
		type = var.semantic_store_refresh_schedule_type

		#Optional
		value = var.semantic_store_refresh_schedule_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Owning compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for a SemanticStore.
* `data_source` - (Required) Defines the data source that the semantic model connects to.
	* `connection_type` - (Required) Specifies the type of underlying connection.
	* `enrichment_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure Database Tools Connection for enrichment.
	* `querying_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure Database Tools Connection for querying.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) An optional description of the SemanticStore.
* `display_name` - (Required) (Updatable) A user-friendly name.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `refresh_schedule` - (Optional) (Updatable) Specifies a refresh schedule. Null represents no automated synchronization schedule.
	* `type` - (Required) (Updatable) Specifies the type of refresh schedule.  
	* `value` - (Required when type=INTERVAL) (Updatable) Specifies the refresh interval value. The interval must be provided using the ISO 8601 extended format, either as PnW or PnYnMnDTnHnMnS,  where 'P' is always required, 'T' precedes any time components less than one day, and each included component is properly suffixed.  For example, "P1DT6H" represents a duration of 1 day and 6 hours. 
* `schemas` - (Required) (Updatable) Array of database schemas or other database objects to include in enrichment pipeline.
	* `connection_type` - (Required) (Updatable) Specifies the type of underlying connection.
	* `schemas` - (Required) (Updatable) Array of database schemas to be included in the connection. Each schema must define a name. A simple schema definition includes only the name, for example: { "schemas": [ { "name": "HR" } ] } Only one schema name is allowed now. Additional configuration options may be supported in extended forms later. 
		* `name` - (Required) (Updatable) 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Owning compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for a SemanticStore.
* `data_source` - Defines the data source that the semantic model connects to.
	* `connection_type` - Specifies the type of underlying connection.
	* `enrichment_connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure Database Tools Connection for enrichment.
	* `querying_connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure Database Tools Connection for querying.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - An optional description of the SemanticStore.
* `display_name` - A user-friendly name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - An [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that uniquely identifies a SemanticStore.
* `lifecycle_details` - A message describing the current state in more detail that can provide actionable information.
* `refresh_schedule` - Specifies a refresh schedule. Null represents no automated synchronization schedule.
	* `type` - Specifies the type of refresh schedule.  
	* `value` - Specifies the refresh interval value. The interval must be provided using the ISO 8601 extended format, either as PnW or PnYnMnDTnHnMnS,  where 'P' is always required, 'T' precedes any time components less than one day, and each included component is properly suffixed.  For example, "P1DT6H" represents a duration of 1 day and 6 hours. 
* `schemas` - Array of database schemas or other database objects to include in enrichment pipeline.
	* `connection_type` - Specifies the type of underlying connection.
	* `schemas` - Array of database schemas to be included in the connection. Each schema must define a name. A simple schema definition includes only the name, for example: { "schemas": [ { "name": "HR" } ] } Only one schema name is allowed now. Additional configuration options may be supported in extended forms later. 
		* `name` - 
* `state` - The lifecycle state of a SemanticStore.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the SemanticStore was created in the format of an RFC3339 datetime string.
* `time_updated` - The date and time that the SemanticStore was updated in the format of an RFC3339 datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Semantic Store
	* `update` - (Defaults to 20 minutes), when updating the Semantic Store
	* `delete` - (Defaults to 20 minutes), when destroying the Semantic Store


## Import

SemanticStores can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_semantic_store.test_semantic_store "id"
```

