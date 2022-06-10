---
subcategory: "Waa"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waa_web_app_acceleration"
sidebar_current: "docs-oci-resource-waa-web_app_acceleration"
description: |-
  Provides the Web App Acceleration resource in Oracle Cloud Infrastructure Waa service
---

# oci_waa_web_app_acceleration
This resource provides the Web App Acceleration resource in Oracle Cloud Infrastructure Waa service.

Creates a new WebAppAcceleration.


## Example Usage

```hcl
resource "oci_waa_web_app_acceleration" "test_web_app_acceleration" {
	#Required
	backend_type = var.web_app_acceleration_backend_type
	compartment_id = var.compartment_id
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	web_app_acceleration_policy_id = oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.web_app_acceleration_display_name
	freeform_tags = {"bar-key"= "value"}
	system_tags = var.web_app_acceleration_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `backend_type` - (Required) Type of the WebAppFirewall, as example LOAD_BALANCER.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) WebAppAcceleration display name, can be renamed.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `load_balancer_id` - (Required) LoadBalancer [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to which the WebAppAccelerationPolicy is attached to.
* `system_tags` - (Optional) (Updatable) Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `web_app_acceleration_policy_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of WebAppAccelerationPolicy, which is attached to the resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backend_type` - Type of the WebAppFirewall, as example LOAD_BALANCER.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - WebAppAcceleration display name, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebAppAcceleration.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in FAILED state. 
* `load_balancer_id` - LoadBalancer [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to which the WebAppAccelerationPolicy is attached to.
* `state` - The current state of the WebAppAcceleration.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the WebAppAcceleration was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the WebAppAcceleration was updated. An RFC3339 formatted datetime string.
* `web_app_acceleration_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of WebAppAccelerationPolicy, which is attached to the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Web App Acceleration
	* `update` - (Defaults to 20 minutes), when updating the Web App Acceleration
	* `delete` - (Defaults to 20 minutes), when destroying the Web App Acceleration


## Import

WebAppAccelerations can be imported using the `id`, e.g.

```
$ terraform import oci_waa_web_app_acceleration.test_web_app_acceleration "id"
```

