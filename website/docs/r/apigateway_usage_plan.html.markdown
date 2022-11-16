---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_usage_plan"
sidebar_current: "docs-oci-resource-apigateway-usage_plan"
description: |-
  Provides the Usage Plan resource in Oracle Cloud Infrastructure API Gateway service
---

# oci_apigateway_usage_plan
This resource provides the Usage Plan resource in Oracle Cloud Infrastructure API Gateway service.

Creates a new usage plan.

## Example Usage

```hcl
resource "oci_apigateway_usage_plan" "test_usage_plan" {
	#Required
	compartment_id = var.compartment_id
	entitlements {
		#Required
		name = var.usage_plan_entitlements_name

		#Optional
		description = var.usage_plan_entitlements_description
		quota {
			#Required
			operation_on_breach = var.usage_plan_entitlements_quota_operation_on_breach
			reset_policy = var.usage_plan_entitlements_quota_reset_policy
			unit = var.usage_plan_entitlements_quota_unit
			value = var.usage_plan_entitlements_quota_value
		}
		rate_limit {
			#Required
			unit = var.usage_plan_entitlements_rate_limit_unit
			value = var.usage_plan_entitlements_rate_limit_value
		}
		targets {
			#Required
			deployment_id = oci_apigateway_deployment.test_deployment.id
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.usage_plan_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `entitlements` - (Required) (Updatable) A collection of entitlements to assign to the newly created usage plan. 
	* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
	* `name` - (Required) (Updatable) An entitlement name, unique within a usage plan. 
	* `quota` - (Optional) (Updatable) Quota policy for a usage plan. 
		* `operation_on_breach` - (Required) (Updatable) What the usage plan will do when a quota is breached: `REJECT` will allow no further requests `ALLOW` will continue to allow further requests 
		* `reset_policy` - (Required) (Updatable) The policy that controls when quotas will reset. Example: `CALENDAR` 
		* `unit` - (Required) (Updatable) The unit of time over which quotas are calculated. Example: `MINUTE` or `MONTH` 
		* `value` - (Required) (Updatable) The number of requests that can be made per time period. 
	* `rate_limit` - (Optional) (Updatable) Rate-limiting policy for a usage plan. 
		* `unit` - (Required) (Updatable) The unit of time over which rate limits are calculated. Example: `SECOND` 
		* `value` - (Required) (Updatable) The number of requests that can be made per time period. 
	* `targets` - (Optional) (Updatable) A collection of targeted deployments that the entitlement will be applied to. 
		* `deployment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a deployment resource. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Usage Plan
	* `update` - (Defaults to 20 minutes), when updating the Usage Plan
	* `delete` - (Defaults to 20 minutes), when destroying the Usage Plan


## Import

UsagePlans can be imported using the `id`, e.g.

```
$ terraform import oci_apigateway_usage_plan.test_usage_plan "id"
```

