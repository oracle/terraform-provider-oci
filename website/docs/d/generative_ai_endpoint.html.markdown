---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_endpoint"
sidebar_current: "docs-oci-datasource-generative_ai-endpoint"
description: |-
  Provides details about a specific Endpoint in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_endpoint
This data source provides details about a specific Endpoint resource in Oracle Cloud Infrastructure Generative AI service.

Gets information about an endpoint.

## Example Usage

```hcl
data "oci_generative_ai_endpoint" "test_endpoint" {
	#Required
	endpoint_id = oci_generative_ai_endpoint.test_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `endpoint_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the endpoint.


## Attributes Reference

The following attributes are exported:

* `description` - An optional description of the endpoint.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `model_id` - The OCID of the model that's used to create this endpoint.
* `state` - The current state of the endpoint.
* `time_updated` - The date and time that the endpoint was updated in the format of an RFC3339 datetime string.

