---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_data_source"
sidebar_current: "docs-oci-resource-generative_ai_agent-data_source"
description: |-
  Provides the Data Source resource in Oracle Cloud Infrastructure Generative Ai Agent service
---

# oci_generative_ai_agent_data_source
This resource provides the Data Source resource in Oracle Cloud Infrastructure Generative Ai Agent service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai-agents/latest/DataSource

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai_agent

Creates a data source.


## Example Usage

```hcl
resource "oci_generative_ai_agent_data_source" "test_data_source" {
	#Required
	compartment_id = var.compartment_id
	data_source_config {
		#Required
		data_source_config_type = var.data_source_data_source_config_data_source_config_type

		#Optional
		object_storage_prefixes {
			#Required
			bucket = var.data_source_data_source_config_object_storage_prefixes_bucket
			namespace = var.data_source_data_source_config_object_storage_prefixes_namespace

			#Optional
			prefix = var.data_source_data_source_config_object_storage_prefixes_prefix
		}
		should_enable_multi_modality = var.data_source_data_source_config_should_enable_multi_modality
	}
	knowledge_base_id = oci_generative_ai_agent_knowledge_base.test_knowledge_base.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.data_source_description
	display_name = var.data_source_display_name
	freeform_tags = {"Department"= "Finance"}
	metadata = var.data_source_metadata
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the data source in. 
* `data_source_config` - (Required) (Updatable) The details of data source. 
	* `data_source_config_type` - (Required) (Updatable) The type of the tool. 
	* `object_storage_prefixes` - (Optional) (Updatable) The locations of data items in Object Storage, can either be an object (File) or a prefix (folder).
		* `bucket` - (Required) (Updatable) The bucket name of an object.
		* `namespace` - (Required) (Updatable) The namespace name of an object.
		* `prefix` - (Optional) (Updatable) The prefix of file object(s) or folder prefix.
	* `should_enable_multi_modality` - (Optional) (Updatable) Flag to enable or disable multi modality such as image processing while ingestion of data. True enable the processing and false exclude the multi modality contents during ingestion.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A description of the data source.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `knowledge_base_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent KnowledgeBase.
* `metadata` - (Optional) (Updatable) Key-value pairs to allow additional configurations.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `data_source_config` - The details of data source. 
	* `data_source_config_type` - The type of the tool. 
	* `object_storage_prefixes` - The locations of data items in Object Storage, can either be an object (File) or a prefix (folder).
		* `bucket` - The bucket name of an object.
		* `namespace` - The namespace name of an object.
		* `prefix` - The prefix of file object(s) or folder prefix.
	* `should_enable_multi_modality` - Flag to enable or disable multi modality such as image processing while ingestion of data. True enable the processing and false exclude the multi modality contents during ingestion.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of the data source.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data source.
* `knowledge_base_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent KnowledgeBase.
* `lifecycle_details` - A message that describes the current state of the data source in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `metadata` - Key-value pairs to allow additional configurations.
* `state` - The current state of the data source.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the data source was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the data source was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Data Source
	* `update` - (Defaults to 20 minutes), when updating the Data Source
	* `delete` - (Defaults to 20 minutes), when destroying the Data Source


## Import

DataSources can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_agent_data_source.test_data_source "id"
```

