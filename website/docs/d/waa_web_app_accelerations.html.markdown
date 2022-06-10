---
subcategory: "Waa"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waa_web_app_accelerations"
sidebar_current: "docs-oci-datasource-waa-web_app_accelerations"
description: |-
  Provides the list of Web App Accelerations in Oracle Cloud Infrastructure Waa service
---

# Data Source: oci_waa_web_app_accelerations
This data source provides the list of Web App Accelerations in Oracle Cloud Infrastructure Waa service.

Gets a list of all WebAppAccelerations in a compartment.


## Example Usage

```hcl
data "oci_waa_web_app_accelerations" "test_web_app_accelerations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.web_app_acceleration_display_name
	id = var.web_app_acceleration_id
	state = var.web_app_acceleration_state
	web_app_acceleration_policy_id = oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) A filter to return only the WebAppAcceleration with the given [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only resources that match the given lifecycleState.
* `web_app_acceleration_policy_id` - (Optional) A filter to return only the WebAppAcceleration with the given [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of related WebAppAccelerationPolicy. 


## Attributes Reference

The following attributes are exported:

* `web_app_acceleration_collection` - The list of web_app_acceleration_collection.

### WebAppAcceleration Reference

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

