---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_api_validation"
sidebar_current: "docs-oci-datasource-apigateway-api_validation"
description: |-
  Provides details about a specific Api Validation in Oracle Cloud Infrastructure API Gateway service
---

# Data Source: oci_apigateway_api_validation
This data source provides details about a specific Api Validation resource in Oracle Cloud Infrastructure API Gateway service.

Gets the API validation results.

## Example Usage

```hcl
data "oci_apigateway_api_validation" "test_api_validation" {
	#Required
	api_id = oci_apigateway_api.test_api.id
}
```

## Argument Reference

The following arguments are supported:

* `api_id` - (Required) The ocid of the API.


## Attributes Reference

The following attributes are exported:

* `validations` - API validation results.
	* `name` - Name of the validation.
	* `result` - Result of the validation.

