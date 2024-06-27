---
subcategory: "Delegate Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_delegate_access_control_service_provider_actions"
sidebar_current: "docs-oci-datasource-delegate_access_control-service_provider_actions"
description: |-
  Provides the list of Service Provider Actions in Oracle Cloud Infrastructure Delegate Access Control service
---

# Data Source: oci_delegate_access_control_service_provider_actions
This data source provides the list of Service Provider Actions in Oracle Cloud Infrastructure Delegate Access Control service.

Lists all the ServiceProviderActions available in the system.


## Example Usage

```hcl
data "oci_delegate_access_control_service_provider_actions" "test_service_provider_actions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.service_provider_action_name
	resource_type = var.service_provider_action_resource_type
	service_provider_service_type = var.service_provider_action_service_provider_service_type
	state = var.service_provider_action_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `resource_type` - (Optional) A filter to return only resources that match the given resource type.
* `service_provider_service_type` - (Optional) A filter to return only resources that match the given Service Provider service type.
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the given Service Provider Action lifecycleState.


## Attributes Reference

The following attributes are exported:

* `service_provider_action_summary_collection` - The list of service_provider_action_summary_collection.

### ServiceProviderAction Reference

The following attributes are exported:

* `component` - Name of the infrastructure layer associated with the Service Provider Action.
* `customer_display_name` - Display Name of the Service Provider Action.
* `description` - Description of the Service Provider Action in terms of associated risk profile, and characteristics of the operating system commands made available to the support operator under this Service Provider Action. 
* `id` - Unique Oracle assigned identifier for the Service Provider Action.
* `name` - Unique name of the Service Provider Action.
* `properties` - Fine grained properties associated with the Delegation Control.
	* `name` - Name of the property
	* `value` - value of the property
* `resource_type` - resourceType for which the ServiceProviderAction is applicable
* `service_provider_service_types` - List of Service Provider Service Types that this Service Provider Action is applicable to.
* `state` - The current lifecycle state of the Service Provider Action.

