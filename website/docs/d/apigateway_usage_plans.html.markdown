---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_usage_plans"
sidebar_current: "docs-oci-datasource-apigateway-usage_plans"
description: |-
  Provides the list of Usage Plans in Oracle Cloud Infrastructure API Gateway service
---

# Data Source: oci_apigateway_usage_plans
This data source provides the list of Usage Plans in Oracle Cloud Infrastructure API Gateway service.

Returns a list of usage plans.

## Example Usage

```hcl
data "oci_apigateway_usage_plans" "test_usage_plans" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.usage_plan_display_name
	state = var.usage_plan_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. Example: `ACTIVE` 


## Attributes Reference

The following attributes are exported:

* `usage_plan_collection` - The list of usage_plan_collection.

### UsagePlan Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `entitlements` - A collection of entitlements currently assigned to the usage plan. 
	* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
	* `name` - An entitlement name, unique within a usage plan. 
	* `quota` - Quota policy for a usage plan. 
		* `operation_on_breach` - What the usage plan will do when a quota is breached: `REJECT` will allow no further requests `ALLOW` will continue to allow further requests 
		* `reset_policy` - The policy that controls when quotas will reset. Example: `CALENDAR` 
		* `unit` - The unit of time over which quotas are calculated. Example: `MINUTE` or `MONTH` 
		* `value` - The number of requests that can be made per time period. 
	* `rate_limit` - Rate-limiting policy for a usage plan. 
		* `unit` - The unit of time over which rate limits are calculated. Example: `SECOND` 
		* `value` - The number of requests that can be made per time period. 
	* `targets` - A collection of targeted deployments that the entitlement will be applied to. 
		* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a deployment resource. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a usage plan resource. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state. 
* `state` - The current state of the usage plan.
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

