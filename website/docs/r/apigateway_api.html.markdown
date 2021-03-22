---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_api"
sidebar_current: "docs-oci-resource-apigateway-api"
description: |-
  Provides the Api resource in Oracle Cloud Infrastructure API Gateway service
---

# oci_apigateway_api
This resource provides the Api resource in Oracle Cloud Infrastructure API Gateway service.

Creates a new API.


## Example Usage

```hcl
resource "oci_apigateway_api" "test_api" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	content = var.api_content
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.api_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `content` - (Optional) (Updatable) API Specification content in json or yaml format
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `lifecycle_details` - A message describing the current lifecycleState in more detail. For ACTIVE state it describes if the document has been validated and the possible values are:
	* 'New' for just updated API Specifications
	* 'Validating' for a document which is being validated.
	* 'Valid' the document has been validated without any errors or warnings
	* 'Warning' the document has been validated and contains warnings
	* 'Error' the document has been validated and contains errors
	* 'Failed' the document validation failed
	* 'Canceled' the document validation was canceled 

	For other states it may provide more details like actionable information. 
* `specification_type` - Type of API Specification file.
* `state` - The current state of the API.
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `validation_results` - Status of each feature available from the API.
	* `name` - Name of the validation.
	* `result` - Result of the validation.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Api
	* `update` - (Defaults to 20 minutes), when updating the Api
	* `delete` - (Defaults to 20 minutes), when destroying the Api


## Import

Apis can be imported using the `id`, e.g.

```
$ terraform import oci_apigateway_api.test_api "id"
```

