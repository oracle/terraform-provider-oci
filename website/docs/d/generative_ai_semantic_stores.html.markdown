---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_semantic_stores"
sidebar_current: "docs-oci-datasource-generative_ai-semantic_stores"
description: |-
  Provides the list of Semantic Stores in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_semantic_stores
This data source provides the list of Semantic Stores in Oracle Cloud Infrastructure Generative AI service.

Lists the SemanticStores given specific filter.

## Example Usage

```hcl
data "oci_generative_ai_semantic_stores" "test_semantic_stores" {

	#Optional
	compartment_id = var.compartment_id
	data_source_querying_connection_id = oci_database_migration_connection.test_connection.id
	display_name = var.semantic_store_display_name
	id = var.semantic_store_id
	state = var.semantic_store_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `data_source_querying_connection_id` - (Optional) A filter to return only resources whose queryingConnectionId matches with this id. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SemanticStore.
* `state` - (Optional) A filter to return only resources whose lifecycle state matches the given array.


## Attributes Reference

The following attributes are exported:

* `semantic_store_collection` - The list of semantic_store_collection.

### SemanticStore Reference

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

