---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_data_sources"
sidebar_current: "docs-oci-datasource-generative_ai_agent-data_sources"
description: |-
  Provides the list of Data Sources in Oracle Cloud Infrastructure Generative Ai Agent service
---

# Data Source: oci_generative_ai_agent_data_sources
This data source provides the list of Data Sources in Oracle Cloud Infrastructure Generative Ai Agent service.

**ListDataSources**

Gets a list of data sources.


## Example Usage

```hcl
data "oci_generative_ai_agent_data_sources" "test_data_sources" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.data_source_display_name
	knowledge_base_id = oci_generative_ai_agent_knowledge_base.test_knowledge_base.id
	state = var.data_source_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `knowledge_base_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the knowledge base.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `data_source_collection` - The list of data_source_collection.

### DataSource Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `data_source_config` - **DataSourceConfig**

	The details of data source. 
	* `data_source_config_type` - The type of the tool. The allowed values are:
		* `OCI_OBJECT_STORAGE`: The data source is Oracle Cloud Infrastructure Object Storage. 
	* `object_storage_prefixes` - The locations of data items in Object Storage, can either be an object (File) or a prefix (folder).
		* `bucket` - The bucket name of an object.
		* `namespace` - The namespace name of an object.
		* `prefix` - The name of the object (file) or prefix (folder).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of the data source.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data source.
* `knowledge_base_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent KnowledgeBase.
* `lifecycle_details` - A message that describes the current state of the data source in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `state` - The current state of the data source.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the data source was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the data source was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

