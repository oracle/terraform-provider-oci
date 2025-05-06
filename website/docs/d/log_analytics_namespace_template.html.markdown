---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_template"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_template"
description: |-
  Provides details about a specific Namespace Template in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_template
This data source provides details about a specific Namespace Template resource in Oracle Cloud Infrastructure Log Analytics service.

Gets detailed information about the template with the specified ocid.


## Example Usage

```hcl
data "oci_log_analytics_namespace_template" "test_namespace_template" {
	#Required
	namespace = var.namespace_template_namespace
	template_id = oci_log_analytics_template.test_template.id
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `template_id` - (Required) Unique ocid of the template. 


## Attributes Reference

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

