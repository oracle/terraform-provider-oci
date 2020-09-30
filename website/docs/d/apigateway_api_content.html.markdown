---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_api_content"
sidebar_current: "docs-oci-datasource-apigateway-api_content"
description: |-
  Provides details about a specific Api Content in Oracle Cloud Infrastructure API Gateway service
---

# Data Source: oci_apigateway_api_content
This data source provides details about a specific Api Content resource in Oracle Cloud Infrastructure API Gateway service.

Get the raw API content.

## Example Usage

```hcl
data "oci_apigateway_api_content" "test_api_content" {
	#Required
	api_id = oci_apigateway_api.test_api.id
}
```

## Argument Reference

The following arguments are supported:

* `api_id` - (Required) The ocid of the API.


## Attributes Reference

The following attributes are exported:


