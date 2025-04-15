---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_templates"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_templates"
description: |-
  Provides the list of Namespace Templates in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_templates
This data source provides the list of Namespace Templates in Oracle Cloud Infrastructure Log Analytics service.

Returns a list of templates, containing detailed information about them. You may limit the number of results, provide sorting order, and filter by information such as template name, type, display name and description.


## Example Usage

```hcl
data "oci_log_analytics_namespace_templates" "test_namespace_templates" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.namespace_template_namespace

	#Optional
	name = var.namespace_template_name
	namespace_template_filter = var.namespace_template_namespace_template_filter
	state = var.namespace_template_state
	template_display_text = var.namespace_template_template_display_text
	type = var.namespace_template_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `name` - (Optional) The template name used for filtering. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `namespace_template_filter` - (Optional) filter
* `state` - (Optional) The template lifecycle state used for filtering. Currently supported values are ACTIVE and DELETED. 
* `template_display_text` - (Optional) The template display text used for filtering. Only templates with the specified name or description will be returned. 
* `type` - (Optional) The template type used for filtering. Only templates of the specified type will be returned. 


## Attributes Reference

The following attributes are exported:

* `log_analytics_template_collection` - The list of log_analytics_template_collection.

### NamespaceTemplate Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `content` - Base64 encoded template content.
* `content_format` - Content format. For example - XML.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description for this resource. 
* `facets` - Facets of the template
	* `name` - The facet name.
	* `value` - The facet value.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud. 
* `is_system` - The system flag.  A value of false denotes a custom, or user defined object.  A value of true denotes a built in object. 
* `name` - The template name.
* `parameters` - Base64 encoded template parameters.
* `parameters_format` - Parameters format.  For example - NAME_VALUE_PAIR.
* `parameters_metadata` - Base64 encoded parameters metadata definition.
* `state` - The current state of the template. 
* `time_created` - The date and time the resource was created, in the format defined by RFC3339. 
* `time_updated` - The date and time the resource was last updated, in the format defined by RFC3339. 
* `type` - The template type.

