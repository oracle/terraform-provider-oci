---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_endpoints"
sidebar_current: "docs-oci-datasource-generative_ai-endpoints"
description: |-
  Provides the list of Endpoints in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_endpoints
This data source provides the list of Endpoints in Oracle Cloud Infrastructure Generative AI service.

Lists the endpoints of a specific compartment.

## Example Usage

```hcl
data "oci_generative_ai_endpoints" "test_endpoints" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.endpoint_display_name
	id = var.endpoint_id
	state = var.endpoint_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the endpoint.
* `state` - (Optional) A filter to return only resources that their lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `endpoint_collection` - The list of endpoint_collection.

### Endpoint Reference

The following attributes are exported:

* `description` - An optional description of the endpoint.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `model_id` - The OCID of the model that's used to create this endpoint.
* `state` - The current state of the endpoint.
* `time_updated` - The date and time that the endpoint was updated in the format of an RFC3339 datetime string.

