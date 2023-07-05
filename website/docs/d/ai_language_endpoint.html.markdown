---
subcategory: "Ai Language"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_language_endpoint"
sidebar_current: "docs-oci-datasource-ai_language-endpoint"
description: |-
  Provides details about a specific Endpoint in Oracle Cloud Infrastructure Ai Language service
---

# Data Source: oci_ai_language_endpoint
This data source provides details about a specific Endpoint resource in Oracle Cloud Infrastructure Ai Language service.

Gets an endpoint by identifier

## Example Usage

```hcl
data "oci_ai_language_endpoint" "test_endpoint" {
	#Required
	endpoint_id = oci_ai_language_endpoint.test_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `endpoint_id` - (Required) The OCID of the endpoint.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the endpoint compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description of the endpoint.
* `display_name` - A user-friendly display name for the resource. It should be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier endpoint OCID of an endpoint that is immutable on creation.
* `inference_units` - Number of replicas required for this endpoint.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in failed state.
* `model_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model to associate with the endpoint.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the Endpoint.
* `state` - The state of the endpoint.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the endpoint was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the endpoint was updated. An RFC3339 formatted datetime string.

