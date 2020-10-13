---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_apis"
sidebar_current: "docs-oci-datasource-apigateway-apis"
description: |-
  Provides the list of Apis in Oracle Cloud Infrastructure API Gateway service
---

# Data Source: oci_apigateway_apis
This data source provides the list of Apis in Oracle Cloud Infrastructure API Gateway service.

Returns a list of APIs.


## Example Usage

```hcl
data "oci_apigateway_apis" "test_apis" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.api_display_name
	state = var.api_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state.  Example: `ACTIVE` 


## Attributes Reference

The following attributes are exported:

* `api_collection` - The list of api_collection.

### Api Reference

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

